package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var socketFilePath string

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got:", string(buf[0:n]))
	}
}

func main() {
	flag.StringVar(&socketFilePath, "socketfile", "/tmp/domainsocketexample.sock", "file path of domain socket")
	flag.Parse()

	c, err := net.Dial("unix", socketFilePath)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	go reader(c)
	for {
		_, err := c.Write([]byte("hi"))
		if err != nil {
			log.Fatal("write error:", err)
			break
		}
		time.Sleep(time.Second)
	}
}
