package main

import (
	"strings"
	"testing"
)

func makeReaderFromString(input string) *strings.Reader {
	return strings.NewReader(input)
}

func TestParseConfig(t *testing.T) {
	var task TaskConfig

	streamReader := makeReaderFromString(`{
                "tasks": [
                        {
                                "name": "foo",
                                "command": "echo foo",
                                "schedule": "@every 30m"
                        },
                        {
                                "name": "bar",
                                "command": "echo bar",
                                "schedule": "@every 1h30",
                                "stdin": "/var/tmp/stdin",
                                "stdout": "/var/tmp/stdout",
                                "stderr": "/var/tmp/stderr"
                        }
                ]
        }`)

	parsed, err := parseFromReader(streamReader)

	if err != nil {
		t.Error(err)
	}
	if parsed == nil {
		t.Fatal("should not be nil")
	}
	lenOfTasks := len(parsed.Tasks)
	if lenOfTasks != 2 {
		t.Errorf("should have 2 tasks, got: %d", lenOfTasks)
	}

	for _, task := range parsed.Tasks {
		if task.Name == "" {
			t.Errorf("task name is required")
		}
		if task.Schedule == "" {
			t.Errorf("task schedule time is required")
		}
		if task.Command == "" {
			t.Errorf("task is required")
		}
	}

	// Fields std* default are `/dev/null`.
	task = parsed.Tasks[0]
	if task.Stdin != "/dev/null" {
		t.Errorf("task's default stdin should be /dev/null")
	}
	if task.Stdout != "/dev/null" {
		t.Errorf("task's default stdout should be /dev/null")
	}
	if task.Stderr != "/dev/null" {
		t.Errorf("task's default stderr should be /dev/null")
	}

	task = parsed.Tasks[1]
	if task.Stdin != "/var/tmp/stdin" {
		t.Errorf("task's stdin should not be changed")
	}
	if task.Stdout != "/var/tmp/stdout" {
		t.Errorf("task's stdout should not be changed")
	}
	if task.Stderr != "/var/tmp/stderr" {
		t.Errorf("task's stderr should not be changed")
	}
}

func TestParseInvalidConfig(t *testing.T) {
	streamReader := makeReaderFromString(`{
                "tasks": [
                        {
                                "name": "missing-command"
                        }
                ]
        }`)

	_, err := parseFromReader(streamReader)
	if err == nil {
		t.Errorf("should return error")
	}
}
