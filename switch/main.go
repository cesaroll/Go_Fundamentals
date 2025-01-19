package main

func main() {

	switch i := 2; i {
		case 1:
			println("One")
		case 2:
			println("Two")
		case 3:
			println("Three")
		default:
			println("Default")
	}
}
