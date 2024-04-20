package main

import (
	"fmt"
	"math/rand"
	"time"
)

var guess, count int

func game() {
	theNum := rand.Intn(100)
	for guess != theNum {
		fmt.Print("Try guessing : ")
		fmt.Scan(&guess)
		if guess > theNum {
			fmt.Println("Oops, the number is lower")
			count++
		} else if guess < theNum {
			fmt.Println("Uh Oh, the number is Higher")
			count++
		} else {
			fmt.Println("Correct, the number is ", theNum)
			fmt.Printf("You guess in %v tries\n", count)

		}
	}
}

func main() {
	var yn string
	isPlaying := true
	rand.NewSource(time.Now().UnixNano())

	for isPlaying {
		fmt.Println("Guessing game")
		game()
		fmt.Print("Want to try again ? (y/n)")
		fmt.Scan(&yn)
		if yn != "y" {
			return
		}
	}
}
