package main

type Incident struct {
	Task TaskConfig
	Err  error
}
type Reporter interface {
	// Report an incident.
	Report(Incident)
}

func getReporter(config ReporterConfig) Reporter {
	switch config.Name {
	case "log":
		return makeLogReporter(config.Options)
	}

	return nil
}
