package nagios

type Command struct {
	CommandName string `json:"command_name"`
	CommandLine string `json:"command_line"`
}

func NewCommand() (*Command, error) {
	return &Command{}, nil
}
