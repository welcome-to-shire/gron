package main

import "encoding/json"

type LogReporter struct{}

func makeLogReporter(options json.RawMessage) *LogReporter {
	return new(LogReporter)
}

func (r LogReporter) Report(err error) {}
