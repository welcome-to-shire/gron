package main

import (
	"encoding/json"
	"errors"
)

type Incident struct {
	Task TaskConfig
	Err  error
}

type Reporter interface {
	// Report an incident.
	Report(Incident)
}

type ReporterConfig struct {
	Name    string          `json:"name"`
	Options json.RawMessage `json:"options,omitempty"`
}

func WaitIncident(reporters []Reporter, incidentCh chan Incident) {
	for {
		select {
		case incident := <-incidentCh:
			for _, reporter := range reporters {
				go reporter.Report(incident)
			}
		}
	}
}

func prepareReporter(reporterConfig ReporterConfig) (Reporter, error) {
	return getReporter(reporterConfig)
}

func getReporter(config ReporterConfig) (Reporter, error) {
	switch config.Name {
	case "log":
		return makeLogReporter(config.Options)
	case "palantir":
		return makePalantirReporter(config.Options)
	}

	return nil, errors.New(`repoter: unable find reporter: ` + config.Name)
}
