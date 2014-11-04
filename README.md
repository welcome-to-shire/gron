# Gron

Cron job runner powered by golang.

## Usage:

```bash
$ ./gron -config commands.json
```

### Commands.json:

```json
{
  "commands": {
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
```

Time schedule format please refer [robfig/cron][robfig-cron].

[robfig-cron]: https://github.com/robfig/cron/blob/master/doc.go


## LICENSE

See [LICENSE.md](LICENSE.md).
