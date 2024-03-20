package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func main() {

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	_, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	port := "8080"

	wg.Add(1)

	go func() {
		tcpListener(&port, cancel)

	}()
	go func() {
		<-quit
		log.Printf("Gracefully shutting down")
		defer wg.Done()
	}()

	wg.Wait()

}

func tcpListener(port *string, cancel context.CancelFunc) {
	log.Printf("Starting TCP Server , listening on port %s ", *port)
	ln, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Could not start TCP Server: %v", err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Unable to accept incoming TCP connection: %v", err)

		}
		go tcpHandleConnection(conn, cancel)
	}

}

func tcpHandleConnection(conn net.Conn, cancel context.CancelFunc) {
	log.Printf("TCP Client %s connected", conn.RemoteAddr().String())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		tcpHandleCommand(conn, scanner.Text(), cancel)
	}
	log.Printf("TCP Client %s disconnected", conn.RemoteAddr().String())
}

func tcpHandleCommand(conn net.Conn, txt string, cancel context.CancelFunc) {
	log.Printf("TCP txt: %v", txt)

	response, addACK, err := handleResponse(txt, cancel)
	if err != nil {
		response = "ERROR: " + response + "\n"
	} else if addACK {
		response = "ACK TCP: " + response + "\n"
	}

	tcpRespond(conn, response)

	if response == "EXIT" {
		exit()
	}
}

func handleResponse(txt string, cancel context.CancelFunc) (response string, addACK bool, responseError error) {
	parts := strings.Split(strings.TrimSpace(txt), " ")
	response = txt
	addACK = true
	responseError = nil

	switch parts[0] {
	case "PRINT":
		log.Printf("print cmd")
		return "print call", true, nil
	}

	return
}

func tcpRespond(conn net.Conn, txt string) {
	log.Printf("Responding to TCP with %q", txt)
	if _, err := conn.Write([]byte(txt + "\n")); err != nil {
		log.Fatalf("Could not write to TCP stream: %v", err)
	}
}

func exit() {
	log.Printf("Received EXIT command. Exiting.")
	// This tells Agones to shutdown this Game Server
	// shutdownErr := s.Shutdown()
	// if shutdownErr != nil {
	// 	log.Printf("Could not shutdown")
	// }
	// The process will exit when Agones removes the pod and the
	// container receives the SIGTERM signal

}
