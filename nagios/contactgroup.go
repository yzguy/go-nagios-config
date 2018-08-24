package nagios

type ContactGroup struct {
	Name                string
	Alias               string
	Members             string // Maybe []Contact
	ContactGroupMembers string // Maybe []ContactGroup
}

func NewContactGroup() (*ContactGroup, error) {
	return &ContactGroup{}, nil
}
