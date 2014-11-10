package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", "tasks.json", "Tasks.json")
	flag.Parse()

	setupLogger()
	config := setupConfig(configFile)
	StartTasks(config.Tasks)

	for {
		time.Sleep(120 * time.Minute)
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
