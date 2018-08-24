package nagios

type Host struct {
	Hostname                   string         `json:"host_name"`
	Alias                      string         `json:"alias"`
	DisplayName                string         `json:"display_name"`
	Address                    string         `json:"address"`
	Parents                    []Host         `json:"parents"`
	HostGroups                 []HostGroup    `json:"hostgroups"`
	CheckCommand               Command        `json:"check_command"`
	InitialState               string         `json:"initial_state"` // Maybe []string
	MaxCheckAttempts           int            `json:"max_check_attempts"`
	CheckInterval              int            `json:"check_interval"`
	RetryInterval              int            `json:"retry_interval"`
	ActiveChecksEnabled        bool           `json:"active_checks_enabled"`
	PassiveChecksEnabled       bool           `json:"passive_checks_enabled"`
	CheckPeriod                TimePeriod     `json:"check_period"` // Maybe TimePeriod.Name
	ObsessOverHost             bool           `json:"obsess_over_host"`
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
	Contacts                   []Contact      `json:"contacts"`
	ContactGroups              []ContactGroup `json:"contact_groups"`
	NotificationInterval       int            `json:"notification_interval"`
	FirstNotificationDelay     int            `json:"first_notification_delay"`
	NotificationPeriod         TimePeriod     `json:"notification_period"`
	NotificationOptions        string         `json:"notification_options"`  // Maybe []string
	NotificationsEnabled       bool           `json:"notifications_enabled"` // Maybe bool
	StalkingOptions            string         `json:"stalking_options"`      // Maybe []string
	Notes                      string         `json:"notes"`                 // Maybe object?
	NotesURL                   string         `json:"notes_url"`             // maybe URL
	ActionURL                  string         `json:"action_url"`            // maybe URL
	IconImage                  string         `json:"icon_image"`
	IconImageAlt               string         `json:"icon_image_alt"`
	VRMLImage                  string         `json:"vrml_image"`
	StatusMapImage             string         `json:"status_map_image"`
	Coords2D                   string         `json:"2d_coords"` // Maybe Coordinates type?
	Coords3D                   string         `json:"3d_coords"` // Maybe Coordinates type?
}

func NewHost() (*Host, error) {
	return &Host{}, nil
}
