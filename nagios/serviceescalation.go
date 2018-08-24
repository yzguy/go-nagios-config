package nagios

type ServiceEscalation struct {
	HostName             string // Maybe []Host
	HostGroupName        string // Maybe []HostGroup
	ServiceDescription   string
	Contacts             string // Maybe []Contact
	ContactGroups        string // Maybe []ContactGroup
	FirstNotification    int
	LastNotification     int
	NotificationInterval int
	EscalationPeriod     string // Maybe TimePeriod.Name
	EscalationOptions    string // Maybe []string
}
