package main

import (
	"errors"
	"fmt"
)

type Incident struct {
	Task TaskConfig
	Err  error
}
type Reporter interface {
	// Report an incident.
	Report(Incident)
}

func getReporter(config ReporterConfig) (Reporter, error) {
	switch config.Name {
	case "log":
		return makeLogReporter(config.Options)
	case "palantir":
		return makePalantirReporter(config.Options)
	}

	return nil, errors.New(fmt.Sprintf("unable to find reporter: %s", config.Name))
}
