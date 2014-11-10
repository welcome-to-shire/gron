package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

const DEFAULT_STDIO_FILE = "/dev/null"

type TaskConfig struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Command  string `json:"command"`
	Stdin    string `json:"stdin,omitempty"`
	Stdout   string `json:"stdout,omitempty"`
	Stderr   string `json:"stderr,omitempty"`
}
type ReporterConfig struct {
	Name    string          `json:"name"`
	Options json.RawMessage `json:"options,omitempty"`
}
type jsonConfig struct {
	Tasks     []TaskConfig     `json:"tasks"`
	Reporters []ReporterConfig `json:"reporters"`
}

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
	var raw jsonConfig
	config := new(Config)

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&raw)
	if err != nil {
		return nil, err
	}

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

func prepareTask(task TaskConfig) (TaskConfig, error) {
	// Check required fields.
	if task.Name == "" {
		return task, errors.New("task: should specify task name")
	}
	if task.Command == "" {
		return task, errors.New("task: should specify task command")
	}
	if task.Schedule == "" {
		return task, errors.New("task: should specify task schedule")
	}

	// Set default values for std*.
	if task.Stderr == "" {
		task.Stderr = DEFAULT_STDIO_FILE
	}
	if task.Stdin == "" {
		task.Stdin = DEFAULT_STDIO_FILE
	}
	if task.Stdout == "" {
		task.Stdout = DEFAULT_STDIO_FILE
	}

	return task, nil
}

func prepareReporter(reporterConfig ReporterConfig) (Reporter, error) {
	return getReporter(reporterConfig)
}
