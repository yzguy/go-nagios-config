package nagios

type ContactGroup struct {
	Name                string `json:"contactgroup_name"`
	Alias               string `json:"alias"`
	Members             string `json:"members"`              // Maybe []Contact
	ContactGroupMembers string `json:"contactgroup_members"` // Maybe []ContactGroup
}

func NewContactGroup() (*ContactGroup, error) {
	return &ContactGroup{}, nil
}
