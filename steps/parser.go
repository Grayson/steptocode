package steps

import (
	"os"

	"github.com/bitrise-io/stepman/models"
	"gopkg.in/yaml.v2"
)

func ParseYamlFile(path string) (*models.StepModel, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseYaml(bytes)
}

func ParseYaml(bytes []byte) (*models.StepModel, error) {
	var sm models.StepModel
	if err := yaml.Unmarshal(bytes, &sm); err != nil {
		return nil, err
	}
	return &sm, nil
}
