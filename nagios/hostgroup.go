package nagios

type HostGroup struct {
	Name             string `json:"hostgroup_name"`
	Alias            string `json:"alias"`
	Members          string `json:"members"`           // Maybe []Host
	HostGroupMembers string `json:"hostgroup_members"` // Maybe []HostGroup
	Notes            string `json:"notes"`
	NotesURL         string `json:"notes_url"`  // maybe URL
	ActionURL        string `json:"action_url"` // maybe URL
}

func NewHostGroup() (*HostGroup, error) {
	return &HostGroup{}, nil
}
