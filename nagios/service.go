package nagios

type Service struct {
	HostName                   []Host         `json:"host_name"`
	HostGroupName              []HostGroup    `json:"hostgroup_name"`
	ServiceDescription         string         `json:"service_description"`
	DisplayName                string         `json:"display_name"`
	ServiceGroups              []ServiceGroup `json:"servicegroups"`
	IsVolatile                 bool           `json:"is_volatile"`
	CheckCommand               Command        `json:"check_command"`
	InitialState               string         `json:"initial_state"` // Maybe []string
	MaxCheckAttempts           int            `json:"max_check_attempts"`
	CheckInterval              int            `json:"check_interval"`
	RetryInterval              int            `json:"retry_interval"`
	ActiveChecksEnabled        bool           `json:"active_checks_enabled"`
	PassiveChecksEnabled       bool           `json:"passive_checks_enabled"`
	CheckPeriod                TimePeriod     `json:"check_period"`
	ObsessOverService          bool           `json:"obsess_over_service"`
	CheckFreshness             bool           `json:"check_freshness"`
	FreshnessThreshold         int            `json:"freshness_threshold"`
	EventHandler               Command        `json:"event_handler"`
	EventHandlerEnabled        bool           `json:"event_handler_enabled"`
	LowFlapThreshold           int            `json:"low_flap_threshold"`
	HighFlapThreshold          int            `json:"high_flap_threshold"`
	FlapDetectionEnabled       bool           `json:"flap_detection_enabled"`
	FlapDetectionOptions       string         `json:"flap_detection_options"` // Maybe []string
	ProcessPerfData            bool           `json:"process_perf_data"`
	RetainStatusInformation    bool           `json:"retain_status_information"`
	RetainNonStatusInformation bool           `json:"retain_non_status_information"`
	NotificationInterval       int            `json:"notification_interval"`
	FirstNotificationDelay     int            `json:"first_notification_delay"`
	NotificationPeriod         TimePeriod     `json:"notification_period"`
	NotificationOptions        string         `json:"notification_options"` // Maybe []string
	NotificationsEnabled       bool           `json:"notifications_enabled"`
	Contacts                   []Contact      `json:"contacts"`
	ContactGroups              []ContactGroup `json:"contact_groups"`
	StalkingOptions            string         `json:"stalking_options"` // Maybe []string
	Notes                      string         `json:"notes"`
	NotesURL                   string         `json:"notes_url"`  // maybe URL
	ActionURL                  string         `json:"action_url"` // maybe URL
	IconImage                  string         `json:"icon_image"`
	IconImageAlt               string         `json:"icon_image_alt"`
}

func NewService() (*Service, error) {
	return &Service{}, nil
}
