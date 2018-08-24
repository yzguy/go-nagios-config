package nagios

type Command struct {
	CommandName string
	CommandLine string
}

func NewCommand() (*Command, error) {
	return &Command{}, nil
}
