package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	randomNumber := rand.Intn(9) + 1

	fmt.Println("You get 5 Tries to guess a number between 1 to 10")

	sumOfTries := 0;
	specialTryUsed := false;

	for try := 0; try < 5; try++ {

		fmt.Println("Attempt ", try+1);


		guessStr,_ := reader.ReadString('\n')
		guess,_ := strconv.Atoi(strings.TrimSpace(guessStr))

		if guess < 1 || guess > 10 {
			fmt.Println("Invalid Guess!")
		}

		if guess == randomNumber {
			fmt.Println("Thats Correct! You Win!!")
			break;
		} else if  guess < randomNumber {
			fmt.Println("Guess is too low!")
		} else {
			fmt.Println("Guess is too high!")
		}

		sumOfTries += guess;
	}

	specialTry:
	fmt.Println("All Tries Finished!");

	if !specialTryUsed {
		averageGuess := sumOfTries/5;

		if averageGuess > randomNumber-2 && averageGuess < randomNumber+2 {
			fmt.Println("You were too close! We give you are Special Attempt : ");
			specialTryUsed = true;

			guessStr,_ := reader.ReadString('\n')
			guess,_ := strconv.Atoi(strings.TrimSpace(guessStr))

			if guess < 1 || guess > 10 {
				fmt.Println("Invalid Guess!")
			}

			if guess == randomNumber {
				fmt.Println("Thats Correct! You Win!!")
			} else if  guess < randomNumber {
				fmt.Println("Guess is too low!")
				goto specialTry
			} else {
				fmt.Println("Guess is too high!")
				goto specialTry
			}			
		}
	}

	fmt.Println("GameOver")
}