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

func findGears(engineMap []string) [][]int {
	var gears [][]int
	gearRegexp := regexp.MustCompile(`\*`)
	for _, line := range engineMap {
		gearBounds := gearRegexp.FindAllStringIndex(line, -1)
		var lineGears []int
		for _, bounds := range gearBounds {
			lineGears = append(lineGears, bounds[0])
		}
		gears = append(gears, lineGears)
	}
	return gears
}

func adjacentNums(index int, nums []NumPair) []int {
	var inRangeNums []int
	for _, num := range nums {
		if index <= num.bounds[1] && index >= num.bounds[0] - 1 {
			inRangeNums = append(inRangeNums, num.val)	
		}
	}
	return inRangeNums
}

func calcGears(engineMap []string, gears [][]int, nums [][]NumPair) int {
	total := 0
	for i, line := range gears {
		for _, gearIndex := range line {
			var validNums []int
			// Line above
			if i != 0 {
				validNums = append(validNums, adjacentNums(gearIndex, nums[i - 1])...)
			}
			// Line below
			if i != len(gears) - 1 {
				validNums = append(validNums, adjacentNums(gearIndex, nums[i + 1])...)
			}
			validNums = append(validNums, adjacentNums(gearIndex, nums[i])...)

			if len(validNums) == 2 {
				total += validNums[0] * validNums[1]
			}
		}
	}
	return total
}

func main() {
	engineMap := parseInput("input.txt")
	gears := findGears(engineMap)
	nums := findNums(engineMap)
	total := calcGears(engineMap, gears, nums)
	fmt.Println(total)
}
