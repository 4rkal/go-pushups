package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/martinlindhe/notify"
)

const (
	padding  = 2
	maxWidth = 80
)

type Routine struct {
	Reps     int `json:"reps"`
	Rest     int `json:"rest"`
	Increase int `json:"increase"`
}

var routine Routine

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

var amount_done int = 0
var reps int

func do(reps, increase int) int {
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

func run(should_save bool) error {
	var confirmation bool
	greet()
	routine, err := routineForm()
	if err != nil {
		fmt.Println("oh oh")
		os.Exit(1)
	}

	confirmation, err2 := saveForm()
	if err2 != nil {
		fmt.Println("oh oh")
		os.Exit(1)
	}
	if routine.Reps <= 0 || routine.Rest <= 0 {
		fmt.Println("Not positive input")
		os.Exit(1)
	}

	if confirmation == true && should_save == true {
		save(routine, "")
	}

	m := model{
		progress: progress.New(progress.WithDefaultGradient()),
	}

	reps = routine.Reps

	for round := 1; ; round++ {
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Oh no!", err)
			os.Exit(1)
		}
		reps = do(reps, routine.Increase)
		fmt.Printf("Round %d: Do %d pushups\n", round, reps)
		alert(reps)
		quit, err := confirmationForm()
		if err != nil {
			fmt.Println("oh oh %s", err)
			os.Exit(1)
		}
		if quit == false {
			fmt.Printf("You did %v pushups\n", amount_done)
			os.Exit(0)
		}
	}
}

func run2(routinee Routine) error {
	clearScreen()
	routine = routinee
	m := model{
		progress: progress.New(progress.WithDefaultGradient()),
	}

	reps = routine.Reps

	for round := 1; ; round++ {
		clearScreen()
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Oh no!", err)
			os.Exit(1)
		}
		reps = do(reps, routine.Increase)
		fmt.Printf("Round %d: Do %d pushups\n", round, reps)
		alert(reps)
		quit, err := confirmationForm()
		if err != nil {
			fmt.Println("oh oh %s", err)
			os.Exit(1)
		}
		if quit == false {
			fmt.Printf("You did %v pushups\n", amount_done)
			os.Exit(0)
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func newRoutine(name string) error {
	routine, err := routineForm()
	if err != nil {
		return err
	}
	save(routine, name)
	run, err := shouldRun()
	if err != nil {
		return err
	}
	if run == true {
		run2(routine)
	} else {
		os.Exit(0)
	}
	return nil
}
