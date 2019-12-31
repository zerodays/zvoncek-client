package main

import (
	"bufio"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os"
	"time"
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer rpio.Close()

	pin := rpio.Pin(27) // Perhaps this is not the write pin number
	pin.Output()
	pin.Low()

	reader := bufio.NewReader(os.Stdin)
	for {
		_, _ = reader.ReadString('\n')
		pin.High()
		time.Sleep(50 * time.Millisecond)
		pin.Low()
	}
}
