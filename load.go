package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model2 struct {
	list list.Model
}

func (m model2) Init() tea.Cmd {
	return nil
}

func (m model2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model2) View() string {
	return docStyle.Render(m.list.View())
}

func load(rootDir string) (error, []string) {
	files := []string{}
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".json") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return err, files
	}
	return nil, files
}

func show_files() {
	var routine Routine
	items := []list.Item{}
	_, files := load(".")
	for i := range files {
		jsonFile, err := os.Open(files[i])
		if err != nil {
			fmt.Println("oh oh")
			os.Exit(1)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &routine)
		fmt.Println(files[i])
		fmt.Println(routine)
		r := fmt.Sprintf("Amount: %d, Rest %d (sec), Increase %d %%", routine.Reps, routine.Rest, routine.Increase)
		name := strings.TrimSuffix(files[i], ".json")
		newItem := item{title: name, desc: r}
		items = append(items, newItem)
	}

	m := model2{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Routines"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
