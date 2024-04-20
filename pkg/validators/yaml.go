package validators

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	configFilePath = "./test.yaml"
	//configFilePath      = "/etc/sc-controller/config.yaml"
)

type Process struct {
	Id               int
	Name             string `yaml:"name"`
	Path             string `yaml:"path"`
	WorkingDirectory string `yaml:"working_directory"`
	AutoStart        bool   `yaml:"auto_start"`
	AutoRestart      bool   `yaml:"auto_restart"`
}

type GlobalConfig struct {
	RestartTimes int `yaml:"restart_times"`
}

type Config struct {
	Global    GlobalConfig `yaml:"global"`
	Processes []Process    `yaml:"processes"`
}

type Validator struct{}

func Init() *Validator {
	return &Validator{}
}

func (v *Validator) Validate() (*Config, error) {
	f, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	data := Config{}

	err = yaml.Unmarshal(f, &data)
	if err != nil {
		return nil, err
	}

	if data.Global.RestartTimes == 0 {
		data.Global.RestartTimes = 3
	}

	for i, p := range data.Processes {
		if strings.TrimSpace(p.Name) == "" {
			return nil, fmt.Errorf("config file error: [process number %d] for each process the name attribute is required", i+1)
		}
		if strings.TrimSpace(p.Path) == "" {
			return nil, fmt.Errorf("config file error: [process number %d] for each process the path attribute is required", i+1)
		}
		if strings.TrimSpace(p.WorkingDirectory) == "" {
			p.WorkingDirectory = "."
		}
	}

	return &data, nil
}
