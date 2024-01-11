package main

import (
	"machine"
	"time"
)

var led = machine.PC18

func main() {

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 2000)
		hello()
		newWord()
		world()
		end()

	}

}

func dot() {
	led.Low()
	time.Sleep(time.Millisecond * 100)
	led.High()
	time.Sleep(time.Millisecond * 500)
}
func dash() {
	led.Low()
	time.Sleep(time.Millisecond * 100)
	led.High()
	time.Sleep(time.Millisecond * 1000)
}

func space() {
	led.Low()
	time.Sleep(time.Millisecond * 1000)

}
func end() { 
	newword()
	newword()
	newword()


}
func hello() {
	dot()
	dot()
	dot()
	dot()
	space()
	dot()
	space()
	dot()
	dash()
	dot()
	dot()
	space()
	dot()
	dash()
	dot()
	dot()
	space()
	dash()
	dash()
	dash()

}

func world() {

	dot()
	dash()
	dash()
	space()
	dash()
	dash()
	dash()
	space()
	dot()
	dash()
	dot()
	space()
	dot()
	dash()
	dot()
	dot()

}

func newWord() {

	space()
	space()

}
