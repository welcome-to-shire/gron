package main

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Tasks     []TaskConfig
	Reporters []Reporter
}

// Parse commands config from json file.
func ParseFile(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return parseFromReader(file)
}

func parseFromReader(reader io.Reader) (*Config, error) {
	var raw struct {
		Tasks     []TaskConfig     `json:"tasks"`
		Reporters []ReporterConfig `json:"reporters"`
	}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&raw)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	for _, task := range raw.Tasks {
		task, err := prepareTask(task)
		if err != nil {
			return nil, err
		}

		config.Tasks = append(config.Tasks, task)
	}

	for _, reporterConfig := range raw.Reporters {
		reporter, err := prepareReporter(reporterConfig)
		if err != nil {
			return nil, err
		}

		config.Reporters = append(config.Reporters, reporter)
	}

	return config, nil
}
