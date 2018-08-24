package nagios

type ServiceEscalation struct {
	HostName             string `json:"host_name"`      // Maybe []Host
	HostGroupName        string `json:"hostgroup_name"` // Maybe []HostGroup
	ServiceDescription   string `json:"service_description"`
	Contacts             string `json:"contacts"`       // Maybe []Contact
	ContactGroups        string `json:"contact_groups"` // Maybe []ContactGroup
	FirstNotification    int    `json:"first_notification"`
	LastNotification     int    `json:"last_notification"`
	NotificationInterval int    `json:"notification_interval"`
	EscalationPeriod     string `json:"escalation_period"`  // Maybe TimePeriod.Name
	EscalationOptions    string `json:"escalation_options"` // Maybe []string
}

func NewServiceEscalation() (*ServiceEscalation, error) {
	return &ServiceEscalation{}, nil
}
