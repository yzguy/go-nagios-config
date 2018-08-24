package nagios

type Object struct {
	Name     string `json:"name"`
	Use      string `json:"use"`
	Register string `json:"register"`
}
