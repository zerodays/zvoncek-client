package main

import (
	"bufio"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"net"
	"time"
)

const (
	address       string = "51.77.83.245:8069"
	pin           int    = 3
	sleepExtended int    = 100
	sleepBetween  int    = 3000
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(pin) // Perhaps this is not the write pin number
	pin.Output()
	pin.High()

	ch := make(chan bool, 1000)
	go bang(pin, ch)

	for {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}

		for {
			msg, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println(err)
				break
			}

			if msg == "bang\n" {
				ch <- true
			}
		}
		_ = conn.Close()
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
			time.Sleep(time.Duration(sleepExtended) * time.Millisecond)
			pin.High()

			time.Sleep(time.Duration(sleepBetween) * time.Millisecond)
		}
	}
}
