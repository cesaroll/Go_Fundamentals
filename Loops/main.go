package main

import (
	"fmt"
	"strconv"
)

func main() {

	type Score struct {
		Name	string
		Score	int
	}

	scores := []Score{}

	for {
		fmt.Print("Enter student name and score: ")
		var name, rawScore string
		fmt.Scanln(&name, &rawScore)
		score, err := strconv.Atoi(rawScore)
		if err != nil {
			fmt.Println("Invalid score")
			continue
		}
		fmt.Println(name, score)
		scores = append(scores, Score{Name: name, Score: score})
	}

}
