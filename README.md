# Gron

[![Build Status](https://travis-ci.org/welcome-to-shire/gron.svg?branch=master)](https://travis-ci.org/welcome-to-shire/gron) ![Current Version](http://img.shields.io/badge/version-0.0.1-brightgreen.svg)

Cron job runner powered by golang.

## Usage:

```bash
$ ./gron -config tasks.json
```

## Tasks.json:

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

### Task

Setting Fields:

| field name | description | required? |
|:----------|-----------|---------|
| name | task name | yes |
| schedule | task schedule format string, please refer [robfig/cron][robfig-cron] | yes |
| command | task command | yes |
| stdin | task running stdin | no, defaults to `/dev/null` |
| stdout | task running stdout | no, defaults to `/dev/null` |
| stderr | task running stderr | no, defaults to `/dev/null` |

[robfig-cron]: https://github.com/robfig/cron/blob/master/doc.go


### Reporter

When a task failed, reporter will receive a incident record and it can do something
on this issue (like warning the SA).

Setting Fields:

| field name | description | required? |
|:----------:|-------------|-----------|
| name | reporter's name | yes |
| options | reporter's options | no |

Supported Reporters:

- `log`: simply write the incident to log (using [log][golang-log] package).
- `palantir`: send the incident report to [palantir][shire-palantir]. See [palantir_reporter.go](palantir_reporter.go) for detailed setting options.


[golang-log]: http://golang.org/pkg/log/
[shire-palantir]: https://github.com/welcome-to-shire/palantir


## LICENSE

See [LICENSE.md](LICENSE.md).
