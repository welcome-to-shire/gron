package gron

import "testing"

func makeTaskConfig(name, cmd, sched, stdin, stdout, stderr string) *TaskConfig {
	t := new(TaskConfig)
	t.Name = name
	t.Command = cmd
	t.Schedule = sched
	t.Stdin = stdin
	t.Stdout = stdout
	t.Stderr = stderr

	return t
}

func TestMakeCronJob(t *testing.T) {
	task := makeTaskConfig(
		"test",
		"echo 1",
		"@every 1h",
		"/dev/null",
		"/dev/null",
		"/dev/null",
	)

	job, err := makeCronJob(*task)
	if err != nil {
		t.Error(err)
	}
	if job == nil {
		t.Error("should return cron job func")
	}
}

func TestMakeCronJobWithInvalidCommand(t *testing.T) {
	task := makeTaskConfig(
		"test",
		"",
		"@every 1h",
		"/dev/null",
		"/dev/null",
		"/dev/null",
	)

	_, err := makeCronJob(*task)
	if err == nil {
		t.Error("should not accept empty command")
	}
}
