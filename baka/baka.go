package main

import (
	"fmt"

	"example.com/try"
)

var shout string
var times int
var guess int
var numberPool int

func main() {

	//fmt.Scan(&shout, &times)
	//newshout := try.NewShout(shout, times)
	//message := try.Face(newshout)
	//fmt.Println(message)
	fmt.Println("Введи максимум:")
	fmt.Scan(&numberPool)
	fmt.Printf("Введи число от 1 до %v\n", numberPool)
	fmt.Scan(&guess)
	tries, result := try.GuessIt(guess, numberPool)
	if result == true {
		fmt.Printf("Угадал с %v попыток!", tries)
	} else {
		fmt.Printf("Ну вот... %v попыток - и все зря.", tries)
	}
}
