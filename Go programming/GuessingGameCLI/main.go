package main

import (
	"bufio"
	"fmt"
	"math/rand/v2" // For random numbers
	"os"
	"strconv"
	"strings"
)

func main() {
	// Setup the reader
	reader := bufio.NewReader(os.Stdin)

	// Generate a number between 0 and 100
	secretNumber := rand.N(101) 

	fmt.Println("I have picked a number between 0 and 100.")
	fmt.Println("Can you guess it?")

	// Infinite loop until they guess correctly
	for {
		fmt.Print("Enter your guess: ")
		
		// Get Input
		//Read the whole line (up to the Enter key)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again")
			continue
		}
		// Clean the input: Windows adds "\r\n", Mac/Linux adds "\n". TrimSpace removes both

		input = strings.TrimSpace(input)
		// Convert ASCII to Integer (Atoi)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("⚠️ That doesn't look like a number. Please enter a valid integer.")
			continue // Skip the rest of the loop and start over
		}

		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("🎉 You got it! The number was", secretNumber)
			break // Exit the loop
		}
	}
}