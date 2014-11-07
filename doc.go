/*
Gron: cron job runner powered by golang.

Usage:

```bash
$ ./gron -config tasks.json
```

Tasks.json:

```json
{
  "tasks": {
    "ping-google.com": {
      "schedule": "0 30 * * * *",
      "command": "ping -c 3 www.google.com",
      "stdout": "/var/tmp/ping-google.com.stdout",
      "stderr": "/var/tmp/ping-google.com.stderr"
    },

    "run-my-awesome-program": {
      "schedule": "@every 1h30m",
      "command": "/path/to/my/awesome/prog",
      "stdin": "/path/to/some/awesome/input"
    }
  }
}

See README.md for comprehensive examples.
```
*/
package main
