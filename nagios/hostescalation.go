package nagios

type HostEscalation struct {
	HostName             string         `json:"host_name"`
	HostGroupName        []HostGroup    `json:"hostgroup_name"`
	Contacts             []Contact      `json:"contacts"`
	ContactGroups        []ContactGroup `json:"contact_groups"`
	FirstNotification    int            `json:"first_notification"`
	LastNotification     int            `json:"last_notification"`
	NotificationInterval int            `json:"notification_interval"`
	EscalationPeriod     TimePeriod     `json:"escalation_period"`
	EscalationOptions    string         `json:"escalation_options"` // maybe []string
}

func NewHostEscalation() (*HostEscalation, error) {
	return &HostEscalation{}, nil
}
