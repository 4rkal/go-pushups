package cmd

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
		} else if msg.String() == "enter" {
			// Get the selected item and check if successful
			selectedItem := m.list.SelectedItem()

			// Access the selected item data (assuming `item` struct)
			selectedRoutine := selectedItem.(item)

			// Load the routine details from the file based on the filename
			routine, err := loadRoutine(selectedRoutine.title)
			if err != nil {
				fmt.Println("Error loading routine:", err)
				return m, nil
			}
			run2(routine)

			// Handle the loaded routine data (e.g., display details, start workout)
			fmt.Printf("Selected routine: %s\n", selectedRoutine.title)
			fmt.Printf("Amount: %d, Rest: %d seconds, Increase: %d%%\n", routine.Reps, routine.Rest, routine.Increase)

			// You can implement further actions based on the routine here

			return m, nil
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
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("failed to get user config directory: %w", err)
	}

	appDir := "go-pushups"
	appDirPath := filepath.Join(configDir, appDir)
	var routine Routine
	items := []list.Item{}
	_, files := load(appDirPath)
	for i := range files {
		jsonFile, err := os.Open(files[i])
		if err != nil {
			fmt.Println("oh oh")
			os.Exit(1)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &routine)
		// fmt.Println(files[i])
		// fmt.Println(routine)
		r := fmt.Sprintf("Amount: %d, Rest %d (sec), Increase %d %%", routine.Reps, routine.Rest, routine.Increase)
		name := strings.TrimSuffix(strings.TrimPrefix(files[i], appDirPath+"/"), ".json")
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

func loadRoutine(file string) (Routine, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("failed to get user config directory: %w", err)
	}

	appDir := "go-pushups"
	var routine Routine

	// Append .json extension to the file name
	filePath := filepath.Join(configDir, appDir, file+".json")

	// Open the JSON file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return routine, fmt.Errorf("error opening file %s: %v", filePath, err)
	}
	defer jsonFile.Close()

	// Read the file contents
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return routine, fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	// Unmarshal JSON into Routine struct
	err = json.Unmarshal(byteValue, &routine)
	if err != nil {
		return routine, fmt.Errorf("error unmarshalling JSON from file %s: %v", filePath, err)
	}
	return routine, nil
}
