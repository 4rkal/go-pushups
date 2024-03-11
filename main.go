package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
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
	var tmp string

	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("go-pushups").
			Description("Welcome to _go-pushups_")),
		huh.NewGroup(
			huh.NewInput().
				Value(&tmp).
				Title("Enter amount of reps:").
				Placeholder("reps").
				Validate(func(s string) error {
					reps, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("probably not int input")
					}
					routine.reps = reps
					return nil
				}).
				Description("To do in each round/cycle"),

			huh.NewInput().
				Value(&tmp).
				Title("Enter amount of rest:").
				Placeholder("rest (sec)").
				Validate(func(s string) error {
					rest, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("probably not int input")
					}
					routine.rest = rest
					return nil
				}).Description("Amount of rest in seconds"),
			huh.NewInput().
				Value(&tmp).
				Title("Enter percent increase per round:").
				Placeholder("percent increase").
				Validate(func(s string) error {
					increase, err := strconv.Atoi(s)
					if err != nil {
						return errors.New("probably not int input")
					}
					routine.increase = increase
					return nil
				}).Description("Percent increase per round"),
		),
	).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	if routine.reps <= 0 || routine.rest <= 0 {
		fmt.Println("Not positive input")
		os.Exit(1)
	}

	for round := 1; ; round++ {
		reps := do(routine.reps, routine.rest, routine.increase)
		fmt.Printf("Round %d: Do %d pushups\n", round, reps)
		alert(reps)
		confirm()
	}
}
