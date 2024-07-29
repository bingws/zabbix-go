package zabbix

// 该接口是否在主机上用作默认的接口，在主机上只能将某种类型的一个接口设置为默认值。  可能的值有: 0 - 不是默认; 1 - 默认。
// For `HostInterfaceObject` field: `Main`
const (
	HostInterfaceMainNotDefault = 0
	HostInterfaceMainDefault    = 1
)

// 接口类型。  可能的值有: 1 - agent; 2 - SNMP; 3 - IPMI; 4 - JMX。
// For `HostInterfaceObject` field: `Type`
const (
	HostInterfaceTypeAgent = 1
	HostInterfaceTypeSNMP  = 2
	HostInterfaceTypeIPMI  = 3
	HostInterfaceTypeJMX   = 4
)

// 是否通过IP进行连接。  可能的值有: 0 - 使用主机DNS进行连接; 1 - 使用主机接口的主机IP进行连接。
// For `HostInterfaceObject` field: `UseIP`
const (
	HostInterfaceUseIpDNS = 0
	HostInterfaceUseIpIP  = 1
)

// 是否使用批量的SNMP请求.  可能的值有: 0 - 不使用批量请求; 1 - (默认) - 使用批量请求。
// For `HostInterfaceDetailsTagObject` field: `Bulk`
const (
	HostInterfaceDetailsTagBulkDontUse = 0
	HostInterfaceDetailsTagBulkUse     = 1
)

// SNMP接口版本。  可能的值： 1 - SNMPv1； 2 - (默认) - SNMPv2c； 3 - SNMPv3
// For `HostInterfaceDetailsTagObject` field: `Version`
const (
	HostInterfaceDetailsTagVersionSNMPv1  = 1
	HostInterfaceDetailsTagVersionSNMPv2c = 2
	HostInterfaceDetailsTagVersionSNMPv3  = 3
)

// SNMPv3安全级别。仅由SNMPv3接口使用。  可能的值： 0 - (默认) - noAuthNoPriv； 1 - authNoPriv； 2 - authPriv。
// For `HostInterfaceDetailsTagObject` field: `SecurityLevel`
const (
	HostInterfaceDetailsTagSecurityLevelNoAuthNoPriv = 0
	HostInterfaceDetailsTagSecurityLevelAuthNoPriv   = 1
	HostInterfaceDetailsTagSecurityLevelAuthPriv     = 2
)

// SNMPv3身份验证协议。仅由SNMPv3接口使用。  可能的值： 0 - (默认) - MD5； 1 - SHA。
// For `HostInterfaceDetailsTagObject` field: `AuthProtocol`
const (
	HostInterfaceDetailsTagAuthProtocolMD5 = 0
	HostInterfaceDetailsTagAuthProtocolSHA = 1
)

// SNMPv3隐私协议。仅由SNMPv3接口使用。  可能的值： 0 - (默认) - DES； 1 - AES。
// For `HostInterfaceDetailsTagObject` field: `PrivProtocol`
const (
	HostInterfaceDetailsTagPrivProtocolDES = 0
	HostInterfaceDetailsTagPrivProtocolAES = 1
)

// HostInterfaceObject 主机接口对象
type HostInterfaceObject struct {
	//接口的ID。
	InterfaceID int `json:"interfaceid,omitempty"`
	//接口使用的DNS名称  如果通过IP直接连接，则可以为空。
	DNS string `json:"dns"`
	//接口所属的主机ID。
	HostID int `json:"hostid,omitempty"`
	//接口使用的IP地址。 如果通过DNS域名连接，可以设置为空。
	IP string `json:"ip"`
	//该接口是否在主机上用作默认的接口，在主机上只能将某种类型的一个接口设置为默认值。  可能的值有: 0 - 不是默认; 1 - 默认。
	Main int `json:"main"` // has defined consts, see above
	//接口使用的端口号，可以包含用户宏。
	Port string `json:"port"`
	//接口类型。  可能的值有: 1 - agent; 2 - SNMP; 3 - IPMI; 4 - JMX。
	Type int `json:"type"` // has defined consts, see above
	//是否通过IP进行连接。  可能的值有: 0 - 使用主机DNS进行连接; 1 - 使用主机接口的主机IP进行连接。
	UseIP int `json:"useip"` // has defined consts, see above
	//接口的附加对象。 如果接口 'type'是SNMP，则必选
	Details HostInterfaceDetailsTagObject `json:"details,omitempty"`

	// Items []ItemObject `json:"items,omitempty"` // not implemented yet
	Hosts []HostObject `json:"hosts,omitempty"`
}

