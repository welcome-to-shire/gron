package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/cron"
)

type Cronjob func()

// Start task manager and run the tasks.
func StartTasks(configs []TaskConfig) {
	c := cron.New()

	for _, taskConfig := range configs {
		job, err := makeCronJob(taskConfig)
		if err != nil {
			log.Fatal(err)
		}

		c.AddFunc(taskConfig.Schedule, job)
	}

	c.Start()
}

// Convert `TaskConfig` object into a cron job func.
// TODO use channel to expose state.
func makeCronJob(task TaskConfig) (Cronjob, error) {
	var err error

	cmdAndArgs := strings.Split(task.Command, " ")
	if len(cmdAndArgs) < 1 || cmdAndArgs[0] == "" {
		return nil, errors.New("cron job: unable to execute empty command")
	}
	return func() {
		cmd := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)

		// Prepare io.
		const (
			mode = os.O_WRONLY | os.O_APPEND | os.O_CREATE
			perm = 0600
		)
		cmd.Stdin, err = os.Open(task.Stdin)
		if err != nil {
			log.Printf("error when opening stdin: %q\n", err)
			return
		}
		cmd.Stdout, err = os.OpenFile(task.Stdout, mode, perm)
		if err != nil {
			log.Printf("error when opening stdout: %q\n", err)
			return
		}
		cmd.Stderr, err = os.OpenFile(task.Stderr, mode, perm)
		if err != nil {
			log.Printf("error when opening stderr: %q\n", err)
			return
		}

		log.Printf("running task: %s\n", task.Name)
		if err = cmd.Run(); err != nil {
			log.Printf(
				"error when running task %s: %q\n",
				task.Name,
				err,
			)
			return
		}
	}, nil
}
