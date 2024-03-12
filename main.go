package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/martinlindhe/notify"
)

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

var amount_done int = 0

type Routine struct {
	reps     int
	rest     int
	increase int
}

var routine Routine

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
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	confirmation_form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("You will be doing %d push ups, with %d seconds rest and a %d increase per round", routine.reps, routine.rest, routine.increase)).
				Value(&confirmation).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err2 := confirmation_form.Run()
	if err2 != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	if routine.reps <= 0 || routine.rest <= 0 {
		fmt.Println("Not positive input")
		os.Exit(1)
	}

	if confirmation == false {
		fmt.Println("Did not confirm")
		os.Exit(0)
	}
	m := model{
		progress: progress.New(progress.WithDefaultGradient()),
	}

	for round := 1; ; round++ {
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Println("Oh no!", err)
			os.Exit(1)
		}
		reps := do(routine.reps, routine.rest, routine.increase)
		fmt.Printf("Round %d: Do %d pushups\n", round, reps)
		alert(reps)
		confirm()
	}
}

type tickMsg time.Time

type model struct {
	progress progress.Model
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		fmt.Print("You did %d", amount_done)
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		cmd := m.progress.IncrPercent(0.1)
		return m, tea.Batch(tickCmd(), cmd)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + "Resting...\n\n" +
		pad + m.progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}

func tickCmd() tea.Cmd {
	rest := time.Duration(routine.rest/10) * time.Second
	return tea.Tick(rest, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
