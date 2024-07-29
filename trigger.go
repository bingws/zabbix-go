package zabbix

// Flags 是一个只读的整数，它代表了原始触发器。可能的值有：0表示普通触发器，4表示自动发现的触发器。
const (
	FlagsNormal    = 0
	FlagsDiscovery = 4
)

// Priority 是一个整数，它代表了触发器的严重性级别。允许的值有：0表示未分类，1表示信息，2是警告，3是一般严重，4是严重，5是灾难。
const (
	PriorityNotClassified = 0
	PriorityInfo          = 1
	PriorityWarn          = 2
	PriorityAlarm         = 3
	PriorityFatal         = 4
)

// State 是一个只读的整数，它代表了触发器的状态。可能的值有：0表示触发器状态是最新的，1表示当前的触发器状态是未知的。
const (
	StateLatest  = 0
	StateUnknown = 1
)

// Status 是一个整数，它表示触发器是否处于启用状态或禁用状态。可能的值有：0表示启用，1表示禁用。
const (
	statusEnabled  = 0
	statusDisabled = 1
)

// Type 是一个整数，它代表了触发器是否能够生成多个故障事件。可能的值有：0表示不生成多个事件，1表示生成多个事件。
const (
	TypeSingleEvent    = 0
	TypeMultipleEvents = 1
)

// Value 是一个只读的整数，它代表了触发器是否处于正常或故障状态。可能的值有：0表示OK，1表示故障。
const (
	ValueOK      = 0
	ValueProblem = 1
)

// RecoveryMode 是一个整数，它代表了事件恢复生成模式。可能的值有：0表示表达式，1表示恢复表达式，2表示无。
const (
	RecoveryModeExpression         = 0
	RecoveryModeRecoveryExpression = 1
	RecoveryModeNone               = 2
)

// CorrelationMode 是一个整数，它代表了事件恢复关闭。可能的值有：0表示所有故障，1表示与标签值匹配的所有故障。
const (
	CorrelationModeAllProblems = 0
	CorrelationModeTagMatch    = 1
)

// ManualClose 是一个整数，它代表了是否允许手动关闭。可能的值有：0表示不允许，1表示允许。这个字段用于控制触发器是否可以由用户手动重置或关闭。
const (
	ManualCloseNotAllowed = 0
	ManualCloseAllowed    = 1
)

// TriggerObject 触发器对象
type TriggerObject struct {

	// TriggerID 触发器的ID。
	TriggerID string `json:"triggerid,omitempty"`

	// Description 触发器的名称。
	Description string `json:"description,omitempty"`

	// Expression 生成的触发表达式。
	Expression string `json:"expression,omitempty"`

	// Comments 触发器的附加说明。
	Comments string `json:"comments,omitempty"`

	// Error 错误概述，如果在更新触发器的状态时出现任何问题。
	Error string `json:"error,omitempty"`

	// Flags 原始触发器。      许可值为：   0 - 普通触发器；   4 - 自动发现的触发器。
	Flags int `json:"flags,omitempty"`

	// LastChangeTimestamp 触发器最后更改其状态的时间。
	LastChangeTimestamp int `json:"lastchangetimestamp,omitempty"`

	// Priority 触发器的严重性级别。      许可值为：   0 - *(**默认**)* 未分类；   1 - 信息；   2 - 警告；   3 - 一般严重；   4 - 严重；   5 - 灾难。
	Priority int `json:"priority,omitempty"`

	// State 触发器的状态。      许可值：   0 - *(**默认**)* 触发器状态是最新的；   1 - 当前的触发器状态是未知的。
	State int `json:"state,omitempty"`

	// Status 触发器是否处于启用状态或禁用状态。      许可值为：   0 - *(**默认**)* 启用；   1 - 禁用。
	Status int `json:"status,omitempty"`

	// TemplateID 父触发器模板ID。
	TemplateID string `json:"templateid,omitempty"`

	// Type 触发器是否能够生成多个故障事件。      许可值为：   0 - *(**默认**)* 不生成多个事件。   1 - 生成多个事件。
	Type int `json:"type,omitempty"`

	// URL与触发器相关联的URL。
	URL string `json:"url,omitempty"`

	// Value 触发器是否处于正常或故障状态。      许可值为：   0 - *(**默认**)* OK; 正常；   1 - 故障。
	Value int `json:"value,omitempty"`

	// RecoveryMode 事件恢复生成模式。      许可值为：   0 - *(**默认**)* 表达式；   1 - 恢复表达式；   2 - 无。
	RecoveryMode int `json:"recovery_mode,omitempty"`

	// RecoveryExpression 生成的触发恢复表达式。
	RecoveryExpression string `json:"recovery_expression,omitempty"`

	// CorrelationMode 事件恢复关闭。      许可值为：   0 - *(**默认**)* 所有故障；   1 - 与标签值匹配的所有故障。
	CorrelationMode int `json:"correlation_mode,omitempty"`

	// CorrelationTag 用于匹配的标签。
	CorrelationTag string `json:"correlation_tag,omitempty"`

	// ManualClose 允许手动关闭。      许可值为：   0 - *(**默认**)* 不允许；   1 - 允许。
	ManualClose int `json:"manual_close,omitempty"`
}

