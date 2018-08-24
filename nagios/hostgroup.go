package nagios

type HostGroup struct {
	Name             string
	Alias            string
	Members          string // Maybe []Host
	HostGroupMembers string // Maybe []HostGroup
	Notes            string
	NotesURL         string // maybe URL
	ActionURL        string // maybe URL
}
