package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var socketFilePath string

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		log.Printf("Server got: %s", string(data))
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func main() {
	flag.StringVar(&socketFilePath, "socketfile", "/tmp/domainsocketexample.sock", "file path of domain socket")
	flag.Parse()

	f, err := os.OpenFile("/tmp/domainsocketexample_server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println("domainsocketexample server start")

	l, err := net.Listen("unix", socketFilePath)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(fd)
	}
}
