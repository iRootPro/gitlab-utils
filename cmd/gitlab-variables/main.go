package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config представляет структуру для хранения переменных
type Config struct {
	Variables map[string]string `yaml:"variables"`
}

func extractVariables(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return config.Variables, nil
}

func main() {
	allVariables := make(map[string]string)

	// Список файлов для обработки
	files := []string{".gitlab-ci.yml", ".gitlab-ci-usr.yml"}

	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			fmt.Printf("Extracting from %s\n", file)
			variables, err := extractVariables(file)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", file, err)
				continue
			}
			// Слияние переменных
			for k, v := range variables {
				allVariables[k] = v
			}
		}
	}

	// Output variable to stdout
	for key, value := range allVariables {
		fmt.Printf("%s=%s\n", key, value)
	}
}
