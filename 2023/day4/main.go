package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
	"strconv"
	"slices"
)

func check(err error){
	if err != nil {
		panic(err)
	}	
}

type Cards struct {
	win []int
	my []int
}

func parseInput(filename string) []Cards {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`\d+`)
	var cardArr []Cards

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line,"|")
		winningStr := r.FindAllString(split[0], -1)[1:]
		myCardsStr := r.FindAllString(split[1], -1)	
		var winning []int
		var myCards []int
		for _, winCard := range winningStr {
			parsed, _ := strconv.Atoi(winCard)
			winning = append(winning, parsed)
		}
		for _, card := range myCardsStr {
			parsed, _ := strconv.Atoi(card)
			myCards = append(myCards, parsed)
		}
		cardArr = append(cardArr, Cards{winning, myCards})
	}
	return cardArr
}

func countCards(cards []Cards) int{
	total := 0
	for _, row := range cards {
		handTotal := 0
		for _, mine := range row.my {
			if slices.Contains(row.win, mine){
				if handTotal == 0 {
					handTotal = 1
				} else {
					handTotal = handTotal * 2
				}
			}
		}
		total += handTotal
	}
	return total
}

func main(){
	cards := parseInput("input.txt")
	total := countCards(cards)
	fmt.Println(total)
}
