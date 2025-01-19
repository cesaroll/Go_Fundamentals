package main

import (
	"fmt"
	"strconv"
)

type Score struct {
	Name	string
	Score	int
}

func main() {

	scores := []Score{}
	shouldContinue := true

	for shouldContinue {

		fmt.Println("(e) Enter a student\n(p) Print all students\n(q) Quit")
		var option string
		fmt.Scanln(&option)

		switch option {
			case "e":
				score, ok := addScore()
				if ok {
					scores = append(scores, score)
				}

			case "p":
				print(scores)

			case "q":
				shouldContinue = false
			default:
				shouldContinue = false
		}
		fmt.Println()
	}
}

func addScore() (score Score, ok bool){
	fmt.Print("\nEnter student name and score: ")
	var name, rawScore string
	fmt.Scanln(&name, &rawScore)
	s, err := strconv.Atoi(rawScore)
	if err != nil {
		fmt.Println("Invalid score")
		return Score{}, false
	}

	return Score{Name: name, Score: s}, true
}

func print(scores []Score) {
	fmt.Println()
	for _, score := range scores {
		fmt.Println(score.Name, " - ", score.Score)
	}
}
