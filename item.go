package zabbix

type ItemType int

const (
	ItemTypeAgent   ItemType = iota
	ItemTypeTrapper          = iota + 1
	ItemTypeSimpleCheck
	ItemTypeInternal    = iota + 2
	ItemTypeAgentActive = iota + 3
	ItemTypeAggregate
	ItemTypeWebItem
	ItemTypeExternalCheck
	ItemTypeDatabaseMonitor
	ItemTypeIPMIAgent
	ItemTypeSSHAgent
	ItemTypeTelnetAgent
	ItemTypeCalculated
	ItemTypeJMXAgent
	ItemTypeSNMPTrap
	ItemTypeDependentItem
	ItemTypeHTTPAgent
	ItemTypeSNMPAgent
)

type ItemAllowTrapsType int

const (
	ItemAllowTrapsNot ItemAllowTrapsType = iota
	ItemAllowTrapsYes
)

type ItemAuthType int

const (
	ItemAuthTypeNone = iota
	ItemAuthTypeBasic
	ItemAuthTypeNTLM
	ItemAuthTypeKerberos
)

type ItemFlagsType int

const (
	ItemFlagsPlain ItemFlagsType = iota
	ItemFlagsDiscovered
)

type ItemFollowRedirectsType int

const (
	ItemFollowRedirectsNo ItemFollowRedirectsType = iota
	ItemFollowRedirectsYes
)

type ItemOutputFormatType int

const (
	ItemOutputFormatRaw ItemOutputFormatType = iota
	ItemOutputFormatJSON
)

type ItemPostType int

const (
	ItemPostTypeRaw ItemPostType = iota
	ItemPostTypeJSON
	ItemPostTypeXML
)

type ItemRequestMethodType int

const (
	ItemRequestMethodGET ItemRequestMethodType = iota
	ItemRequestMethodPOST
	ItemRequestMethodPUT
	ItemRequestMethodHEAD
)

type ItemRetrieveModeType int

const (
	ItemRetrieveModeBody ItemRetrieveModeType = iota
	ItemRetrieveModeHeaders
	ItemRetrieveModeBoth
)

type ItemStateType int

const (
	ItemStateNormal ItemStateType = iota
	ItemStateNotSupported
)

type ItemStatusType int

const (
	ItemStatusEnabled ItemStatusType = iota
	ItemStatusDisabled
)

type ItemVerifyHostType int

const (
	ItemVerifyHostNo ItemVerifyHostType = iota
	ItemVerifyHostYes
)

type ItemVerifyPeerType int

const (
	ItemVerifyPeerNo ItemVerifyPeerType = iota
	ItemVerifyPeerYes
)

