package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	log.Println("Container from Scratch")

	currentPid := os.Getegid()

	log.Println(currentPid)

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Run()

}

// https://itnext.io/container-from-scratch-348838574160
