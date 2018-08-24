package nagios

type HostDependency struct {
	DependentHostName          string `json:"dependent_host_name"`          // Maybe []Host
	DependentHostGroupName     string `json:"dependent_hostgroup_name"`     // Maybe []Hostgroup
	HostName                   string `json:"host_name"`                    // Maybe []Host
	HostGroupName              string `json:"hostgroup_name"`               // Maybe []HostGroup
	InheritsParent             int    `json:"inherits_parent"`              // Maybe bool
	ExecutionFailureCritera    string `json:"execution_failure_critera"`    // Maybe []string
	NotificationFailureCritera string `json:"notification_failure_critera"` // Maybe []string
	DependencyPeriod           string `json:"dependency_period"`            // Maybe TimePeriod.Name
}

func NewHostDependency() (*HostDependency, error) {
	return &HostDependency{}, nil
}
