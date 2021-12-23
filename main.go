package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var port, name string

func main() {

	// use the executable name if not passed
	name = os.Getenv("NAME")
	if len(name) == 0 {
		if _, err := exec.LookPath(os.Args[0]); err != nil {
			log.Print(err)
			os.Exit(1)
		}
		name = filepath.Base(os.Args[0])
	}

	// assign and validate port
	port = os.Getenv("PORT")
	if len(port) == 0 {
		port = "1337"
	}
	if n, _ := strconv.Atoi(port); !(0 < n && n <= 65535) {
		log.Printf("Port must be between 1 and 65535")
		os.Exit(1)
	}

	// start listening
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Print(err)
	}

	log.Printf("Listening on port %v with name %q", port, name)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Println("Accepted ", conn.RemoteAddr())
		conn.Write([]byte("Well hello there!\n> "))
		go handleConnection(conn)
	}

}

func write(i interface{}, conn net.Conn) {
	w := bufio.NewWriter(conn)
	w.WriteString(fmt.Sprintf("%v\n> ", i))
	w.Flush()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	s := bufio.NewScanner(conn)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		switch line {
		case "exit":
			write("Goodbye.", conn)
			time.Sleep(500 * time.Millisecond)
			return
		case "host":
			host, _ := os.Hostname()
			write(host, conn)
		case "ip":
			write(conn.LocalAddr(), conn)
		case "whoami":
			write(conn.RemoteAddr(), conn)
		case "name":
			write(name, conn)
		case "", "help":
			fallthrough
		default:
			write("usage: (help|host|ip|whoami|name|exit)", conn)
		}
	}
}
