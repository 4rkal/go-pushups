package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/martinlindhe/notify"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

var amount_done int = 0

func do(reps, rest, increase int) int {
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
	var tmp string
	var confirmation bool
	routine, err := form1(tmp)
	if err != nil {
		fmt.Println("oh oh")
		os.Exit(1)
	}

	confirmation, err2 := form2()
	if err2 != nil {
		fmt.Println("oh oh")
		os.Exit(1)
	}
	if routine.Reps <= 0 || routine.Rest <= 0 {
		fmt.Println("Not positive input")
		os.Exit(1)
	}

	if confirmation == true {
		save(routine)
	}

	m := model{
		progress: progress.New(progress.WithDefaultGradient()),
	}

	for round := 1; ; round++ {
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Oh no!", err)
			os.Exit(1)
		}
		reps := do(routine.Reps, routine.Rest, routine.Increase)
		fmt.Printf("Round %d: Do %d pushups\n", round, reps)
		alert(reps)
		quit, err := form3()
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
