package main

import (
	"testing"
)

func TestGetReporter(t *testing.T) {
	testCases := []struct {
		config ReporterConfig
		isNil  bool
	}{
		{ReporterConfig{"log", nil}, false},
		{ReporterConfig{"unknown-repoter", nil}, true},
	}

	for _, test := range testCases {
		testIsNil := getReporter(test.config) == nil
		if testIsNil != test.isNil {
			t.Errorf("expected %t, got %t", test.isNil, testIsNil)
		}
	}
}
