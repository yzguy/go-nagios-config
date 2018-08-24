package nagios_test

import (
	"testing"

	"github.com/yzguy/go-nagios-config/nagios"
)

func TestNewCommand(t *testing.T) {
	var tests = []struct {
		CommandName string
		CommandLine string
	}{
		{
			"check_uptime",
			"check_uptime -H node1.test.local",
		},
		{
			"check_empty",
			"",
		},
	}

	for _, tt := range tests {
		cmd, err := nagios.NewCommand(tt.CommandName, tt.CommandLine)
		if err != nil {
			if err.Error() != nagios.ErrCommandArgsEmpty.Error() {
				t.Errorf("Command: Error received did not match expected (Got: %s, Expected: %s)", err.Error(), nagios.ErrCommandArgsEmpty.Error())
			}
		} else {
			if cmd.CommandName != tt.CommandName {
				t.Errorf("Command: CommandName did not match expected (Got: %s, Expected: %s)", cmd.CommandName, tt.CommandName)
			}

			if cmd.CommandLine != tt.CommandLine {
				t.Errorf("Command: CommandLine did not match expected (Got: %s, Expected: %s)", cmd.CommandLine, tt.CommandLine)
			}
		}
	}
}
