package nagios

type Service struct {
	HostName                   string // Maybe []Host
	HostGroupName              string // Maybe []HostGroup
	ServiceDescription         string
	DisplayName                string
	ServiceGroups              string // Maybe []ServiceGroup
	IsVolatile                 int    // Maybe bool
	CheckCommand               string // Maybe Command.Name
	InitialState               string // Maybe []string
	MaxCheckAttempts           int
	CheckInterval              int
	RetryInterval              int
	ActiveChecksEnabled        int    // Maybe bool
	PassiveChecksEnabled       int    // Maybe bool
	CheckPeriod                string // Maybe TimePeriod.Name
	ObsessOverService          int    // Maybe bool
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
	NotificationInterval       int
	FirstNotificationDelay     int
	NotificationPeriod         string // Maybe TimePeriod.Name
	NotificationOptions        string // Maybe []string
	NotificationsEnabled       int    // Maybe bool
	Contacts                   string // Maybe []Contact
	ContactGroups              string // Maybe []ContactGroup
	StalkingOptions            string // Maybe []string
	Notes                      string // Maybe object?
	NotesURL                   string // maybe URL
	ActionURL                  string // maybe URL
	IconImage                  string
	IconImageAlt               string
}

func NewService() (*Service, error) {
	return &Service{}, nil
}
