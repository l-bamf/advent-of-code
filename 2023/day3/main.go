package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func check(err error){
	if err != nil {
		panic(err)
	}
}

func parseInput(filename string) []string{
	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	
	var engineMap []string
	for scanner.Scan() {
		line := scanner.Text()
		engineMap = append(engineMap, line)
	}
	return engineMap
}

type NumPair struct {
	val int
	bounds []int
}

func findNums(engineMap []string) [][]NumPair {
	var pairs [][]NumPair
	for _, line := range engineMap {
		r := regexp.MustCompile(`[0-9]+`)	
		nums := r.FindAllString(line, -1)
		allBounds := r.FindAllStringIndex(line, -1)
		var linePairs []NumPair
		for j, num := range nums {
			bounds := allBounds[j]
			parsedVal, _ := strconv.Atoi(num)
			linePairs = append(linePairs, NumPair{parsedVal, bounds})
		}
		pairs = append(pairs, linePairs)
	}
	return pairs
}

func isPart (rowIndex int, numPair NumPair, engineMap []string) bool{
	r := regexp.MustCompile(`[\!\@\#\$\%\^\&\*\(\)\-\_\+\=\\\/]`)

	left := numPair.bounds[0] - 1
	if left < 0 {
		left = 0
	}

	right := numPair.bounds[1] + 1
	lineLength := len(engineMap[0])
	if right >= lineLength{
		right = lineLength 
	}

	if rowIndex != 0 {
		rowAbove := engineMap[rowIndex - 1]
		trimmedRow := rowAbove[left:right]
		if r.MatchString(trimmedRow) {
			return true
		} else {
			fmt.Println(trimmedRow)
		}
	}

	if rowIndex != len(engineMap) -1 {
		rowBelow := engineMap[rowIndex + 1]
		trimmedRow := rowBelow[left:right]
		if r.MatchString(trimmedRow) {
			return true
		} else {
			fmt.Println(trimmedRow)
		}
	}
	
	trimmedRow := engineMap[rowIndex][left:right]
	if r.MatchString(trimmedRow){
		return true
	} else {
		fmt.Println(trimmedRow)
	}
	fmt.Println(numPair.val)
	return false
}

func main() {
	engineMap := parseInput("input.txt")
	pairs := findNums(engineMap)
	total := 0
	for i, linePairs := range pairs {
		for _, pair := range linePairs {
			if isPart(i, pair, engineMap) {
				total += pair.val
			}
		}
	}
	fmt.Println(total)
}
