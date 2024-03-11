package main

import (
	"fmt"
	"os"
	"time"

	"github.com/martinlindhe/notify"
)

var amount_done int = 0

type Routine struct {
	reps     int
	rest     int
	increase int
}

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
	notify.Alert("go-pushups", "Pushup time!", message, "assets/logo.png")
}

func confirm() {
	confirm := ""
	for confirm != "y" {
		if confirm == "q" {
			fmt.Printf("You did %v pushups\n", amount_done)
			os.Exit(0)
		}
		fmt.Print("Did you do them? (y/q)")
		fmt.Scan(&confirm)
	}
}

func main() {
	var routine Routine

	fmt.Print("Enter amount of reps: ")
	fmt.Scan(&routine.reps)

	fmt.Print("Enter amount of rest (sec): ")
	fmt.Scan(&routine.rest)

	if routine.reps <= 0 || routine.rest <= 0 {
		fmt.Println("Not positive input")
		os.Exit(1)
	}

	fmt.Print("Enter percent increase per round: ")
	fmt.Scan(&routine.increase)

	for round := 1; ; round++ {
		reps := do(routine.reps, routine.rest, routine.increase)
		fmt.Printf("Round %d: Do %d pushups\n", round, reps)
		alert(reps)
		confirm()
	}
}
