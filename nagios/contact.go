package nagios

type Contact struct {
	Name                        string
	Alias                       string
	ContactGroups               string // Maybe []ContactGroup
	HostNotificationsEnabled    int    // Maybe bool
	ServiceNotificationsEnabled int    // Maybe bool
	HostNotificationPeriod      string // Maybe TimePeriod.Name
	ServiceNotificationPeriod   string // Maybe TimePeriod.Name
	HostNotificationOptions     string // Maybe []string
	ServiceNotificationOptions  string // Maybe []string
	HostNotificationCommands    string // Maybe Command.Name
	ServiceNotificationCommands string // Maybe Command.Name
	Email                       string
	Pager                       string
	AddressX                    string
	CanSubmitCommands           int // Maybe bool
	RetainStatusInformation     int // Maybe bool
	RetainNonStatusInformation  int // Maybe bool
}

func NewContact() (*Contact, error) {
	return &Contact{}, nil
}
