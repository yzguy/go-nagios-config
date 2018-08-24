package nagios

type Host struct {
	Hostname                   string `json:"host_name"`
	Alias                      string `json:"alias"`
	DisplayName                string `json:"display_name"`
	Address                    string `json:"address"`
	Parents                    string `json:"parents"`       // Maybe []Host.Name
	HostGroups                 string `json:"hostgroups"`    // Maybe []HostGroup
	CheckCommand               string `json:"check_command"` // Maybe Command.Name
	InitialState               string `json:"initial_state"` // Maybe []string
	MaxCheckAttempts           int    `json:"max_check_attempts"`
	CheckInterval              int    `json:"check_interval"`
	RetryInterval              int    `json:"retry_interval"`
	ActiveChecksEnabled        int    `json:"active_checks_enabled"`  // Maybe bool
	PassiveChecksEnabled       int    `json:"passive_checks_enabled"` // Maybe bool
	CheckPeriod                string `json:"check_period"`           // Maybe TimePeriod.Name
	ObsessOverHost             int    `json:"obsess_over_host"`       // Maybe bool
	CheckFreshness             int    `json:"check_freshness"`        // Maybe bool
	FreshnessThreshold         int    `json:"freshness_threshold"`
	EventHandler               string `json:"event_handler"`         // Maybe Command.Name
	EventHandlerEnabled        int    `json:"event_handler_enabled"` // Maybe bool
	LowFlapThreshold           int    `json:"low_flap_threshold"`
	HighFlapThreshold          int    `json:"high_flap_threshold"`
	FlapDetectionEnabled       int    `json:"flap_detection_enabled"`        // Maybe bool
	FlapDetectionOptions       string `json:"flap_detection_options"`        // Maybe []string
	ProcessPerfData            int    `json:"process_perf_data"`             // Maybe bool
	RetainStatusInformation    int    `json:"retain_status_information"`     // Maybe bool
	RetainNonStatusInformation int    `json:"retain_non_status_information"` // Maybe bool
	Contacts                   string `json:"contacts"`                      // Maybe []Contact
	ContactGroups              string `json:"contact_groups"`                // Maybe []ContactGroup
	NotificationInterval       int    `json:"notification_interval"`
	FirstNotificationDelay     int    `json:"first_notification_delay"`
	NotificationPeriod         string `json:"notification_period"`   // Maybe TimePeriod.Name
	NotificationOptions        string `json:"notification_options"`  // Maybe []string
	NotificationsEnabled       int    `json:"notifications_enabled"` // Maybe bool
	StalkingOptions            string `json:"stalking_options"`      // Maybe []string
	Notes                      string `json:"notes"`                 // Maybe object?
	NotesURL                   string `json:"notes_url"`             // maybe URL
	ActionURL                  string `json:"action_url"`            // maybe URL
	IconImage                  string `json:"icon_image"`
	IconImageAlt               string `json:"icon_image_alt"`
	VRMLImage                  string `json:"vrml_image"`
	StatusMapImage             string `json:"status_map_image"`
	Coords2D                   string `json:"2d_coords"` // Maybe Coordinates type?
	Coords3D                   string `json:"3d_coords"` // Maybe Coordinates type?
}

func NewHost() (*Host, error) {
	return &Host{}, nil
}
