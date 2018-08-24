package nagios

type ServiceDependency struct {
	DependentHostName           string `json:"dependent_host_name"`         // Maybe []Host
	DependentHostGroupName      string `json:"dependent_hostgroup_name"`    // Maybe []HostGroup
	ServiceGroupName            string `json:"servicegroup_name"`           // Maybe []ServiceGroup
	DependentServiceGroupName   string `json:"dependent_servicegroup_name"` // Maybe []ServiceGroups
	DependentServiceDescription string `json:"dependent_service_description"`
	HostName                    string `json:"host_name"`      // Maybe []Host
	HostGroupName               string `json:"hostgroup_name"` // Maybe []HostGroup
	ServiceDescription          string `json:"service_description"`
	InheritsParent              int    `json:"inherits_parent"`              // Maybe bool
	ExecutionFailureCritera     string `json:"execution_failure_critera"`    // Maybe []string
	NotificationFailureCritera  string `json:"notification_failure_critera"` // Maybe []string
	DependencyPeriod            string `json:"dependency_period"`            // Maybe TimePeriod
}

func NewServiceDependency() (*ServiceDependency, error) {
	return &ServiceDependency{}, nil
}
