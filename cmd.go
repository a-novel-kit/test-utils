package testutils

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

// CMDResult is the result of a command execution by RunCMD.
type CMDResult struct {
	// Err is the error returned by exec.Command Run.
	Err error
	// Success is true if the process exited with a zero status.
	Success bool

	// STDOut is the captured standard output.
	STDOut string
	// STDErr is the captured standard error.
	STDErr string
}

// CMDConfig is used to run a command in a controlled environment.
type CMDConfig struct {
	// CmdFn is the function that will be executed under the separate process. Its result will then be passed to
	// MainFn.
	CmdFn func(t *testing.T)
	// MainFn is a controller function, that oversees the result of CmdFn. It is used as the actual test result.
	MainFn func(t *testing.T, res *CMDResult)
	// Env is a replacement env that will be available to the CmdFn.
	//
	// It is populated with os.Environ first. If present, the variable `JUST_CHECKING` will be filtered out.
	Env []string
}

// RunCMD setups a pattern to run test in controlled environments.
// https://stackoverflow.com/a/33404435/9021186
//
// IMPORTANT: this must be run only once per test.
func RunCMD(t *testing.T, config *CMDConfig) {
	t.Helper()

	if os.Getenv("JUST_CHECKING") == "bruh" {
		config.CmdFn(t)
		// Ensure proper exit anyway, otherwise we create a memory leak.
		os.Exit(0)
		return
	}

	// Filter out reserved env.
	env := append(append(os.Environ(), config.Env...), "JUST_CHECKING=bruh")

	outWriter, outCapture, err := CreateSTDCapture(t)
	require.NoError(t, err)
	errWriter, errCapture, err := CreateSTDCapture(t)
	require.NoError(t, err)

	cmd := exec.Command(os.Args[0], "-test.run="+t.Name()) //nolint:gosec
	cmd.Stdout = outWriter
	cmd.Stderr = errWriter
	cmd.Env = env

	err = cmd.Run()

	report := &CMDResult{
		Err:     err,
		Success: true,
		STDOut:  outCapture(),
		STDErr:  errCapture(),
	}

	if err != nil {
		var exErr *exec.ExitError
		if errors.As(err, &exErr) {
			report.Success = exErr.Success()
		} else {
			report.Success = false
		}
	}

	config.MainFn(t, report)
}
