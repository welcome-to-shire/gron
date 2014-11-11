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
	WaitIncident(config.Reporters, incidentCh)
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
