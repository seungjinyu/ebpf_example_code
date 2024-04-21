package main

import (
	"log"
	"time"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

const mapKey uint32 = 0

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go bpf kprobe_percpu.c -- -I../headers

func main() {

	fn := "sys_execve"

	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}
	// Load pre-compiled programs and maps into the kernel.
	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	// Open a Kprobe at the entry point of the kernel function and attach the
	// pre-compiled program. Each time the kernel function enters, the program
	// will increment the execution counter by 1. The read loop below polls this
	// map value once per second.
	kp, err := link.Kprobe(fn, objs.KprobeExecve, nil)
	if err != nil {
		log.Fatalf("opening kprobe: %s", err)
	}
	defer kp.Close()

	// Read loop reporting the total amount of times the kernel
	// function was entered, once per second.
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	log.Println("Waiting for events..")

	for range ticker.C {
		var all_cpu_value []uint64
		if err := objs.KprobeMap.Lookup(mapKey, &all_cpu_value); err != nil {
			log.Fatalf("reading map: %v", err)
		}
		for cpuid, cpuvalue := range all_cpu_value {
			log.Printf("%s called %d times on CPU%v\n", fn, cpuvalue, cpuid)
		}
		log.Printf("\n")
	}
}
