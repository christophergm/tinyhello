package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define your character set as a slice of strings
var charSet = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
	"q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
}

func main() {
	// Example usage
	encodedMessage := sendPackage("Hello, World!")
	decodedMessage := receivePackage(encodedMessage)

	fmt.Println("Decoded message:", decodedMessage)
}

func sendPackage(data string) string {
	var binData []string

	fmt.Println("Sending package:", data)

	for _, char := range data {
		index := getIndex(char)
		binaryRep := IntegerToBinary(index)

		for _, bit := range binaryRep {
			if bit == '0' {
				fmt.Print("0")
				binData = append(binData, "0")
			} else if bit == '1' {
				fmt.Print("1")
				binData = append(binData, "1")
			}
		}
	}

	fmt.Println("\nPackage sent")

	// Return the binary data as a single string
	return strings.Join(binData, "")
}

func receivePackage(data string) string {
	fmt.Println("Receiving package:", data)

	var bytes []string
	for i := 0; i < len(data); i += 8 {
		end := i + 8
		if end > len(data) {
			end = len(data)
		}
		bytes = append(bytes, data[i:end])
	}

	var receivedChars []string

	for _, byteStr := range bytes {
		index, err := strconv.ParseInt(byteStr, 2, 0)
		if err != nil {
			fmt.Printf("Error parsing binary string %s: %v\n", byteStr, err)
			continue
		}

		if int(index) < len(charSet) && index >= 0 {
			char := charSet[index]
			receivedChars = append(receivedChars, char)
		} else {
			fmt.Printf("Invalid index %d for binary string %s\n", index, byteStr)
		}
	}

	receivedMessage := strings.Join(receivedChars, "")

	fmt.Println("Received characters:", receivedMessage)
	fmt.Println("Package received")

	return receivedMessage
}

func getIndex(char rune) int {
	for i, ch := range charSet {
		if ch == string(char) {
			return i
		}
	}
	return -1
}

func IntegerToBinary(n int) string {
	return fmt.Sprintf("%08b", n)
}
