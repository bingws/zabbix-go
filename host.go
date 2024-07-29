package zabbix

// Zabbix agent的可用性。  可能的值为： 0 - （默认）未知 1 - 可用； 2 - 不可用。
// For `HostObject` field: `Available`
const (
	HostAvailableUnknown     = 0
	HostAvailableAvailable   = 1
	HostAvailableUnavailable = 2
)

// 主机的来源。  可能的值： 0 - 普通主机； 4 - 自动发现的主机。
// For `HostObject` field: `Flags`
const (
	HostFlagsPlain      = 0
	HostFlagsDiscovered = 4
)

// 主机的资产清单填充模式。  可能的值： -1 - 禁用； 0 - *(默认)* 手动； 1 - 自动。
// For `HostObject` field: `inventory_mode`
const (
	HostInventoryModeDisabled  = -1
	HostInventoryModeManual    = 0
	HostInventoryModeAutomatic = 1
)

// IPMI 认证算法。  可能的值： -1 - *(默认)* 默认； 0 - 无； 1 - MD2； 2 - MD5 4 - straight； 5 - OEM； 6 - RMCP+。
// For `HostObject` field: `IpmiAuthType`
const (
	HostIpmiAuthTypeDefault  = -1
	HostIpmiAuthTypeNone     = 0
	HostIpmiAuthTypeMD2      = 1
	HostIpmiAuthTypeMD5      = 2
	HostIpmiAuthTypeStraight = 4
	HostIpmiAuthTypeOEM      = 5
	HostIpmiAuthTypeRMCP     = 6
)

// IPMI agent的可用性。  可能的值： 0 - *(默认)* 未知； 1 - 可用； 2 - 不可用。
// For `HostObject` field: `IpmiAvailable`
const (
	HostIpmiAvailableUnknown     = 0
	HostIpmiAvailableAvailable   = 1
	HostIpmiAvailableUnavailable = 2
)

// IPMI 权限等级。  可能的值： 1 - 回调； 2 - *(默认)* 用户； 3 - 操作员； 4 - 管理员； 5 - OEM原厂。
// For `HostObject` field: `IpmiPrivilege`
const (
	HostIpmiPrivilegeCallback = 1
	HostIpmiPrivilegeUser     = 2
	HostIpmiPrivilegeOperator = 3
	HostIpmiPrivilegeAdmin    = 4
	HostIpmiPrivilegeOEM      = 5
)

// JMX agent的可用性。  可能的值： 0 - *(默认)* 未知； 1 - 可用； 2 - 不可用。
// For `HostObject` field: `JmxAvailable`
const (
	HostJmxAvailableUnknown     = 0
	HostJmxAvailableAvailable   = 1
	HostJmxAvailableUnavailable = 2
)

// 有效维护的状态。  可能的值：\ 0 - *(默认)* 没有维护； 1 - 有效维护。
// For `HostObject` field: `MaintenanceStatus`
const (
	HostMaintenanceStatusDisable = 0
	HostMaintenanceStatusEnable  = 1
)

// 有效维护类型。  可能的值： 0 - *(默认)* 维护期间搜集数据； 1 - 维护期间不搜集数据。
// For `HostObject` field: `MaintenanceType`
const (
	HostMaintenanceTypeDataCollectionEnabled  = 0
	HostMaintenanceTypeDataCollectionDisabled = 1
)

// SNMP agent的可用性。  可能的值： 0 - *(默认)* 未知； 1 - 可用； 2 - 不可用。
// For `HostObject` field: `SnmpAvailable`
const (
	HostSnmpAvailableUnknown     = 0
	HostSnmpAvailableAvailable   = 1
	HostSnmpAvailableUnavailable = 2
)

// 主机的状态。  可能的值： 0 - *(默认)* 已监控的主机； 1 - 未监控的主机。
// For `HostObject` field: `Status`
const (
	HostStatusMonitored   = 0
	HostStatusUnmonitored = 1
)

