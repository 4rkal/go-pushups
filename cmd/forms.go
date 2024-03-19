package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

func greet() error {
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("go-pushups").
			Description("Welcome to _go-pushups_, your personal pushup companion and counter!")),
	).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		return err
	}
	return nil
}

func routineForm() (Routine, error) {
	var tmp string
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	form := huh.NewForm(
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
					routine.Reps = reps
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
					routine.Rest = rest
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
					routine.Increase = increase
					return nil
				}).Description("Percent increase per round")),
	).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		return Routine{}, err
	}
	return routine, nil
}

func saveForm() (bool, error) {
	var confirmation bool
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	confirmation_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("You will be doing %d push ups, with %d seconds rest and a %d increase per round. Do you want to save for future use?", routine.Reps, routine.Rest, routine.Increase)).
				Value(&confirmation).
				Affirmative("Yes save!").
				Negative("No thnx"),
		),
	).WithAccessible(accessible)

	err2 := confirmation_form.Run()
	if err2 != nil {
		return confirmation, err2
	}
	return confirmation, nil
}

func confirmationForm() (bool, error) {
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

func shouldRun() (bool, error) {
	var run bool
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))
	confirmation_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Routine Saved. Do you want to run it?").
				Value(&run).
				Affirmative("Yes run it!").
				Negative("No, quit"),
		),
	).WithAccessible(accessible)

	err := confirmation_form.Run()
	if err != nil {
		return run, err
	}
	return run, nil
}
