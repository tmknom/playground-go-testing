package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	shutdown()
	os.Exit(exitCode)
}

func setup() {
	println("before all testing...")
}

func shutdown() {
	println("after all testing...")
}

func TestFoo(t *testing.T) {}
func TestBar(t *testing.T) { t.Fatal("fail example") }
func TestBaz(t *testing.T) {}
