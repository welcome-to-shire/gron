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
func StartTasks(configs []TaskConfig, incidentCh chan Incident) {
	c := cron.New()

	for _, taskConfig := range configs {
		job, err := makeCronJob(taskConfig, incidentCh)
		if err != nil {
			log.Fatal(err)
		}

		c.AddFunc(taskConfig.Schedule, job)
	}

	c.Start()
}

// Convert `TaskConfig` object into a cron job func.
func makeCronJob(task TaskConfig, incidentCh chan Incident) (Cronjob, error) {
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
			incidentCh <- Incident{task, err}
			return
		}
		cmd.Stdout, err = os.OpenFile(task.Stdout, mode, perm)
		if err != nil {
			incidentCh <- Incident{task, err}
			return
		}
		cmd.Stderr, err = os.OpenFile(task.Stderr, mode, perm)
		if err != nil {
			incidentCh <- Incident{task, err}
			return
		}

		log.Printf("running task: %s\n", task.Name)
		if err = cmd.Run(); err != nil {
			incidentCh <- Incident{task, err}
			return
		}
	}, nil
}
