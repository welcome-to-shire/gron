.PHONY: test

PACKAGE=github.com/bcho/gron/pkg

test:
	go test $(PACKAGE)
