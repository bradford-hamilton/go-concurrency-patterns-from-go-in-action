// Package main is an example package from Go In Action: https://www.manning.com/books/go-in-action.

// This sample program demonstrates how to use a channel to monitor the amount of
// time the program is running and terminate the program if it runs too long.
package main

import (
	"log"
	"os"
	"time"

	"github.com/bradford-hamilton/concurrency_patterns/runner"
)

// timeout is the number of seconds the program has to finish
const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	// Create a new timer value for this run.
	r := runner.New(timeout)

	// Add the tasks to be run.
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks and handle the result.
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
