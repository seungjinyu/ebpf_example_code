package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {

	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock")
	}

	var objs xdpObjects
	if err := loadXdpObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	ifname := "eth0"
	iface, err := net.InterfaceByName(ifname)

	if err != nil {
		log.Fatalf("Getting interface %s : %s", ifname, err)
	}

	link, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.XdpProgSimple,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatal("Attaching XDP:", err)
	}

	defer link.Close()

	log.Printf("attached the byte code on XDP %s..", ifname)
	tick := time.Tick(time.Second)
	stop := make(chan os.Signal, 5)
	signal.Notify(stop, os.Interrupt)

	for {

		select {

		case <-tick:
			log.Println("ticking")

		case <-stop:
			log.Println("Signal accepted")
			return
		}
	}
}
