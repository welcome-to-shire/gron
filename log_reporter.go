package main

import (
	"encoding/json"
	"log"
)

type LogReporter struct{}

func makeLogReporter(options json.RawMessage) (*LogReporter, error) {
	return new(LogReporter), nil
}

func (r LogReporter) Report(incident Incident) {
	log.Printf("Task %s failed: %q", incident.Task.Name, incident.Err)
}