type ItemObject struct {
	ItemID string `json:"itemid,omitempty"`
	// 更新监控项的时间间隔。 接受具有后缀(30s,1m,2h,1d)时间单位  也接受用户宏
	Delay string `json:"delay"`
	// 监控项原型所属的主机的ID 对于更新操作，这个字段是 *只读*.
	HostID string `json:"hostid"`
	// 监控项主机接口的ID。 仅用于主机项。
	InterfaceID string `json:"interfaceid"`
	// 监控项原型的键
	Key string `json:"key_"`
	// 监控项原型的名称
	Name string `json:"name"`
	// 监控项原型的类型。
	// # todo 确认 get请求示例是 string "0"  object页面介绍是int
	Type ItemType `json:"type"`
	// 仅在HTTP agent监控项原型有要求的URL字符串
	Url string `json:"url"`
	// HTTP agent监控项字段
	AllowTraps ItemAllowTrapsType `json:"allow_traps,omitempty"`
	// 仅在SSH agent 监控项 或 HTTP agent 监控项中使用。
	AuthType ItemAuthType `json:"authtype,omitempty"`
	// 监控项说明。
	Description string `json:"description,omitempty"`
	// *只读* 当更新监控项出错时的错误文本。
	Error string `json:"error,omitempty"`
	// *只读* 原始监控项
	Flags ItemFlagsType `json:"flags,omitempty"`
	// HTTP agent 监控项字段，合并数据时跟随重定向。
	FollowRedirects ItemFollowRedirectsType `json:"follow_redirects,omitempty"`
	// HTTP agent 监控项字段。带有HTTP(S)请求报头的对象，报头名为键名，报头值为值。
	//  事例: { "User-Agent": "Zabbix" }
	Headers []interface{} `json:"headers,omitempty"`
	// 一个历史数据被保存的时长的时间单位。接受用户宏。 默认值: 90d.
	History string `json:"history,omitempty"`
	// HTTP agent 监控项字段。HTTP(S)代理连接字符串。
	HttpProxy string `json:"http_proxy,omitempty"`
	// 监控项填充的主机资产的ID。 默认0
	InventoryLink int `json:"inventory_link,omitempty"`
	// IPMI传感器。仅用于IPMI监控项
	IpmiSensor string `json:"ipmi_sensor,omitempty"`
	// JMX agent自定义的连接字符串。 默认值: service:jmx:rmi:///jndi/rmi://{HOST.CONN}:{HOST.PORT}/jmxrmi
	JmxEndpoint string `json:"jmx_endpoint,omitempty"`
	// (**只读**)* 监控项最后被更新的时间。 此属性将只返回[ZBX_HISTORY_PERIOD](https://www.zabbix.com/documentation/5.0/zh/manual/web_interface/definitions)中配置的周期值。
	LastClock string `json:"lastclock,omitempty"`
	// (**只读**)* 监控项最后被更新的纳秒。 此属性将只返回[ZBX_HISTORY_PERIOD](https://www.zabbix.com/documentation/5.0/zh/manual/web_interface/definitions)中配置的周期值。
	LasTns string `json:"lastns,omitempty"`
	//*(**只读**)* 监控项最新的值。 此属性将只返回[ZBX_HISTORY_PERIOD](https://www.zabbix.com/documentation/5.0/zh/manual/web_interface/definitions)中配置的周期值。
	LastValue string `json:"lastvalue,omitempty"`
	// 日志条目的时间格式。仅用于日志监控项。
	LogTimeFmt string `json:"logtimefmt,omitempty"`
	// 主监控项ID 允许多达3个依赖监控项的递归和监控项的最大计数等于999 要求依赖项。
	MasterItemID string `json:"master_itemid,omitempty"`
	// 上次更新所监视日志文件的时间。仅用于日志项。
	Mtime string `json:"mtime,omitempty"`
	// HTTP agent监控项字段。返回数据应被转换成JSON。 0 - *(**默认值**)* 存储raw. 1 - 转成JSON.
	OutputFormat int `json:"output_format,omitempty"`
	// 取决于监控项类型的附加参数： - SSH、Telnet项执行脚本 - SQL查询数据库监控项; - 计算项目公式.
	Params string `json:"params,omitempty"`
	// 认证的密码。用于simple check, SSH, Telnet, database monitor, JMX and HTTP agent items. 当JMX使用用户名时，用户名应该和密码一起指定，或者两个属性都应该留空。
	Password string `json:"password,omitempty"`
	// 监控项监控的端口。仅用于SMNP监控项。
	Port string `json:"port,omitempty"`
	// HTTP agent字段。存储在post属性的post的数据类型。 0 - *(**默认值**)* Raw data. 2 - JSON data. 3 - XML data.
	PostType ItemPostType `json:"post_type,omitempty"`
	// HTTP agent字段。HTTP(S)请求报文。仅用于post_type。
	Posts string `json:"posts,omitempty"`
	// *(**只读**)*监控项的前一个值。 此属性将只返回[ZBX_HISTORY_PERIOD](https://www.zabbix.com/documentation/5.0/zh/manual/web_interface/definitions)中配置的周期值。
	PrevValue string `json:"prevvalue,omitempty"`
	//私钥文件名。
	PrivateKey string `json:"privatekey,omitempty"`
	//公钥的文件名。
	PublicKey string `json:"publickey,omitempty"`
	// HTTP agent监控项字段。查询参数。带有键值对的数组对象，值可为空。
	QueryFields []interface{} `json:"query_fields,omitempty"`
	// HTTP agent监控项字段。请求方法的类型。 0 - GET 1 - *(**默认值**)* POST 2 - PUT 3 - HEAD
	RequestMethod ItemRequestMethodType `json:"request_method,omitempty"`
	// HTTP agent监控项字段。被存储的响应的部分。
	RetrieveMode ItemRetrieveModeType `json:"retrieve_mode,omitempty"`
	// SNMP OID.
	SnmpOid string `json:"snmp_oid,omitempty"`
	// HTTP agent监控项字段。公共SSL 秘钥的文件路径。
	SslCertFile string `json:"ssl_cert_file,omitempty"`
	// HTTP agent监控项字段。私有SLL秘钥的文件路径。
	SslKeyFile string `json:"ssl_key_file,omitempty"`
	// HTTP agent监控项字段。SSL 秘钥的文件密码。
	SslKeyPassword string `json:"ssl_key_password,omitempty"`
	// (**只读**)* 监控项状态. 取值范围: 0 - *(**默认值**)* 标准; 1 - 不支持.
	State ItemStateType `json:"state,omitempty"`
	// 监控项状态. 取值范围: 0 - *(**默认值**)* 启用; 1 - 禁用.
	Status ItemStatusType `json:"status,omitempty"`
	// HTTP agent监控项字段。以逗号分隔的HTTP 状态码的范围。也支持作为逗号分隔的用户宏列表。 事例: 200,200-{$M},{$M},200-400
	StatusCodes string `json:"status_codes,omitempty"`
	//（只读）父模板的ID。
	TemplateID string `json:"templateid,omitempty"`
	// HTTP agent监控项字段。监控项数据轮询超时时间。支持用户宏。 默认值: 3s 最大值: 60s
	Timeout string `json:"timeout,omitempty"`
	//接受的主机。仅用于trapper监控项或者HTTP agent监控项。
	TrapperHosts string `json:"trapper_hosts,omitempty"`
	//时间单位，数据数据被保存的时间长度。也接受用户宏。 默认值: 365d.
	Trends string `json:"trends,omitempty"`
	// Value units. 值的单位。
	Units string `json:"units,omitempty"`
	//认证的用户名。用于simple check, SSH, Telnet, database monitor, JMX and HTTP agent 监控项。 要求提供。\\当被JMX使用时，密码也要和用户名一起被提供或者一起留空。
	Username string `json:"username,omitempty"`
	//关联映射值的ID。
	ValueMapID string `json:"valuemapid,omitempty"`
	// HTTP agent字段。验证URL中的主机名处于通用名称字段或主机证书的主题备用名称字段 0 - *(**默认值**)* 不验证. 1 - 验证.
	VerifyHost ItemVerifyHostType `json:"verify_host,omitempty"`
	// HTTP agent字段。验证主机的合法性。 0 - *(**默认值**)* 不验证. 1 - 验证
	VerifyPeer ItemVerifyPeerType `json:"verify_peer,omitempty"`

	// 这些字段很奇怪  只有在get例子里有 https://www.zabbix.com/documentation/5.0/en/manual/api/reference/item/get
	ValueType string `json:"value_type,omitempty"`
	Lifetime  string `json:"lifetime,omitempty"`
	EvalType  string `json:"evaltype,omitempty"`
}

