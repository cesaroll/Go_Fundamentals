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
	shouldContinue := true

	for shouldContinue {

		fmt.Println("(e) Enter a student\n(p) Print all students\n(q) Quit")
		var option string
		fmt.Scanln(&option)

		switch option {
			case "e":
				fmt.Print("\nEnter student name and score: ")
				var name, rawScore string
				fmt.Scanln(&name, &rawScore)
				score, err := strconv.Atoi(rawScore)
				if err != nil {
					fmt.Println("Invalid score")
					continue
				}
				scores = append(scores, Score{Name: name, Score: score})

			case "p":
				fmt.Println()
				for _, score := range scores {
					fmt.Println(score.Name, " - ", score.Score)
				}

			case "q":
				shouldContinue = false
			default:
				shouldContinue = false
		}
		fmt.Println()
	}
}
