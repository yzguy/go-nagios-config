package nagios

type HostEscalation struct {
	HostName             string `json:"host_name"`
	HostGroupName        string `json:"hostgroup_name"` // Maybe HostGroup.Name
	Contacts             string `json:"contacts"`       // Maybe []Contacts
	ContactGroups        string `json:"contact_groups"` // Maybe []ContactGroup
	FirstNotification    int    `json:"first_notification"`
	LastNotification     int    `json:"last_notification"`
	NotificationInterval int    `json:"notification_interval"`
	EscalationPeriod     string `json:"escalation_period"`  // Maybe TimePeriod.Name
	EscalationOptions    string `json:"escalation_options"` // maybe []string
}

func NewHostEscalation() (*HostEscalation, error) {
	return &HostEscalation{}, nil
}
