package nagios

type HostEscalation struct {
	Name                 string
	HostGroupName        string // Maybe HostGroup.Name
	Contacts             string // Maybe []Contacts
	ContactGroups        string // Maybe []ContactGroup
	FirstNotification    int
	LastNotification     int
	NotificationInterval int
	EscalationPeriod     string // Maybe TimePeriod.Name
	EscalationOptions    string // maybe []string
}

func NewHostEscalation() (*HostEscalation, error) {
	return &HostEscalation{}, nil
}
