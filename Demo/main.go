package main

import (
	"fmt"
	"strings"
)

func main() {
	// var name string = "Cesar";
	// name := "Cesar";

	// var score = 99;
	// score := 99;

	// name, score := "Cesar", 99

	students := []string{
		"Cesar",
		"Denise",
		"Jocelyn",
	}

	scores := []int{99,93,96}

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 16))
	fmt.Println(students[0], scores[0]);
	fmt.Println(students[1], scores[1]);
	fmt.Println(students[2], scores[2]);
}
