package main

import (
	"fmt"
	"os"
	"bufio"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello advent")
	file, err := os.Open(os.Args[1])
	check(err)

	scanner := bufio.NewScanner(file)
	total := 0
	zeroRune := rune('0')
	nineRune := rune('9')
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		lineNumber := 0
		for i := 0; i < len(runes); i++ {
			char := runes[i]
			if char > zeroRune && char <= nineRune {
				lineNumber += int(char - '0') * 10
				break
			}
		}
		for i := len(runes)-1; i >= 0; i-- {
			char := runes[i]
			if char > zeroRune && char <= nineRune {
				lineNumber += int(char - '0') 
				break
			}
		}
		fmt.Println(lineNumber)
		total += lineNumber
	}
	fmt.Println(total)
}
