package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	palantir "github.com/welcome-to-shire/palantir-go"
)

// Palantir setting options:
//
// - host: palantir server's host, string like 127.0.0.1:5566
// - subject: report subject
type PalantirReporter struct {
	Host    string `json:"host"`
	Subject string `json:"subject"`
	Client  *palantir.Client
}

func makePalantirReporter(raw json.RawMessage) (*PalantirReporter, error) {
	reporter := new(PalantirReporter)
	err := json.Unmarshal(raw, reporter)
	if err != nil {
		return nil, err
	}

	reporter.Client = palantir.MakeClient(reporter.Host)

	return reporter, nil
}

func (r PalantirReporter) Report(incident Incident) {
	message := palantir.Message{
		fmt.Sprintf("Gron: task %s failed!", incident.Task.Name),
		incident.Err.Error(),
		palantir.ISO8601Time(time.Now()),
	}

	_, err := r.Client.CreateMessage(r.Subject, message)
	if err != nil {
		log.Printf("failed to send message to palantir: %q", err)
	}
}
