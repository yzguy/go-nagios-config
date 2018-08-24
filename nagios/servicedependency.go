package nagios

type ServiceDependency struct {
	DependentHostName           []Host         `json:"dependent_host_name"`
	DependentHostGroupName      []HostGroup    `json:"dependent_hostgroup_name"`
	ServiceGroupName            []ServiceGroup `json:"servicegroup_name"`
	DependentServiceGroupName   []ServiceGroup `json:"dependent_servicegroup_name"`
	DependentServiceDescription string         `json:"dependent_service_description"`
	HostName                    []Host         `json:"host_name"`
	HostGroupName               []HostGroup    `json:"hostgroup_name"`
	ServiceDescription          string         `json:"service_description"`
	InheritsParent              bool           `json:"inherits_parent"`
	ExecutionFailureCriteria    string         `json:"execution_failure_criteria"`   // Maybe []string
	NotificationFailureCritera  string         `json:"notification_failure_critera"` // Maybe []string
	DependencyPeriod            TimePeriod     `json:"dependency_period"`
}

func NewServiceDependency() (*ServiceDependency, error) {
	return &ServiceDependency{}, nil
}
