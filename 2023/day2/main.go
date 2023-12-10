package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	check(err)

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		halves := strings.Split(line, ":")
		id, err := strconv.Atoi(halves[0][5:])
		check(err)

		sets := strings.Split(halves[1], ";")

		maxBlue := 0
		maxRed := 0
		maxGreen := 0

		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, draw := range balls {
				split := strings.Split(draw[1:], " ")
				num, err := strconv.Atoi(split[0])
				check(err)
				if split[1] == "blue" && num > maxBlue {
					maxBlue = num
				}
				if split[1] == "red" && num > maxRed {
					maxRed = num
				}
				if split[1] == "green" && num > maxGreen {
					maxGreen = num
				}
			}
		}
		power := maxBlue * maxRed * maxGreen
		fmt.Println(id, power)
		total += power
	}
	fmt.Println(total)
}
