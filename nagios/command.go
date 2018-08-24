package nagios

import "errors"

var (
	ErrCommandArgsEmpty = errors.New("Command: command_name and/or command_line cannot be empty")
)

type Command struct {
	CommandName string `json:"command_name"`
	CommandLine string `json:"command_line"`
}

func NewCommand(commandName, commandLine string) (*Command, error) {
	if commandName == "" || commandLine == "" {
		return &Command{}, ErrCommandArgsEmpty
	}

	return &Command{
		CommandName: commandName,
		CommandLine: commandLine,
	}, nil
}
