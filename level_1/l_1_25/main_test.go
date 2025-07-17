package main

import (
	"testing"
	"time"
)

var testCases = []struct {
	name     string
	duration time.Duration
}{
	{"100ms", 100 * time.Millisecond},
	{"500ms", 500 * time.Millisecond},
	{"1s", 1 * time.Second},
}

var testFuncs = []struct {
	name      string
	sleepFunc func(time.Duration)
}{
	{"V1", sleepV1},
	{"V2", sleepV2},
	{"V3", sleepV3},
}

func checkSleep(t *testing.T, sleepFunc func(time.Duration), duration time.Duration) {
	startedAt := time.Now()
	maxRuntime := time.Duration(float64(duration) * 1.01)

	done := make(chan struct{})
	go func() {
		sleepFunc(duration)
		close(done)
	}()

	select {
	case <-done:
		if time.Since(startedAt) < duration {
			t.Error("Waiting too little time")
		}

	case <-time.After(maxRuntime):
		t.Error("Waiting to long")

	}
}

func TestSleep(t *testing.T) {

	for _, tf := range testFuncs {
		t.Run(tf.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					checkSleep(t, tf.sleepFunc, tc.duration)
				})
			}

		})
	}
}
