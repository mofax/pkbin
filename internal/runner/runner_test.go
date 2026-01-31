package runner

import (
	"os"
	"runtime"
	"testing"
)

func TestRunScript_Success(t *testing.T) {
	exitCode, err := RunScript("echo 'test'")
	if err != nil {
		t.Fatalf("RunScript() failed: %v", err)
	}
	if exitCode != 0 {
		t.Errorf("Expected exit code 0, got %d", exitCode)
	}
}

func TestRunScript_Failure(t *testing.T) {
	// Use a command that will fail
	exitCode, err := RunScript("false")
	if err != nil {
		t.Fatalf("RunScript() should not return error for failed command: %v", err)
	}
	if exitCode == 0 {
		t.Error("Expected non-zero exit code for failed command")
	}
}

func TestRunScript_EmptyCommand(t *testing.T) {
	exitCode, err := RunScript("")
	if err == nil {
		t.Error("RunScript() should return error for empty command")
	}
	if exitCode != 1 {
		t.Errorf("Expected exit code 1 for empty command, got %d", exitCode)
	}
}

func TestRunScript_EnvironmentPassthrough(t *testing.T) {
	// Set a test environment variable
	testVar := "PKBIN_TEST_VAR"
	testValue := "test_value_12345"
	os.Setenv(testVar, testValue)
	defer os.Unsetenv(testVar)

	// Run a command that checks for the environment variable
	var checkCmd string
	if runtime.GOOS == "windows" {
		checkCmd = "cmd /c \"if %" + testVar + "%==test_value_12345 exit 0 else exit 1\""
	} else {
		checkCmd = "sh -c '[ \"$" + testVar + "\" = \"test_value_12345\" ]'"
	}

	exitCode, err := RunScript(checkCmd)
	if err != nil {
		t.Fatalf("RunScript() failed: %v", err)
	}
	if exitCode != 0 {
		t.Errorf("Environment variable not passed through correctly, exit code: %d", exitCode)
	}
}

func TestRunScript_WorkingDirectory(t *testing.T) {
	// Run a command that prints working directory
	var pwdCmd string
	if runtime.GOOS == "windows" {
		pwdCmd = "cd"
	} else {
		pwdCmd = "pwd"
	}

	exitCode, err := RunScript(pwdCmd)
	if err != nil {
		t.Fatalf("RunScript() failed: %v", err)
	}
	if exitCode != 0 {
		t.Errorf("Command failed with exit code: %d", exitCode)
	}
	// Note: We can't easily verify the output here since it goes to os.Stdout
	// But we can verify the command runs successfully
}
