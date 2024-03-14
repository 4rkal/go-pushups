package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucasepe/codename"
)

func save(routine Routine) error {
	rng, err := codename.DefaultRNG()
	if err != nil {
		return fmt.Errorf("failed to get default RNG: %w", err)
	}

	// Get the user's config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config directory: %w", err)
	}

	appDir := "go-pushups"
	appDirPath := filepath.Join(configDir, appDir)

	if err := os.MkdirAll(appDirPath, 0755); err != nil {
		return fmt.Errorf("error creating app directory: %w", err)
	}

	filename := fmt.Sprintf("%s.json", codename.Generate(rng, 0))
	filePath := filepath.Join(appDirPath, filename)

	jsonData, err := json.MarshalIndent(routine, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("error closing file: %v\n", closeErr)
		}
	}()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing JSON to file: %w", err)
	}

	return nil
}
