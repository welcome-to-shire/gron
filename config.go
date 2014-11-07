package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

const DEFAULT_STDIO_FILE = "/dev/null"

// Commands.json schema
type TaskConfig struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Command  string `json:"command"`
	Stdin    string `json:"stdin,omitempty"`
	Stdout   string `json:"stdout,omitempty"`
	Stderr   string `json:"stderr,omitempty"`
}
type TasksConfig struct {
	Tasks []TaskConfig `json:"tasks"`
}

// Parse commands config from json file.
func ParseFile(filepath string) (*TasksConfig, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return parseFromReader(file)
}

func parseFromReader(reader io.Reader) (*TasksConfig, error) {
	tasks := new(TasksConfig)

	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	for i, task := range tasks.Tasks {
		tasks.Tasks[i], err = prepareTask(task)
		if err != nil {
			return nil, err
		}
	}
	return tasks, nil
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
