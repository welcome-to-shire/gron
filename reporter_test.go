package main

import (
	"encoding/json"
	"testing"
)

func TestGetReporter(t *testing.T) {
	testCases := []struct {
		config        ReporterConfig
		expectedError bool
	}{
		{ReporterConfig{"log", nil}, false},
		{
			ReporterConfig{
				"palantir",
				json.RawMessage(`{"host": "hostname", "subject": "subjectname"}`),
			},
			false,
		},
		{ReporterConfig{"unknown-repoter", nil}, true},
	}

	for _, test := range testCases {
		_, err := getReporter(test.config)
		if err != nil && !test.expectedError {
			t.Errorf("failed to create repoter")
		}
		if err == nil && test.expectedError {
			t.Errorf("should not create reporter")
		}
	}
}
