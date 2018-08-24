package nagios

type ServiceExtInfo struct {
	HostName           string // Maybe Host.Name
	ServiceDescription string
	Notes              string
	NotesURL           string // Maybe URL
	ActionURL          string // Maybe URL
	IconImage          string
	IconImageAlt       string
}

func NewServiceExtInfo() (*ServiceExtInfo, error) {
	return &ServiceExtInfo{}, nil
}
