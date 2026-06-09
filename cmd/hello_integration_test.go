//go:build integration

package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	hello := "./hello"
	buildCmd := exec.Command("go", "build", "-o", hello, "hello.go")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("No compile: %v", err)
	}
	defer func() {
		if err := os.Remove(hello); err != nil {
			log.Fatalf("Error in remove binary: %v", err)
		}
	}()
	mainCmd := exec.Command(hello)
	var stdout, stderr bytes.Buffer
	mainCmd.Stdout = &stdout
	mainCmd.Stderr = &stderr
	if err := mainCmd.Run(); err != nil {
		t.Fatalf("The program exit with error %v. Stderr: %s", err, stderr.String())
	}
	gotOutput := stdout.String()
	wantOutput := "Hello world!"
	if !strings.Contains(gotOutput, wantOutput) {
		t.Errorf("Expexted out %q, but got %q", wantOutput, gotOutput)
	}
}
