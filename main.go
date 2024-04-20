package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess int
	count := 0
	rand.NewSource(time.Now().UnixNano())

	theNum := rand.Intn(100)

	fmt.Println("Guessing game")
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
			fmt.Printf("You guess in %v tries", count)
		}
	}
}
