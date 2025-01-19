package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Score struct {
	Name	string
	Score	int
}

func main() {
	fmt.Println("Students")
	fmt.Println(strings.Repeat("-", 16))

	scores := []Score{
		{Name: "Cesar", Score: 99},
		{"Denise", 98},
		{"Jocelyn", 97},
	}

	for idx, score := range scores{
		fmt.Println(idx, " - ", score.Name)
		// fmt.Printf("%s %d\n", score.Name, score.Score)
	}

	fmt.Print("Select score to print: ")
	var option string
	fmt.Scanln(&option)

	opt, err := strconv.Atoi(option)
	if err != nil {
		fmt.Println("Invalid option")
	} else if opt >= 0 && opt < len(scores) {
		fmt.Println(scores[opt].Name, ": ", scores[opt].Score)
	} else {
		fmt.Println("Invalid option")
	}
}
