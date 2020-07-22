package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandom() int {
	now := time.Now()
	rand.Seed(now.Unix())
	return rand.Intn(100)
}

func main() {
	number := getRandom()
	var guess int

	fmt.Println("Guess the number!")
	for guess != number {
		fmt.Scanf("%d\n", &guess)
		if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("Too low")
		}
	}
	fmt.Println("Win")
}
