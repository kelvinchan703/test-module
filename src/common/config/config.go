package config

import (
	commonModels "ctint-conv/src/common/models"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var GlobalConfig commonModels.GlobalConfig

func InitGolobalConfig() {
	env := "dev" // Default to dev if not specified
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	// Read YAML file based on the environment
	filename := fmt.Sprintf("./ctint-global-config-%s.yaml", env)
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	// Unmarshal YAML data into struct
	if err := yaml.Unmarshal(data, &GlobalConfig); err != nil {
		log.Fatalf("error unmarshalling YAML data: %v", err)
	}

	// Access configuration values
	fmt.Printf("Load %s global config successfully\n", env)
}
