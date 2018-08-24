package nagios

type ServiceGroup struct {
	Name                string
	Alias               string
	Members             string // Maybe []Service
	ServiceGroupMembers string // Maybe []ServiceGroup
	Notes               string
	NotesURL            string // maybe URL
	ActionURL           string // maybe URL
}

func NewServiceGroup() (*ServiceGroup, error) {
	return &ServiceGroup{}, nil
}
