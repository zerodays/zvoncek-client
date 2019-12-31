package main

import (
	"bufio"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"net"
	"time"
)

const (
	address string = "51.77.83.245:69"
	pin     int    = 3
	sleep   int    = 50
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(pin) // Perhaps this is not the write pin number
	pin.Output()
	ch := make(chan bool, 1000)
	go bang(pin, ch)

	for {
		conn, _ := net.Dial("tcp", address)
		for {
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println(err)
				break
			}

			if msg == "bang" {
				ch <- true
			}
		}
	}
}

func bang(pin rpio.Pin, ch chan bool) {
	for {
		shouldBang, ok := <-ch
		if !ok {
			break
		}

		if shouldBang {
			pin.Low()
			time.Sleep(time.Duration(sleep) * time.Millisecond)
			pin.High()
		}
	}
}
