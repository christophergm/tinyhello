package main

import (
    "machine" ///grandcentral-m4"
    "time"
	// "tinygo.org/x/drivers/ws2812"
)

func main() {
	led := machine.PC24
	led = machine.PC18
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    // n := ws2812.New(machine.PC24)
	for {
        led.Low()
        time.Sleep(time.Microsecond * 500)

        led.High()
        time.Sleep(time.Microsecond * 5000)
    }
}