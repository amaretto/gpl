package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// 接続先を定義
type clock struct {
	host string
	port string
}

var clocks = []clock{
	{"localhost", "8010"},
	//	{"localhost", "8020"},
	//	{"localhost", "8030"},
}

func main() {
	var timeChans []chan string
	for _, target := range clocks {
		timeChan := make(chan string, 1)
		timeChans = append(timeChans, timeChan)
		go getTime(100*time.Millisecond, target, timeChan)
	}
	displayTime(100*time.Millisecond, timeChans)
}

func getTime(delay time.Duration, target clock, timeChan chan string) {
	conn, err := net.Dial("tcp", target.host+":"+target.port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		clockTime, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		timeChan <- clockTime
		time.Sleep(delay)
	}
}

func displayTime(delay time.Duration, timeChans []chan string) {
	var temp, timeString string
	for {
		for _, timeChan := range timeChans {
			temp = <-timeChan
			timeString += strings.TrimRight(temp, "\n") + "\t"
		}
		fmt.Printf("\r%s", timeString)
		time.Sleep(delay)
		timeString = ""
	}
}
