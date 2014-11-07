package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	configFile := flag.String("config", "tasks.json", "Tasks.json")
	flag.Parse()

	if *configFile == "" {
		showUsage()
		os.Exit(1)
	}
	tasksConfig, err := ParseFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	Start(*tasksConfig)

	for {
		time.Sleep(120 * time.Minute)
	}
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "Usage: \n")
	flag.PrintDefaults()
}

func setupLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
