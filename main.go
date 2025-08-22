package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//print a welcome message
	intro()
	//create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)
	//start a goroutine to read the user's input and and run the prime function
	go readUserData(os.Stdin, doneChan)
	//block until the done channel gets a value
	<-doneChan
	//close the channel
	//say Goodbye
	GoodBye()
}
func intro() {
	fmt.Println("Hello User , It's nice to finally see you")
	fmt.Println("Instructions !!!")
	fmt.Println("Please enter a number to check whether it's prime or not, to quit press q")
	prompt()
}
func prompt() {
	fmt.Printf("->")
}
func readUserData(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)
	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		fmt.Println(res)
		prompt()
	}
}
func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()

	// check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	numToCheck, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(int(numToCheck))

	return msg, false
}

func GoodBye() {
	fmt.Println("Thanks for visiting")
}
func isPrime(n int) (bool, string) {
	if n == 1 || n == 0 {
		return false, fmt.Sprintf("0 and 1 are not prime numbers")
	}
	if n < 0 {
		return false, fmt.Sprintf("Negative numbers are not prime numbers")
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("The number entered %d is not a prime number as it got divided by %d", n, i)
		}
	}
	return true, fmt.Sprintf("The number entered %d is indeed a prime number", n)
}
