package nagios

type ServiceExtInfo struct {
	HostName           HostName `json:"host_name"` // Maybe Host.Name
	ServiceDescription string   `json:"service_description"`
	Notes              string   `json:"notes"`
	NotesURL           string   `json:"notes_url"`  // Maybe URL
	ActionURL          string   `json:"action_url"` // Maybe URL
	IconImage          string   `json:"icon_image"`
	IconImageAlt       string   `json:"icon_image_alt"`
}

func NewServiceExtInfo() (*ServiceExtInfo, error) {
	return &ServiceExtInfo{}, nil
}
