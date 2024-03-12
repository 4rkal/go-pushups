package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Routine struct {
	reps     int
	rest     int
	increase int
}

var routine Routine

func form1(tmp string) (Routine, error) {
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("go-pushups").
			Description("Welcome to _go-pushups_, your personal pushup companion and counter!")),
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
				}).Description("Percent increase per round")),
	).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		return Routine{}, err
	}
	return routine, nil
}

func form2() (bool, error) {
	var confirmation bool
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	confirmation_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("You will be doing %d push ups, with %d seconds rest and a %d increase per round", routine.reps, routine.rest, routine.increase)).
				Value(&confirmation).
				Affirmative("Yes!").
				Negative("No."),
		),
	).WithAccessible(accessible)

	err2 := confirmation_form.Run()
	if err2 != nil {
		return confirmation, err2
	}
	return confirmation, nil
}

func form3() (bool, error) {
	var quit bool
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	confirmation_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Did you do them?").
				Value(&quit).
				Affirmative("Yes!").
				Negative("Quit"),
		),
	).WithAccessible(accessible)

	err2 := confirmation_form.Run()
	if err2 != nil {
		return quit, err2
	}
	return quit, nil
}
