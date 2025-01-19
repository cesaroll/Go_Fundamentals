package main

import (
	"fmt"
	"strings"
)

type Score struct {
	Name	string
	Score	int
}

func main() {
	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 16))

	scores := []Score{
		{Name: "Cesar", Score: 99},
		{"Denise", 98},
		{"Jocelyn", 97},
	}

	for _, score := range scores{
		fmt.Println(score.Name, score.Score)
		// fmt.Printf("%s %d\n", score.Name, score.Score)
	}
}
