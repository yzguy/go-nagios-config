package nagios

type ServiceGroup struct {
	Name                string              `json:"servicegroup_name"`
	Alias               string              `json:"alias"`
	Members             map[string][]string `json:"members"`
	ServiceGroupMembers []ServiceGroup      `json:"servicegroup_members"` // Maybe []ServiceGroup
	Notes               string              `json:"notes"`
	NotesURL            string              `json:"notes_url"`  // maybe URL
	ActionURL           string              `json:"action_url"` // maybe URL
}

func NewServiceGroup() (*ServiceGroup, error) {
	return &ServiceGroup{}, nil
}
