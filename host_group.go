package zabbix

// 主机组的来源。  取值范围: 0—普通主机组; 4—自动发现的主机组。
// For `HostGroupObject` field: `Status`
const (
	HostGroupFlagsPlain       = 0
	HostGroupFlagsDiscrovered = 4
)

// 组是否被系统内部使用。内部组不能被删除。  取值范围: 0 - *(默认值)* 非内部; 1 - 内部。
// For `HostGroupObject` field: `Internal`
const (
	HostGroupInternalFalse = 0
	HostGroupInternalTrue  = 1
)

// HostGroupObject 定义主机组
type HostGroupObject struct {
	//主机组ID。
	GroupID int `json:"groupid,omitempty"`
	//主机组的名称。
	Name string `json:"name,omitempty"`
	//主机组的来源。  取值范围: 0—普通主机组; 4—自动发现的主机组。
	Flags int `json:"flags,omitempty"` // has defined consts, see above
	//组是否被系统内部使用。内部组不能被删除。  取值范围: 0 - *(默认值)* 非内部; 1 - 内部。
	Internal int `json:"internal,omitempty"` // has defined consts, see above
	//主机列表
	Hosts     []HostObject     `json:"hosts,omitempty"`
	Templates []TemplateObject `json:"templates,omitempty"`
}

// HostGroupGetParams struct is used for hostGroup get requests
type HostGroupGetParams struct {
	GetParameters

	GraphIDs       []int `json:"graphids,omitempty"`
	GroupIDs       []int `json:"groupids,omitempty"`
	HostIDs        []int `json:"hostids,omitempty"`
	MaintenanceIDs []int `json:"maintenanceids,omitempty"`
	MonitoredHosts bool  `json:"monitored_hosts,omitempty"`
	RealHosts      bool  `json:"real_hosts,omitempty"`
	TemplatedHosts bool  `json:"templated_hosts,omitempty"`
	TemplateIDs    []int `json:"templateids,omitempty"`
	TriggerIDs     []int `json:"triggerids,omitempty"`

	WithApplications              bool `json:"with_applications,omitempty"`
	WithGraphs                    bool `json:"with_graphs,omitempty"`
	WithGraphPrototypes           bool `json:"with_graph_prototypes,omitempty"`
	WithHostsAndTemplates         bool `json:"with_hosts_and_templates,omitempty"`
	WithHttptests                 bool `json:"with_httptests,omitempty"`
	WithItems                     bool `json:"with_items,omitempty"`
	WithItemPrototypes            bool `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes bool `json:"with_simple_graph_item_prototypes,omitempty"`
	WithMonitoredHttptests        bool `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            bool `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers         bool `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems          bool `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  bool `json:"with_triggers,omitempty"`

	// SelectDiscoveryRule  SelectQuery `json:"selectDiscoveryRule,omitempty"` // not implemented yet
	// SelectGroupDiscovery SelectQuery `json:"selectGroupDiscovery,omitempty"` // not implemented yet
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`
}

// Structure to store creation result
type hostGroupCreateResult struct {
	GroupIDs []int `json:"groupids"`
}

// Structure to store deletion result
type hostGroupDeleteResult struct {
	GroupIDs []int `json:"groupids"`
}

// HostGroupGet gets hostGroups
func (z *Context) HostGroupGet(params HostGroupGetParams) ([]HostGroupObject, int, error) {

	var result []HostGroupObject

	status, err := z.request("hostgroup.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostGroupCreate creates hostGroups
func (z *Context) HostGroupCreate(params []HostGroupObject) ([]int, int, error) {

	var result hostGroupCreateResult

	status, err := z.request("hostgroup.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GroupIDs, status, nil
}

// HostGroupDelete deletes hostGroups
func (z *Context) HostGroupDelete(groupIDs []int) ([]int, int, error) {

	var result hostGroupDeleteResult

	status, err := z.request("hostgroup.delete", groupIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GroupIDs, status, nil
}

// HostGroupUpdate update hostGroups
func (z *Context) HostGroupUpdate(params []HostGroupObject) ([]int, int, error) {

	var result hostGroupCreateResult

	status, err := z.request("hostgroup.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GroupIDs, status, nil
}
