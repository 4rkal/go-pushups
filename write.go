package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lucasepe/codename"
)

func save(routine Routine) error {
	rng, err := codename.DefaultRNG()
	if err != nil {
		return fmt.Errorf("failed to get default RNG: %w", err)
	}

	name := fmt.Sprintf("%s.json", codename.Generate(rng, 0))

	jsonData, err := json.MarshalIndent(routine, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	file, err := os.Create(name)
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
