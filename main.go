package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inisialisasi generator angka acak dengan seed berdasarkan waktu
	rand.Seed(time.Now().UnixNano())

	// Generate angka acak antara 1 dan 100
	secretNumber := rand.Intn(100) + 1
	attempts := 0
	maxAttempts := 10

	fmt.Println("Welcome to Guess the Number!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Printf("You have %d attempts to guess it.\n", maxAttempts)

	for attempts < maxAttempts {
		fmt.Print("Enter your guess: ")
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		attempts++
		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You've guessed the number %d in %d attempts.\n", secretNumber, attempts)
			return
		}
	}

	fmt.Printf("Sorry, you've reached the maximum number of attempts. The secret number was %d.\n", secretNumber)
}
