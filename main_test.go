package main

import (
	"strings"
	"testing"
	"time"
)

func TestWorkerOne(t *testing.T) {
	c1 := make(chan string)
	go workerOne(c1)

	select {
	case msg := <-c1:
		if !strings.Contains(msg, "worker one") {
			t.Errorf("Expected 'worker one' in message, got %s", msg)
		}
	case <-time.After(time.Second):
		t.Error("Timeout waiting for message from worker one")
	}
}

func TestWorkerTwo(t *testing.T) {
	c2 := make(chan string)
	go workerTwo(c2)

	select {
	case msg := <-c2:
		if !strings.Contains(msg, "worker two") {
			t.Errorf("Expected 'worker two' in message, got %s", msg)
		}
	case <-time.After(3 * time.Second):
		t.Error("Timeout waiting for message from worker two")
	}
}
