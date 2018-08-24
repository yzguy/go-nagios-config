package nagios

type HostExtInfo struct {
	HostName       string `json:"host_name"` // Maybe Host.Name
	Notes          string `json:"notes"`
	NotesURL       string `json:"notes_url"`  // Maybe URL
	ActionURL      string `json:"action_url"` // Maybe URL
	IconImage      string `json:"icon_image"`
	IconImageAlt   string `json:"icon_image_alt"`
	VRMLImage      string `json:"vrml_image"`
	StatusMapImage string `json:"statusmap_image"`
	Coords2D       string `json:"2d_coords"` // maybe coords type?
	Coords3D       string `json:"3d_coords"` // maybe coords type?
}

func NewHostExtInfo() (*HostExtInfo, error) {
	return &HostExtInfo{}, nil
}
