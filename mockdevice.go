package elmobd

import (
	"fmt"
	"strings"
	"time"
)

/*==============================================================================
 * External
 */

// MockResult represents the raw text output of running a raw command,
// including information used in debugging to show what input caused what
// error, how long the command took, etc.
type MockResult struct {
	input     string
	outputs   []string
	error     error
	writeTime time.Duration
	readTime  time.Duration
	totalTime time.Duration
}

// Failed checks if the result is successful or not
func (res *MockResult) Failed() bool {
	return res.error != nil
}

// GetError returns the results current error
func (res *MockResult) GetError() error {
	return res.error
}

// GetOutputs returns the outputs of the result
func (res *MockResult) GetOutputs() []string {
	return res.outputs
}

// FormatOverview formats a result as an overview of what command was run and
// how long it took.
func (res *MockResult) FormatOverview() string {
	lines := []string{
		"=======================================",
		" Mocked command \"%s\"",
		"=======================================",
	}

	return fmt.Sprintf(
		strings.Join(lines, "\n"),
		res.input,
	)
}

// MockDevice represent a mocked serial connection
type MockDevice struct {
}

// RunCommand mocks the given AT/OBD command by just returning a result for the
// mocked outputs set earlier.
func (dev *MockDevice) RunCommand(command string) RawResult {
	return &MockResult{
		input:     command,
		outputs:   mockOutputs(command),
		writeTime: 0,
		readTime:  0,
		totalTime: 0,
	}
}

/*==============================================================================
 * Internal
 */

func mockMode1Outputs(subcmd string) []string {
	if strings.HasPrefix(subcmd, "00") {
		// PIDs supported part 1
		// Support all of the ones in this library except "obd standard" command
		return []string{
			"41 00 1F FD 80 02",
		}
	} else if strings.HasPrefix(subcmd, "20") {
		// PIDs supported part 2
		return []string{
			"41 20 00 00 00 00", // None
		}
	} else if strings.HasPrefix(subcmd, "40") {
		// PIDs supported part 3
		return []string{
			"41 40 00 00 00 00", // None
		}
	} else if strings.HasPrefix(subcmd, "60") {
		// PIDs supported part 4
		return []string{
			"41 60 00 00 00 00", // None
		}
	} else if strings.HasPrefix(subcmd, "80") {
		// PIDs supported part 5
		return []string{
			"41 80 00 00 00 00", // None
		}
	} else if strings.HasPrefix(subcmd, "04") {
		// Engine Load
		return []string{
			"41 04 7F",
		}
	} else if strings.HasPrefix(subcmd, "05") {
		// Coolant Temperature
		return []string{
			"41 05 64",
		}
	} else if strings.HasPrefix(subcmd, "06") {
		// Short Term Fuel Trim Bank 1
		return []string{
			"41 06 64",
		}
	} else if strings.HasPrefix(subcmd, "07") {
		// Long Term Fuel Trim Bank 1
		return []string{
			"41 07 45",
		}
	} else if strings.HasPrefix(subcmd, "08") {
		// Short Term Fuel Trim Bank 2
		return []string{
			"41 08 66",
		}
	} else if strings.HasPrefix(subcmd, "09") {
		// Long Term Fuel Trim Bank 2
		return []string{
			"41 09 75",
		}
	} else if strings.HasPrefix(subcmd, "0A") {
		// Fuel Pressure
		return []string{
			"41 0A 80",
		}
	} else if strings.HasPrefix(subcmd, "0B") {
		// 	Intake manifold absolute pressure
		return []string{
			"41 0B 80",
		}
	} else if strings.HasPrefix(subcmd, "0C") {
		// RPM
		return []string{
			"41 0C 0F A0",
		}
	} else if strings.HasPrefix(subcmd, "0D") {
		// Speed
		return []string{
			"41 0D FF",
		}
	} else if strings.HasPrefix(subcmd, "0E") {
		// Timing advance
		return []string{
			"41 0E 80",
		}
	} else if strings.HasPrefix(subcmd, "10") {
		// MAF air flow rate
		return []string{
			"41 10 80 80",
		}
	} else if strings.HasPrefix(subcmd, "11") {
		// Throttle Position
		return []string{
			"41 11 80",
		}
	} else if strings.HasPrefix(subcmd, "1F") {
		// Runtime Since Engine Start
		return []string{
			"41 1F 30 A0",
		}
	} else {
		return []string{"NOT SUPPORTED"}
	}
}

func mockOutputs(cmd string) []string {
	if cmd == "ATSP0" {
		return []string{"OK"}
	} else if cmd == "AT@1" {
		return []string{"OBDII by elm329@gmail.com"}
	} else if strings.HasPrefix(cmd, "01") {
		return mockMode1Outputs(cmd[2:])
	}

	return []string{"NOT SUPPORTED"}
}
