package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// go run container.go run <cmd> <args>

func main() {

	log.Printf("PID %d\n", os.Getpid())

	switch os.Args[1] {
	case "run":
		log.Println("run running")
		run()

	case "child":
		log.Println("child running")
		child()
	default:
		panic("invalid command ")
	}

}

func run() {
	log.Println("[RUN]")

	log.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())

	// run 이후 args 에 child 가 추가 된다.
	args := append([]string{"child"}, os.Args[2:]...)

	// 외부의 프로그램을 go 프로그램 내부에서 실행 할 수 있는 exec
	// cmd := exec.Command(os.Args[2], os.Args[3:]...)

	// 스스로를 다시 실행 시킨다.
	cmd := exec.Command("/proc/self/exe", args...)

	// cmd stdin, stdout, stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	//cmd 가 실행이 된다.
	cmd.Run()
}

func child() {
	log.Println("[CHILD]")
	log.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())

	syscall.Sethostname([]byte("container-demo"))
	controlgroup()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	syscall.Chroot("/containerfs")

	os.Chdir("/")

	syscall.Mount("proc", "proc", "proc", 0, "")

	cmd.Run()

}

func controlgroup() {
	cgPath := filepath.Join("/sys/fs/cgroup/memory", "prabhu")
	os.Mkdir(cgPath, 0755)
	os.WriteFile(filepath.Join(cgPath, "memory.limit_in_bytes"), []byte("100000000"), 0700)
	os.WriteFile(filepath.Join(cgPath, "memory.swappiness"), []byte("100000000"), 0700)
	os.WriteFile(filepath.Join(cgPath, "tasks"), []byte("100000000"), 0700)
}

// https://itnext.io/container-from-scratch-348838574160
