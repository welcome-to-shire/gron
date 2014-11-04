package gron

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/cron"
)

type Cronjob func()

// Start task manager and run the tasks.
func Start(config TasksConfig) {
	c := cron.New()

	for _, taskConfig := range config.Tasks {
		job, err := makeCronJob(taskConfig)
		if err != nil {
			// TODO handle error like a gentleman
			panic(err)
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
		const create_and_append = os.O_WRONLY | os.O_APPEND | os.O_CREATE
		cmd.Stdin, err = os.Open(task.Stdin)
		if err != nil {
			return
		}
		cmd.Stdout, err = os.OpenFile(task.Stdout, create_and_append, 0600)
		if err != nil {
			return
		}
		cmd.Stderr, err = os.OpenFile(task.Stderr, create_and_append, 0600)
		if err != nil {
			return
		}

		cmd.Run()
	}, nil
}