type itemUpdateResult struct {
	ItemIDs []int `json:"itemids"`
}

type ItemTagFilter struct {
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Operator int    `json:"operator"`
}

type ItemGetParams struct {
	GetParameters

	// ItemIDs filters search results to items with the given Item ID's.
	ItemIDs []string `json:"itemids,omitempty"`

	// GroupIDs filters search results to items belong to the hosts
	// of the given Group ID's.
	GroupIDs []string `json:"groupids,omitempty"`

	// TemplateIDs filters search results to items belong to the
	// given templates of the given Template ID's.
	TemplateIDs []string `json:"templateids,omitempty"`

	// HostIDs filters search results to items belong to the
	// given Host ID's.
	HostIDs []string `json:"hostids,omitempty"`

	// ProxyIDs filters search results to items that are
	// monitored by the given Proxy ID's.
	ProxyIDs []string `json:"proxyids,omitempty"`

	// InterfaceIDs filters search results to items that use
	// the given host Interface ID's.
	InterfaceIDs []string `json:"interfaceids,omitempty"`

	// GraphIDs filters search results to items that are used
	// in the given graph ID's.
	GraphIDs []string `json:"graphids,omitempty"`

	// TriggerIDs filters search results to items that are used
	// in the given Trigger ID's.
	TriggerIDs []string `json:"triggerids,omitempty"`

	// ApplicationIDs filters search results to items that
	// belong to the given Applications ID's.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// WebItems flag includes web items in the result.
	WebItems bool `json:"webitems,omitempty"`

	// Inherited flag return only items inherited from a template
	// if set to 'true'.
	Inherited bool `json:"inherited,omitempty"`

	// Templated flag return only items that belong to templates
	// if set to 'true'.
	Templated bool `json:"templated,omitempty"`

	// Monitored flag return only enabled items that belong to
	// monitored hosts if set to 'true'.
	Monitored bool `json:"monitored,omitempty"`

	// Group filters search results to items belong to a group
	// with the given name.
	Group string `json:"group,omitempty"`

	// Host filters search results to items that belong to a host
	// with the given name.
	Host string `json:"host,omitempty"`

	// Application filters search results to items that belong to
	// an application with the given name.
	Application string `json:"application,omitempty"`

	// WithTriggers flag return only items that are used in triggers
	WithTriggers bool `json:"with_triggers,omitempty"`

	// Filter by tags
	Tags []ItemTagFilter `json:"tags,omitempty"`
}

func (z *Context) ItemCreate(params []ItemObject) ([]int, int, error) {
	var result struct {
		ItemIDs []int `json:"itemids"`
	}

	status, err := z.request("item.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemIDs, status, nil
}

func (z *Context) ItemDelete(itemIDs []string) ([]string, int, error) {
	var result struct {
		ItemIDs []string `json:"itemids"`
	}

	status, err := z.request("item.delete", itemIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemIDs, status, nil
}

func (z *Context) ItemUpdate(params []ItemObject) ([]int, int, error) {
	var result itemUpdateResult

	status, err := z.request("item.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ItemIDs, status, nil
}

func (z *Context) ItemGet(params ItemGetParams) ([]ItemObject, int, error) {
	var result []ItemObject

	status, err := z.request("item.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
