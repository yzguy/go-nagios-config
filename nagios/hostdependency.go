package nagios

type HostDependency struct {
	DependentHostName          []Host      `json:"dependent_host_name"`
	DependentHostGroupName     []HostGroup `json:"dependent_hostgroup_name"`
	HostName                   []Host      `json:"host_name"`
	HostGroupName              []HostGroup `json:"hostgroup_name"`
	InheritsParent             bool        `json:"inherits_parent"`
	ExecutionFailureCriteria   string      `json:"execution_failure_criteria"`
	NotificationFailureCritera string      `json:"notification_failure_critera"`
	DependencyPeriod           TimePeriod  `json:"dependency_period"`
}

func NewHostDependency() (*HostDependency, error) {
	return &HostDependency{}, nil
}
