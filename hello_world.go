package main
import (
	"machine"
	"time"
	)


func main() {
	
	led := machine.PC18
	led.configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		Led.Low()
		time.Sleep(time.Millisecond * 1000)
		hello()
		newWord()
		world()
 		newWord()
		newWord()
		newWord()



	} 

	println("hello world")
	
}

func dot() {
	Led.Low()
	time.Sleep(time.Millisecond * 100)
	Led.High()
	time.Sleep(time.Millisecond * 500)
}
func dash() {
	Led.Low()
	time.Sleep(time.Millisecond * 100)
	Led.High()
	time.Sleep(time.Millisecond * 1000)
}

func space() {
	Led.Low()
	time.Sleep(time.Millisecond * 1000)



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


}
func world(){

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
