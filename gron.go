package main

import (
	"flag"
	"log"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "tasks.json", "Tasks.json")
	flag.Parse()

	setupLogger()
	config := setupConfig(configFile)

	incidentCh := make(chan Incident)
	StartTasks(config.Tasks, incidentCh)

	// Listen on incident.
	for {
		select {
		case incident := <-incidentCh:
			for _, reporter := range config.Reporters {
				reporter.Report(incident)
			}
		}
	}
}

func setupLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("GRON ")
}

func setupConfig(configFile string) *Config {
	config, err := ParseFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