// 到主机的连接。  可能的值： 1 - *(默认)* 没有加密； 2 - PSK； 4 - 证书。
// For `HostObject` field: `TLSConnect`
const (
	TLSConnectNoEncryption = 1
	TLSConnectPSK          = 2
	TLSConnectCertificate  = 4
)

// 来自主机的连接。  可能的值： 1 - *(默认)* 没有加密； 2 - PSK； 4 - 证书。
// For `HostObject` field: `TLSAccept`
const (
	TLSAcceptNoEncryption = 1
	TLSAcceptPSK          = 2
	TLSAcceptCertificate  = 4
)

// For `HostGetParams` field: `EvalType`
const (
	HostEvalTypeAndOr = 0
	HostEvalTypeOr    = 2
)

// For `HostTag` field: `Operator`
const (
	HostTagOperatorContains = 0
	HostTagOperatorEquals   = 1
)

// HostObject struct is used to store host operations results
type HostObject struct {
	//主机的ID。
	HostID int `json:"hostid,omitempty"`
	//主机的正式名称。
	Host string `json:"host,omitempty"`
	//Zabbix agent的可用性。  可能的值为： 0 - （默认）未知 1 - 可用； 2 - 不可用。
	Available int `json:"available,omitempty"` // has defined consts, see above
	//说明。
	Description string `json:"description,omitempty"`
	//Zabbix agent不可用状态的下一次轮询时间。
	DisableUntil int `json:"disable_until,omitempty"`
	//Zabbix agent不可用时的错误信息。
	Error string `json:"error,omitempty"`
	//Zabbix agent不可用时的时间。
	ErrorsFrom int `json:"errors_from,omitempty"`
	//主机的来源。  可能的值： 0 - 普通主机； 4 - 自动发现的主机。
	Flags int `json:"flags,omitempty"` // has defined consts, see above
	//主机的资产清单填充模式。  可能的值： -1 - 禁用； 0 - *(默认)* 手动； 1 - 自动。
	InventoryMode int `json:"inventory_mode,omitempty"` // has defined consts, see above
	//IPMI 认证算法。  可能的值： -1 - *(默认)* 默认； 0 - 无； 1 - MD2； 2 - MD5 4 - straight； 5 - OEM； 6 - RMCP+。
	IpmiAuthType int `json:"ipmi_authtype,omitempty"` // has defined consts, see above
	//IPMI agent的可用性。  可能的值： 0 - *(默认)* 未知； 1 - 可用； 2 - 不可用。
	IpmiAvailable int `json:"ipmi_available,omitempty"` // has defined consts, see above
	//IPMI agent不可用状态的下一次轮询时间。
	IpmiDisableUntil int `json:"ipmi_disable_until,omitempty"`
	//IPMI agent不可用时的错误信息。
	IpmiError string `json:"ipmi_error,omitempty"`
	//IPMI agent不可用时的时间。
	IpmiErrorsFrom int `json:"ipmi_errors_from,omitempty"`
	//IPMI 密码。
	IpmiPassword string `json:"ipmi_password,omitempty"`
	//IPMI 权限等级。  可能的值： 1 - 回调； 2 - *(默认)* 用户； 3 - 操作员； 4 - 管理员； 5 - OEM原厂。
	IpmiPrivilege int `json:"ipmi_privilege,omitempty"` // has defined consts, see above
	//IPMI 用户名。
	IpmiUsername string `json:"ipmi_username,omitempty"`
	//JMX agent的可用性。  可能的值： 0 - *(默认)* 未知； 1 - 可用； 2 - 不可用。
	JmxAvailable int `json:"jmx_available,omitempty"` // has defined consts, see above
	//JMX agent不可用状态的下一次轮询时间。
	JmxDisableUntil int `json:"jmx_disable_until,omitempty"`
	//JMX agent不可用时的错误信息。
	JmxError string `json:"jmx_error,omitempty"`
	//JMX agent不可用时的时间。
	JmxErrorsFrom int `json:"jmx_errors_from,omitempty"`
	//有效维护的开始时间。
	MaintenanceFrom int `json:"maintenance_from,omitempty"`
	//有效维护的状态。  可能的值：\ 0 - *(默认)* 没有维护； 1 - 有效维护。
	MaintenanceStatus int `json:"maintenance_status,omitempty"` // has defined consts, see above
	//有效维护类型。  可能的值： 0 - *(默认)* 维护期间搜集数据； 1 - 维护期间不搜集数据。
	MaintenanceType int `json:"maintenance_type,omitempty"` // has defined consts, see above
	//目前对主机生效的维护模式ID。
	MaintenanceID int `json:"maintenanceid,omitempty"`
	//主机的可见名。  默认: `host` 属性值。
	Name string `json:"name,omitempty"`
	//被控主机的Proxy服务器的hostId。
	ProxyHostID int `json:"proxy_hostid,omitempty"`
	//SNMP agent的可用性。  可能的值： 0 - *(默认)* 未知； 1 - 可用； 2 - 不可用。
	SnmpAvailable int `json:"snmp_available,omitempty"` // has defined consts, see above
	//SNMP agent不可用状态的下一次轮询时间。
	SnmpDisableUntil int `json:"snmp_disable_until,omitempty"`
	//SNMP agent不可用时的错误信息。
	SnmpError string `json:"snmp_error,omitempty"`
	//SNMP agent不可用时的时间。
	SnmpErrorsFrom int `json:"snmp_errors_from,omitempty"`
	//主机的状态。  可能的值： 0 - *(默认)* 已监控的主机； 1 - 未监控的主机。
	Status int `json:"status,omitempty"` // has defined consts, see above
	//到主机的连接。  可能的值： 1 - *(默认)* 没有加密； 2 - PSK； 4 - 证书。
	TLSConnect int `json:"tls_connect,omitempty"` // has defined consts, see above
	//来自主机的连接。  可能的值： 1 - *(默认)* 没有加密； 2 - PSK； 4 - 证书。
	TLSAccept int `json:"tls_accept,omitempty"` // has defined consts, see above
	//证书发行机构。
	TLSIssuer string `json:"tls_issuer,omitempty"`
	//证书的主题。
	TLSSubject string `json:"tls_subject,omitempty"`
	//PSK认证。 如果`tls_connect` 或 `tls_accept`启用了PSK，那么该选项是必选。
	//不要将敏感信息放在PSK身份中，它会通过网络以未加密的方式传输，以通知接收者要使用哪个PSK。
	TLSPSKIdentity string `json:"tls_psk_identity,omitempty"`
	//预共享密钥。至少需要32位16进制数字构成。如果`tls_connect` 或 `tls_accept`启用了PSK，那么该选项是必选。
	TLSPSK string `json:"tls_psk,omitempty"`

	//所属主机组
	Groups []HostGroupObject `json:"groups,omitempty"`
	//主机接口列表
	Interfaces []HostInterfaceObject `json:"interfaces,omitempty"`
	//主机标签
	Tags []HostTagObject `json:"tags,omitempty"`
	//继承的主机标签
	InheritedTags []HostTagObject `json:"inheritedTags,omitempty"`
	//宏
	Macros []UserMacroObject `json:"macros,omitempty"`
	//关联的模板，用于创建操作
	Templates []TemplateObject `json:"templates,omitempty"` // Used for `create` operations
	//关联的父模板，用于获取操作时的结构
	ParentTemplates []TemplateObject `json:"parentTemplates,omitempty"` // Used to store result for `get` operations
}

