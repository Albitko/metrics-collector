package contracts_settings

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Albitko/metrics-collector/internal/entity"
)

func New() (entity.ContractsSettings, error) {
	cfg := entity.ContractsSettings{}

	yamlFile, err := os.ReadFile("contracts_settings.yaml")
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
