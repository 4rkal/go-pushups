package main

import (
	"fmt"
	"os"
	"time"

	"github.com/martinlindhe/notify"
)

var amount_done int = 0

func do(reps, rest, increase int) int {
	time.Sleep(time.Duration(rest) * time.Second)
	if increase != 0 {
		reps += reps * increase / 100
	}
	amount_done += reps
	return reps
}

func alert(reps int) {
	message := fmt.Sprintf("Do %v pushups", reps)
	notify.Alert("go-pushups", "Pushup time!", message, "logo.png")
}

func main() {
	var reps int
	var rest int
	var increase int
	var confirm string

	fmt.Print("Enter amount of reps: ")
	fmt.Scan(&reps)

	fmt.Print("Enter amount of rest (sec): ")
	fmt.Scan(&rest)

	if reps <= 0 || rest <= 0 {
		fmt.Println("Not positive input")
		os.Exit(1)
	}

	fmt.Print("Enter percent increase per round: ")
	fmt.Scan(&increase)

	for round := 1; ; round++ {
		reps := do(reps, rest, increase)
		fmt.Printf("Round %d: Do %d pushups\n", round, reps)
		alert(reps)
		fmt.Print("Did you do them? (y/q)")
		fmt.Scan(&confirm)
		if confirm == "q" {
			fmt.Printf("You did %v pushups", amount_done)
			break
		}
	}
}
