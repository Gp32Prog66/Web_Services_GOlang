package main

/*
Movie Script from Genius.com
https://genius.com/Monty-python-monty-python-and-the-holy-grail-bridge-of-death-annotated
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Type anything to continue ")

	//Test Reading Only One Input
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	fmt.Println(input)

	//Testing String Import Library
	str1 := "Testing String"

	res1 := strings.ToLower(str1)

	fmt.Println(res1)

	//Declaring Variables
	//var name string

	name := ""
	nameStr := strings.ToLower(name)

	fmt.Print("What is your name? ")
	nameInput, _ := inputReader.ReadString('\n')
	fmt.Println(nameInput)

	//var quest string

	quest := ""
	questStr := strings.ToLower(quest)

	fmt.Print("What is your quest? ")
	questInput, _ := inputReader.ReadString('\n')
	fmt.Println(questInput)

	//var color string
	color := ""
	colorStr := strings.ToLower(color)

	fmt.Print("What is your favorite color? ")
	colorInput, _ := inputReader.ReadString('\n')
	fmt.Println(colorInput)

	//Strings

	//var capital string
	capital := ""
	capitalStr := strings.ToLower(capital)

	fmt.Println("What is the capital of Bahrain? ")
	fmt.Scanln(&capitalStr)

	if capitalStr == "Manama" {
		fmt.Println("Right. Off you go.")
	} else {
		fmt.Println("Auuuuuuuuuuuuuuuuugh")
	}


/*

	//var finalQuestion string
	finalQuestion := ""
	finalQuestionStr := strings.ToLower(finalQuestion)

	fmt.Print("Final Question ")
	finalQuestionInput, _ := inputReader.ReadString('\n')
	fmt.Println(finalQuestionInput)

	if finalQuestionInput == "What do you mean? An African or European swallow?" {
		fmt.Println("What? I don't know that! Auuuuuuuuuuuuuuuuugh")
	} else {
		fmt.Println("Auuuuuuuuuuuuuuuuugh")
	}

	*/
	//Questions
	fmt.Println("What is your name? ")
	fmt.Scanln(&nameStr)

	//fmt.Println("What is your quest? ")
	fmt.Scanln(&questStr)
	fmt.Println("What is your favorite color? ")
	fmt.Scanln(&colorStr)

	//Off guard questions
	//fmt.Println("What is the capital of Bahrain?")
	//fmt.Scanln(&capitalStr)


	/*
	//Determine Correct Answer for Bahrain Capital
	if capital == "Manama" {
		fmt.Println("Right. Off you go.")
	} else {
		fmt.Println("Auuuuuuuuuuuuuuuuugh")
	}
*/
	//Final Questions
	//fmt.Println("What is the air-speed velocity of an unladen swallow? ")
	//fmt.Scanln(&finalQuestionStr)

/*

	//Answer
	if finalQuestion == "What do you mean? An African or European swallow?" {
		fmt.Println("What? I don't know that! Auuuuuuuuuuuuuuuuugh")
	} else {
		fmt.Println("Auuuuuuuuuuugh")
	}
*/
}
