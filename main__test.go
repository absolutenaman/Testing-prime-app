package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test__isPrime(t *testing.T) {
	tests := []struct {
		name     string
		msg      string
		expected bool
		testNum  int
	}{
		{name: "IsPrime", msg: "The number entered 7 is indeed a prime number", testNum: 7, expected: true},
		{name: "isNotPrime", msg: "0 and 1 are not prime numbers", testNum: 1, expected: false},
		{name: "isNotPrime", msg: "Negative numbers are not prime numbers", testNum: -1, expected: false},
		{name: "isNotPrime", msg: "The number entered 20 is not a prime number as it got divided by 2", testNum: 20, expected: false},
	}
	for i := 0; i < len(tests); i++ {
		res, msg := isPrime(tests[i].testNum)
		if !res && tests[i].expected {
			t.Errorf("Expected %v but got %v", tests[i].expected, res)
		}
		if res && !tests[i].expected {
			t.Errorf("Expected %v but got %v", tests[i].expected, res)
		}
		if msg != tests[i].msg {
			t.Errorf("Expected message to be %s but got %s", tests[i].msg, msg)
		}
	}
}
func Test__prompt(t *testing.T) {

	// Save the original stdout (console)
	oldStdout := os.Stdout

	// Create a pipe (a secret tunnel)
	r, w, _ := os.Pipe()

	// Redirect all prints to the pipe
	os.Stdout = w

	// Call the function (it prints, but we catch it in the pipe)
	prompt()

	// Close writer so we can read what was written
	_ = w.Close()

	// Restore real console again
	os.Stdout = oldStdout

	// Read the captured output
	out, _ := io.ReadAll(r)

	// Show what was captured
	fmt.Println("Captured output:", string(out))

	if string(out) != "->" {
		t.Errorf("expected -> but got %s", string(out))
	}

}

func Test__intro(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldStdout
	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Hello User , It's nice to finally see you") {
		t.Errorf("expected `Hello User , It's nice to finally see you` but got %s", string(out))
	}
}

func Test__GoodBye(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	GoodBye()
	_ = w.Close()
	os.Stdout = oldStdout
	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Thanks for visiting") {
		t.Errorf("expected `Thanks for visiting` but got %s", string(out))
	}
}

func Test__checkNumbers(t *testing.T) {
	testData := []struct {
		name  string
		msg   string
		input string
		res   bool
	}{
		{name: "isPrime", msg: "The number entered 7 is indeed a prime number", input: "7", res: true},
		{name: "isNotPrime", msg: "The number entered 20 is not a prime number as it got divided by 2", input: "20", res: false},
		{name: "QUIT", msg: "", input: "q", res: false},
		{name: "isNotNumber", msg: "Please enter a whole number!", input: "hgjh", res: false},
	}

	for i := 0; i < len(testData); i++ {
		input := strings.NewReader(testData[i].input)
		reader := bufio.NewScanner(input)
		msg, res := checkNumbers(reader)
		if res != testData[i].res && msg != testData[i].msg {
			t.Errorf("Expected %s but got %s", testData[i].msg, msg)
		}
	}
}
func Test__readUserData(t *testing.T) {
	done := make(chan bool)
	var stdin bytes.Buffer
	stdin.Write([]byte("1\nq\n"))
	go readUserData(&stdin, done)
	<-done
}
