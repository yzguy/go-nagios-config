package nagios

type ServiceEscalation struct {
	HostName             []Host         `json:"host_name"`
	HostGroupName        []HostGroup    `json:"hostgroup_name"`
	ServiceDescription   string         `json:"service_description"`
	Contacts             []Contact      `json:"contacts"`
	ContactGroups        []ContactGroup `json:"contact_groups"`
	FirstNotification    int            `json:"first_notification"`
	LastNotification     int            `json:"last_notification"`
	NotificationInterval int            `json:"notification_interval"`
	EscalationPeriod     TimePeriod     `json:"escalation_period"`
	EscalationOptions    string         `json:"escalation_options"` // Maybe []string
}

func NewServiceEscalation() (*ServiceEscalation, error) {
	return &ServiceEscalation{}, nil
}
