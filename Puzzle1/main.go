package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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
	digitStrs := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		line := scanner.Text()
		firstStrDigit := -1
		firstDigitVal := 0
		lastStrDigit := -1
		lastDigitVal := 0
		for i, s := range digitStrs {
			index := strings.Index(line, s)
			if index != -1 && (index < firstStrDigit || firstStrDigit == -1) {
				firstStrDigit = index
				firstDigitVal = i + 1
			}

			lastIndex := strings.LastIndex(line, s)
			if index != -1 && (index > lastStrDigit || lastStrDigit == -1) {
				lastStrDigit = lastIndex
				lastDigitVal = i + 1
			}
		}

		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			char := runes[i]
			if char > zeroRune && char <= nineRune && (firstStrDigit == -1 || (firstStrDigit != -1 && i < firstStrDigit)) { 
				firstDigitVal = int(char - '0')
				break
			}
		}
		for i := len(runes)-1; i >= 0; i-- {
			char := runes[i]
			if char > zeroRune && char <= nineRune && (lastStrDigit == -1 || (lastStrDigit != -1 && i > lastStrDigit)) {
				lastDigitVal = int(char - '0') 
				break
			}
		}
		lineNumber := firstDigitVal * 10 + lastDigitVal
		fmt.Println(lineNumber)
		total += lineNumber
	}
	fmt.Println(total)
}
