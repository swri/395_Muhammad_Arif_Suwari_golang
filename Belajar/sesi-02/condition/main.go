package main

import "fmt"

func main() {
	println("Condition")

	var currentYear = 2014

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim =>", age)
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	var score = 0

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Switch with relational operators
	var skor = 9

	switch {
	case skor == 8:
		fmt.Println("perfect")
	case (skor < 8) && (skor >= 3):
		fmt.Println("not bad")
		fallthrough // Switch (fallthrough keyword)
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	var skore = 4

	if skore > 7 {
		switch skore {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if skore <= 5 {
			fmt.Println("not bad")
		} else if skore <= 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if skore == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
