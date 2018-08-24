package nagios

type ContactGroup struct {
	Name                string         `json:"contactgroup_name"`
	Alias               string         `json:"alias"`
	Members             []Contact      `json:"members"`
	ContactGroupMembers []ContactGroup `json:"contactgroup_members"`
}

func NewContactGroup() (*ContactGroup, error) {
	return &ContactGroup{}, nil
}
