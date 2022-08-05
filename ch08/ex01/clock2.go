package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

// Server Location is Asia/Tokyo
const ASIA = "Asia/tokyo"
const US = "US/Eastern"
const EU = "Europe/London"

func main() {
	host := flag.String("host", "localhost", "host ip")
	port := flag.String("port", "8080", "port number")
	flag.Parse()

	var tz string = ASIA
	if timezone := os.Getenv("TZ"); timezone != "" {
		tz = timezone
	}

	listener, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, tz)
	}
}

func handleConn(c net.Conn, tz string) {
	defer c.Close()
	var err error
	for {
		if tz == ASIA {
			_, err = io.WriteString(c, time.Now().Format("15:04:05\n"))
		} else if tz == US {
			_, err = io.WriteString(c, time.Now().Add(-time.Hour*14).Format("15:04:05\n"))
			fmt.Println("hoge")
		} else if tz == EU {
			_, err = io.WriteString(c, time.Now().Add(-time.Hour*9).Format("15:04:05\n"))
		}
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
