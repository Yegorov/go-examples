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

func TestHelloIntegration(t *testing.T) {
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

func BenchmarkHelloCmd(b *testing.B) {
	hello := "./hello"
	buildCmd := exec.Command("go", "build", "-o", hello, "hello.go")
	if err := buildCmd.Run(); err != nil {
		b.Fatalf("No compile: %v", err)
	}
	defer func() {
		if err := os.Remove(hello); err != nil {
			log.Fatalf("Error in remove binary: %v", err)
		}
	}()
	// b.ResetTimer()
	// for i := 0; i < b.N; i++ {
	// 	_ = exec.Command(hello)
	// }
	b.ResetTimer()
	for b.Loop() {
		_ = exec.Command(hello)
	}
}
