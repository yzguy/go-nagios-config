package nagios

type ServiceDependency struct {
	DependentHostName           string // Maybe []Host
	DependentHostGroup          string // Maybe []HostGroup
	ServiceGroupName            string // Maybe []ServiceGroup
	DependentServiceGroupName   string // Maybe []ServiceGroups
	DependentServiceDescription string
	HostName                    string // Maybe []Host
	HostGroupName               string // Maybe []HostGroup
	ServiceDescription          string
	InheritsParent              int    // Maybe bool
	ExecutionFailureCritera     string // Maybe []string
	NotificationFailureCritera  string // Maybe []string
	DependencyPeriod            string // Maybe TimePeriod
}
