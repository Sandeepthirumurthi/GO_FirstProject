package main

import "fmt"

func main() {
	// data type
	var textment string = "Hello World"
	var num1 int = 100
	var num2 float64 = 2.6
	var num3 uint = 100 // accept only positive number.
	var boolean bool = true

	fmt.Println(textment)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(boolean)

	// another way of declare a variable
	text := "Sandeep"
	age1 := 27

	fmt.Println(text)
	fmt.Println(age1)

	// fmt.Printf("Hello %v, you are %v old", text, age) // %v --> Act like a placeholder. Here, "name" variable will be printed in the place %v.

	// quiz game
	fmt.Println("Welcome to my quiz game!")
	fmt.Printf("Enter your name: ")
	var name string

	fmt.Scan(&name) // "Scan" use to collect the user data and store in an variable. "&name" whatever the user typed in. It'll be store in name variable.

	fmt.Printf("Hello, %v, welcome to the game!", name)

	fmt.Printf("\n Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You cannot play")
		return
	}

	score := 0
	num_question := 2

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var ans1 string
	var ans2 string
	fmt.Scan(&ans1, &ans2)

	if ans1+" "+ans2 == "RTX 3090" || ans1+" "+ans2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v, out of %v \n", score, num_question)
	percent := (float64(score) / float64(num_question)) * 100

	fmt.Printf("You scored: %v%%.", percent)
}
