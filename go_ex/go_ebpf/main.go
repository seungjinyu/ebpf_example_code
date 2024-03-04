package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/iovisor/gobpf/bcc"
)

const source string = `
#include <uapi/linux/ptrace.h>

int hello(void *ctx) {
    bpf_trace_printk("Hello World! Someone connected via SSH\n");
    return 0;
}
`

func main() {
	// Create a new eBPF module
	module := bcc.NewModule(source, []string{})

	// Attach the eBPF program to the SSH login event
	sshLogin, err := module.LoadKprobe("hello")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load Kprobe: %v\n", err)
		os.Exit(1)
	}

	err = module.AttachKprobe("tcp_v4_connect", sshLogin, -1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach Kprobe: %v\n", err)
		os.Exit(1)
	}

	// Open the trace pipe to read eBPF program output
	tracePipe := make(chan []byte)
	perfMap, err := bcc.InitPerfMap(module, "events", tracePipe)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize PerfMap: %v\n", err)
		os.Exit(1)
	}

	// Start the PerfMap
	perfMap.Start()

	// Listen for Ctrl+C signal to stop the PerfMap and exit
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalCh
		perfMap.Stop()
		os.Exit(0)
	}()

	// Print the SSH connection logs
	for {
		data := <-tracePipe
		message := string(data)
		if strings.Contains(message, "Hello World! Someone connected via SSH") {
			fmt.Println(message)
		}
	}
}
