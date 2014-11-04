.PHONY: test

PACKAGE=github.com/bcho/gron/pkg

test:
	go test $(PACKAGE)

install:
	go get github.com/robfig/cron
