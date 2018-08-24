package nagios

type TimePeriod struct {
	Name       string            `json:"timeperiod_name"`
	Alias      string            `json:"alias"`
	Periods    map[string]string `json:"periods"`
	Exceptions map[string]string `json:"exceptions"`
	Exclude    []string          `json:"exclude"`
}

func NewTimePeriod() (*TimePeriod, error) {
	return &TimePeriod{}, nil
}
