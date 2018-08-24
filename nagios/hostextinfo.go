package nagios

type HostExtInfo struct {
	HostName       string // Maybe Host.Name
	Notes          string
	NotesURL       string // Maybe URL
	ActionURL      string // Maybe URL
	IconImage      string
	IconImageAlt   string
	VRMLImage      string
	StatusMapImage string
	Coords2D       string // maybe coords type?
	Coords3D       string // maybe coords type?
}
