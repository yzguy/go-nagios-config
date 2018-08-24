package nagios

type Host struct {
	Hostname                   string
	Alias                      string
	DisplayName                string
	Address                    string
	Parents                    string // Maybe []Host.Name
	HostGroups                 string // Maybe []HostGroups
	CheckCommand               string // Maybe Command.Name
	InitialState               string // Maybe []string
	MaxCheckAttempts           int
	CheckInterval              int
	RetryInterval              int
	ActiveChecksEnabled        int    // Maybe bool
	PassiveChecksEnabled       int    // Maybe bool
	CheckPeriod                string // Maybe TimePeriod.Name
	ObsessOverHost             int    // Maybe bool
	CheckFreshness             int    // Maybe bool
	FreshnessThreshold         int
	EventHandler               string // Maybe Command.Name
	EventHandlerEnabled        int    // Maybe bool
	LowFlapThreshold           int
	HighFlapThreshold          int
	FlapDetectionEnabled       int    // Maybe bool
	FlapDetectionOptions       string // Maybe []string
	ProcessPerfData            int    // Maybe bool
	RetainStatusInformation    int    // Maybe bool
	RetainNonStatusInformation int    // Maybe bool
	Contacts                   string // Maybe []Contacts
	ContactGroups              string // Maybe []ContactGroups
	NotificationInterval       int
	FirstNotificationDelay     int
	NotificationPeriod         string // Maybe TimePeriod.Name
	NotificationOptions        string // Maybe []string
	NotificationsEnabled       int    // Maybe bool
	StalkingOptions            string // Maybe []string
	Notes                      string // Maybe object?
	NotesURL                   string // maybe URL
	ActionURL                  string // maybe URL
	IconImage                  string
	IconImageAlt               string
	VRMLImage                  string
	StatusMapImage             string
	Coords2D                   string // Maybe Coordinates type?
	Coords3D                   string // Maybe Coordinates type?
}

func NewHost() (*Host, error) {
	return &Host{}, nil
}
