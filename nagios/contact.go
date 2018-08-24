package nagios

type Contact struct {
	Name                        string         `json:"contact_name"`
	Alias                       string         `json:"alias"`
	ContactGroups               []ContactGroup `json:"contactgroups"`
	HostNotificationsEnabled    bool           `json:"host_notifications_enabled"`
	ServiceNotificationsEnabled bool           `json:"service_notifications_enabled"`
	HostNotificationPeriod      TimePeriod     `json:"host_notification_period"`
	ServiceNotificationPeriod   TimePeriod     `json:"service_notification_period"`
	HostNotificationOptions     string         `json:"host_notification_options"`    // Maybe []string
	ServiceNotificationOptions  string         `json:"service_notification_options"` // Maybe []string
	HostNotificationCommands    []Command      `json:"host_notification_commands"`
	ServiceNotificationCommands []Command      `json:"service_notification_commands"`
	Email                       string         `json:"email"`
	Pager                       string         `json:"pager"`
	AddressX                    string         `json:"addressx"`
	CanSubmitCommands           bool           `json:"can_submit_commands"`
	RetainStatusInformation     bool           `json:"retain_status_information"`
	RetainNonStatusInformation  bool           `json:"retain_non_status_information"`
}

func NewContact() (*Contact, error) {
	return &Contact{}, nil
}