// TriggerTagObject struct is used to store trigger tag
type TriggerTagObject struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// TriggerGetParams struct is used for trigger get requests
type TriggerGetParams struct {
	GetParameters

	TriggerIDs     []int `json:"triggerids,omitempty"`
	GroupIDs       []int `json:"groupids,omitempty"`
	HostIDs        []int `json:"hostids,omitempty"`
	ItemIDs        []int `json:"itemids,omitempty"`
	ApplicationIDs []int `json:"applicationids,omitempty"`

	Functions                   []string    `json:"functions,omitempty"`
	Group                       string      `json:"group,omitempty"`
	Host                        string      `json:"host,omitempty"`
	OnlyTrue                    bool        `json:"only_true,omitempty"`
	Monitored                   bool        `json:"monitored,omitempty"`
	Active                      bool        `json:"active,omitempty"`
	Maintenance                 bool        `json:"maintenance,omitempty"`
	WithUnacknowledged          bool        `json:"withUnacknowledged,omitempty"`
	WithAcknowledged            bool        `json:"withAcknowledged,omitempty"`
	WithLastEventUnacknowledged bool        `json:"withLastEventUnacknowledged,omitempty"`
	RecentProblemOnly           bool        `json:"recentProblemOnly,omitempty"`
	SortField                   []string    `json:"sortfield,omitempty"`
	ExpandData                  string      `json:"expandData,omitempty"`
	SelectGroups                SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts                 SelectQuery `json:"selectHosts,omitempty"`
	SelectItems                 SelectQuery `json:"selectItems,omitempty"`
	SelectFunctions             SelectQuery `json:"selectFunctions,omitempty"`
	SelectDependencies          SelectQuery `json:"selectDependencies,omitempty"`
	SelectDiscoveryRule         SelectQuery `json:"selectDiscoveryRule,omitempty"`
	SelectLastEvent             SelectQuery `json:"selectLastEvent,omitempty"`
	SelectTags                  SelectQuery `json:"selectTags,omitempty"`
	SelectTriggerDiscovery      SelectQuery `json:"selectTriggerDiscovery,omitempty"`
	LimitSelects                int         `json:"limitSelects,omitempty"`
}

// Structure to store creation result
type triggerCreateResult struct {
	TriggerIDs []int `json:"triggerids"`
}

// Structure to store updation result
type triggerUpdateResult struct {
	TriggerIDs []int `json:"triggerids"`
}

// Structure to store deletion result
type triggerDeleteResult struct {
	TriggerIDs []int `json:"triggerids"`
}

// TriggerAddDependenciesResult Structure to store AddDependencies result
type TriggerAddDependenciesResult struct {
	TriggerID           string   `json:"triggerid"`
	DependsOnTriggerIDs []string `json:"dependsOnTriggerid"`
}

// TriggerDeleteDependenciesResult Structure to store DeleteDependencies result
type TriggerDeleteDependenciesResult struct {
	TriggerID           string   `json:"triggerid"`
	DependsOnTriggerIDs []string `json:"dependsOnTriggerid"`
}

// TriggerGet gets triggers
func (z *Context) TriggerGet(params TriggerGetParams) ([]TriggerObject, int, error) {

	var result []TriggerObject

	status, err := z.request("trigger.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// TriggerCreate creates triggers
func (z *Context) TriggerCreate(params []TriggerObject) ([]int, int, error) {

	var result triggerCreateResult

	status, err := z.request("trigger.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TriggerIDs, status, nil
}

// TriggerUpdate updates triggers
func (z *Context) TriggerUpdate(params []TriggerObject) ([]int, int, error) {

	var result triggerUpdateResult

	status, err := z.request("trigger.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TriggerIDs, status, nil
}

// TriggerDelete deletes triggers
func (z *Context) TriggerDelete(triggerIDs []int) ([]int, int, error) {

	var result triggerDeleteResult

	status, err := z.request("trigger.delete", triggerIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TriggerIDs, status, nil
}

// TriggerAddDependencies adds dependencies to triggers
func (z *Context) TriggerAddDependencies(params TriggerAddDependenciesResult) ([]int, int, error) {
	var result triggerCreateResult

	status, err := z.request("trigger.adddependencies", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TriggerIDs, status, nil
}

// TriggerDeleteDependencies deletes dependencies from triggers
func (z *Context) TriggerDeleteDependencies(params TriggerDeleteDependenciesResult) ([]int, int, error) {
	var result triggerCreateResult

	status, err := z.request("trigger.deletedependencies", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TriggerIDs, status, nil
}