// HostInterfaceDetailsTagObject 主机SNMP接口详情对象
type HostInterfaceDetailsTagObject struct {
	//SNMP接口版本。  可能的值： 1 - SNMPv1； 2 - (默认) - SNMPv2c； 3 - SNMPv3
	Version int `json:"version,omitempty"` // has defined consts, see above
	//是否使用批量的SNMP请求.  可能的值有: 0 - 不使用批量请求; 1 - (默认) - 使用批量请求。
	Bulk int `json:"bulk,omitempty"` // has defined consts, see above
	//SNMP 团体字 (必选)。 仅在SNMPv1和SNMPv2接口中使用。
	Community string `json:"community,omitempty"`
	//SNMPv3 安全名称。仅在SNMPv3接口中使用。
	SecurityName string `json:"securityname,omitempty"`
	//SNMPv3安全级别。仅由SNMPv3接口使用。  可能的值： 0 - (默认) - noAuthNoPriv； 1 - authNoPriv； 2 - authPriv。
	SecurityLevel int `json:"securitylevel,omitempty"` // has defined consts, see above
	//SNMPv3认证密码。仅由SNMPv3接口使用。
	AuthPassPhrase string `json:"authpassphrase,omitempty"`
	//SNMPv3私有密码。仅由SNMPv3接口使用。
	PrivPassPhrase string `json:"privpassphrase,omitempty"`
	//SNMPv3身份验证协议。仅由SNMPv3接口使用。  可能的值： 0 - (默认) - MD5； 1 - SHA。
	AuthProtocol int `json:"authprotocol,omitempty"` // has defined consts, see above
	//SNMPv3隐私协议。仅由SNMPv3接口使用。  可能的值： 0 - (默认) - DES； 1 - AES。
	PrivProtocol int `json:"privprotocol,omitempty"` // has defined consts, see above
	//SNPv3上下文名称。仅由SNMPv3接口使用。
	ContextName string `json:"contextname,omitempty"`
}

// HostInterfaceGetParams struct is used for hostInterface get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/hostInterface/get#parameters
type HostInterfaceGetParams struct {
	GetParameters

	HostIDs      []int `json:"hostids,omitempty"`
	InterfaceIDs []int `json:"interfaceids,omitempty"`
	ItemIDs      []int `json:"itemids,omitempty"`
	TriggerIDs   []int `json:"triggerids,omitempty"`

	// SelectItems SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`
}

// Structure to store creation result
type hostInterfaceCreateResult struct {
	InterfaceIDs []int `json:"interfaceids"`
}

// Structure to store deletion result
type hostInterfaceDeleteResult struct {
	InterfaceIDs []int `json:"interfaceids"`
}

// HostInterfaceGet gets hostInterfaces
func (z *Context) HostInterfaceGet(params HostInterfaceGetParams) ([]HostInterfaceObject, int, error) {

	var result []HostInterfaceObject

	status, err := z.request("hostInterface.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostInterfaceCreate creates hostInterfaces
func (z *Context) HostInterfaceCreate(params []HostInterfaceObject) ([]int, int, error) {

	var result hostInterfaceCreateResult

	status, err := z.request("hostInterface.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.InterfaceIDs, status, nil
}

// HostInterfaceDelete deletes hostInterfaces
func (z *Context) HostInterfaceDelete(hostInterfaceIDs []int) ([]int, int, error) {

	var result hostInterfaceDeleteResult

	status, err := z.request("hostInterface.delete", hostInterfaceIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.InterfaceIDs, status, nil
}
