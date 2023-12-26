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
	num int
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
		cardNum, _ := strconv.Atoi(r.FindAllString(split[0], -1)[0])
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
		cardArr = append(cardArr, Cards{cardNum, winning, myCards})
	}
	return cardArr
}

func intersect(a1 []int, a2 []int) int{
	total := 0
	for _, item := range a2 {
		if slices.Contains(a1, item){
			total += 1
		}
	}
	return total
}

func tallyCards(cards []Cards, subset []Cards) int {
	total := 0
	for _, card := range subset {
		count := intersect(card.win, card.my)		
		total += tallyCards(cards, cards[card.num: card.num + count])
		total += 1
	}
	return total
}

func main(){
	cards := parseInput("input.txt")
	total := tallyCards(cards, cards)
	fmt.Println(total)
}
