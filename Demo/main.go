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

	name, score := "Cesar", 99

	fmt.Println("Student scores")
	fmt.Println(strings.Repeat("-", 16))
	fmt.Println(name, score);
}
