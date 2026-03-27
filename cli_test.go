package cli

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// helper function to capture stdout
func captureStdout(f func()) string {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic("captureStdout: failed to create pipe: " + err.Error())
	}
	os.Stdout = w
	restored := false
	defer func() {
		if !restored {
			w.Close()
			os.Stdout = oldStdout
		}
	}()

	f()

	w.Close()
	os.Stdout = oldStdout
	restored = true

	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

// TestRegisterAndExecuteCommand tests registering a new command and its execution
func TestRegisterAndExecuteCommand(t *testing.T) {
	testCmdName := "test"
	testCmdOutput := "Test command executed\n"
	Command(testCmdName, func() {
		fmt.Println("Test command executed")
	}, "Test command")

	output := captureStdout(func() {
		executeCommand(testCmdName)
	})

	if output != testCmdOutput {
		t.Errorf("Expected command output to be %q, got %q", testCmdOutput, output)
	}
}

// TestListCommands tests the listing of registered commands
func TestListCommands(t *testing.T) {
	expectedOutputContains := "Available commands:"
	output := captureStdout(func() {
		listCommands()
	})

	if !bytes.Contains([]byte(output), []byte(expectedOutputContains)) {
		t.Errorf("Expected listCommands output to contain %q", expectedOutputContains)
	}
}

// TestUnknownCommand tests the handling of an unknown command
func TestUnknownCommand(t *testing.T) {
	unknownCmd := "helloworld"
	expectedOutput := fmt.Sprintln("Unknown command:", unknownCmd)
	output := captureStdout(func() {
		executeCommand(unknownCmd)
	})

	if output != expectedOutput {
		t.Errorf("Expected unknown command output to be %q, got %q", expectedOutput, output)
	}
}
