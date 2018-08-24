package nagios

type HostGroup struct {
	Name             string      `json:"hostgroup_name"`
	Alias            string      `json:"alias"`
	Members          []Host      `json:"members"`
	HostGroupMembers []HostGroup `json:"hostgroup_members"`
	Notes            string      `json:"notes"`
	NotesURL         string      `json:"notes_url"`  // maybe URL
	ActionURL        string      `json:"action_url"` // maybe URL
}

func NewHostGroup() (*HostGroup, error) {
	return &HostGroup{}, nil
}
