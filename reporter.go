package main

type Reporter interface {
	// Report an incident.
	Report(err error)
}

func getReporter(config ReporterConfig) Reporter {
	switch config.Name {
	case "log":
		return makeLogReporter(config.Options)
	}

	return nil
}
