package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	gron "github.com/bcho/gron/pkg"
)

func main() {
	configFile := flag.String("config", "", "Tasks.json")
	flag.Parse()

	if *configFile == "" {
		showUsage()
		os.Exit(1)
	}
	tasksConfig, err := gron.ParseFile(*configFile)
	if err != nil {
		// TODO handle error like a gentleman
		panic(err)
	}
	gron.Start(*tasksConfig)

	for {
		time.Sleep(120 * time.Minute)
	}
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "Usage: \n")
	flag.PrintDefaults()
}
