package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/iovisor/gobpf/bcc"
)

const source string = `
#include <uapi/linux/ptrace.h>

int hello(void *ctx) {
    bpf_trace_printk("Hello World! Someone connected via SSH\\n");
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

	fmt.Println("Press Ctrl+C to stop...")

	// Listen for Ctrl+C signal to detach the program and exit
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

	// Close the module to detach the eBPF program and unload the module
	module.Close()
}
