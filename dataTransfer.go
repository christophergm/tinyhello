package main

import (
	"machine"
	"time"
)

var led = machine.PC18
var sync = machine.PC19
var receve = machine.PC20
var send = machine.PC21
func main() {
	string charSet = ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","P","q","r","s","t","u","v","w","x","y","z","1","2","3","4","5","6","7","8","9","0"];
	sendPackage("Hello, World");
	
}

func sendPackage(string data) {
	send.high()
	for i := 0; i := range data {	
		dex = charSet.index(data);
	}	
	
	dex = IntegerToBinary(dex);
	dex = strconv.FormatInt(int64(dex), 10);
	
	for i := 0; i := range dex {
		if i == "0" && sync == high() && receve == low() {
			led.low();
		} else if i == "1" && sync == high() && receve == low() {
			led.high();
		}
	} 
	send.low()
}

func recevePackage(dataIn) {
	receve.high()
	for i := 0; i := range data {	
		dex = charSet.index(data);
	}	
	
	dex = IntegerToBinary(dex);
	dex = strconv.FormatInt(int64(dex), 10);
	
	for i := 0; i := range dex {
		if i == "0" && sync == high() && send != high() {
			fmt.Println("0");
		} else if i == "1" && sync == high() && send != high() {
			fmt.Println("1");
		}
	} 
}

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
