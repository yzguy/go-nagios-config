package nagios

type Contact struct {
	Name                        string `json:"contact_name"`
	Alias                       string `json:"alias"`
	ContactGroups               string `json:"contactgroups"`                 // Maybe []ContactGroup
	HostNotificationsEnabled    int    `json:"host_notifications_enabled"`    // Maybe bool
	ServiceNotificationsEnabled int    `json:"service_notifications_enabled"` // Maybe bool
	HostNotificationPeriod      string `json:"host_notification_period"`      // Maybe TimePeriod.Name
	ServiceNotificationPeriod   string `json:"service_notification_period"`   // Maybe TimePeriod.Name
	HostNotificationOptions     string `json:"host_notification_options"`     // Maybe []string
	ServiceNotificationOptions  string `json:"service_notification_options"`  // Maybe []string
	HostNotificationCommands    string `json:"host_notification_commands"`    // Maybe Command.Name
	ServiceNotificationCommands string `json:"service_notification_commands"` // Maybe Command.Name
	Email                       string `json:"email"`
	Pager                       string `json:"pager"`
	AddressX                    string `json:"addressx"`
	CanSubmitCommands           int    `json:"can_submit_commands"`           // Maybe bool
	RetainStatusInformation     int    `json:"retain_status_information"`     // Maybe bool
	RetainNonStatusInformation  int    `json:"retain_non_status_information"` // Maybe bool
}

func NewContact() (*Contact, error) {
	return &Contact{}, nil
}
