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
		valid := true

		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, draw := range balls {
				split := strings.Split(draw[1:], " ")
				num, err := strconv.Atoi(split[0])
				check(err)
				if split[1] == "blue" && num > 14 {
					valid = false
				}
				if split[1] == "red" && num > 12 {
					valid = false
				}
				if split[1] == "green" && num > 13 {
					valid = false
				}
			}
		}
		fmt.Println(id, valid)

		if valid {
			total += id
		}

	}
	fmt.Println(total)
}
