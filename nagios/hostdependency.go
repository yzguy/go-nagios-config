package nagios

type HostDependency struct {
	DependentHostName          string // Maybe []Host
	DependentHostGroupName     string // Maybe []Hostgroup
	HostName                   string // Maybe []Host
	HostGroupName              string // Maybe []HostGroup
	InheritsParent             int    // Maybe bool
	ExecutionFailureCritera    string // Maybe []string
	NotificationFailureCritera string // Maybe []string
	DependencyPeriod           string // Maybe TimePeriod.Name
}

func NewHostDependency() (*HostDependency, error) {
	return &HostDependency{}, nil
}
