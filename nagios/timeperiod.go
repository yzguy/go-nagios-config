package nagios

type TimePeriod struct {
	Name       string
	Alias      string
	Periods    map[string]string
	Exceptions map[string]string
	Exclude    string // Maybe map[string]string
}
