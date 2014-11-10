/*
Gron: cron job runner powered by golang.

Usage:

```bash
$ ./gron -config tasks.json
```

Tasks.json:

```json
{
  "tasks": [
    {
      "name": "ping-google.com",
      "schedule": "0 30 * * * *",
      "command": "ping -c 3 www.google.com",
      "stdout": "/var/tmp/ping-google.com.stdout",
      "stderr": "/var/tmp/ping-google.com.stderr"
    },

    {
      "name": "run-my-awesome-program",
      "schedule": "@every 1h30m",
      "command": "/path/to/my/awesome/prog",
      "stdin": "/path/to/some/awesome/input"
    }
  ],

  "reporters": [
    {
      "name": "my-fail-reporter",
      "options": {
        "reporter-options-field": "set something"
      }
    }
  ]
}
```

See README.md for comprehensive examples.
*/
package main
