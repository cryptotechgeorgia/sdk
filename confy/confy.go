package confy

import (
	"encoding/json"
	"fmt"
	"os"
)

// parse json configuration file into out T type
func ParseConfiguration[T any](filename string, out *T) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&out); err != nil {
		return fmt.Errorf("failed to decode config file: %w", err)
	}

	return nil
}