// HostTagObject struct is used to store host tag
type HostTagObject struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// HostGetParams struct is used for host get requests
type HostGetParams struct {
	GetParameters
	GroupIDs       []int `json:"groupids,omitempty"`
	ApplicationIDs []int `json:"applicationids,omitempty"`
	DserviceIDs    []int `json:"dserviceids,omitempty"`
	GraphIDs       []int `json:"graphids,omitempty"`
	HostIDs        []int `json:"hostids,omitempty"`
	HttptestIDs    []int `json:"httptestids,omitempty"`
	InterfaceIDs   []int `json:"interfaceids,omitempty"`
	ItemIDs        []int `json:"itemids,omitempty"`
	MaintenanceIDs []int `json:"maintenanceids,omitempty"`
	MonitoredHosts bool  `json:"monitored_hosts,omitempty"`
	ProxyHosts     bool  `json:"proxy_hosts,omitempty"`
	ProxyIDs       []int `json:"proxyids,omitempty"`
	TemplatedHosts bool  `json:"templated_hosts,omitempty"`
	TemplateIDs    []int `json:"templateids,omitempty"`
	TriggerIDs     []int `json:"triggerids,omitempty"`

	WithItems                     bool            `json:"with_items,omitempty"`
	WithItemPrototypes            bool            `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes bool            `json:"with_simple_graph_item_prototypes,omitempty"`
	WithApplications              bool            `json:"with_applications,omitempty"`
	WithGraphs                    bool            `json:"with_graphs,omitempty"`
	WithGraphPrototypes           bool            `json:"with_graph_prototypes,omitempty"`
	WithHttpTests                 bool            `json:"with_httptests,omitempty"`
	WithMonitoredHttpTests        bool            `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            bool            `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers         bool            `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems          bool            `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  bool            `json:"with_triggers,omitempty"`
	WithProblemsSuppressed        bool            `json:"withProblemsSuppressed,omitempty"`
	EvalType                      int             `json:"evaltype,omitempty"` // has defined consts, see above
	Severities                    []int           `json:"severities,omitempty"`
	Tags                          []HostTagObject `json:"tags,omitempty"`
	InheritedTags                 bool            `json:"inheritedTags,omitempty"`

	// SelectApplications    SelectQuery `json:"selectApplications,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectDiscoveryRule   SelectQuery `json:"selectDiscoveryRule ,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	SelectGroups SelectQuery `json:"selectGroups,omitempty"`
	// SelectHostDiscovery   SelectQuery `json:"selectHostDiscovery ,omitempty"` // not implemented yet
	// SelectHTTPTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	SelectInterfaces SelectQuery `json:"selectInterfaces,omitempty"`
	// SelectInventory       SelectQuery `json:"selectInventory,omitempty"` // not implemented yet
	// SelectItems           SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	SelectMacros          SelectQuery `json:"selectMacros,omitempty"`
	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	// SelectScreens         SelectQuery `json:"selectScreens,omitempty"` // not implemented yet
	SelectTags          SelectQuery `json:"selectTags,omitempty"`
	SelectInheritedTags SelectQuery `json:"selectInheritedTags,omitempty"`
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
}

// Structure to store creation result
type hostCreateResult struct {
	HostIDs []int `json:"hostids"`
}

// Structure to store updation result
type hostUpdateResult struct {
	HostIDs []int `json:"hostids"`
}

// Structure to store deletion result
type hostDeleteResult struct {
	HostIDs []int `json:"hostids"`
}

// HostGet gets hosts
func (z *Context) HostGet(params HostGetParams) ([]HostObject, int, error) {

	var result []HostObject

	status, err := z.request("host.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostCreate creates hosts
func (z *Context) HostCreate(params []HostObject) ([]int, int, error) {

	var result hostCreateResult

	status, err := z.request("host.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostIDs, status, nil
}

// HostUpdate updates hosts
func (z *Context) HostUpdate(params []HostObject) ([]int, int, error) {

	var result hostUpdateResult

	status, err := z.request("host.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostIDs, status, nil
}

// HostDelete deletes hosts
func (z *Context) HostDelete(hostIDs []int) ([]int, int, error) {

	var result hostDeleteResult

	status, err := z.request("host.delete", hostIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostIDs, status, nil
}
