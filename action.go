package zabbix

type ActionStatusType int

// ！！！ zabbix本身json不规范、数据库不规范  snake等问题  没办法

const (
	ActionStatusEnabled ActionStatusType = iota
	ActionStatusDisabled
)

type ActionPauseSuppressedType int

const (
	ActionPauseSuppressedEnabled  ActionPauseSuppressedType = iota
	ActionPauseSuppressedDisabled                           = 1
)

//  ############################################### ActionOperationType ##############################################

type ActionsOperationType int

const (
	ActionOperationTypeSendMsg              ActionsOperationType = iota //动作更新操作 || 动作操作 || 动作恢复操作
	ActionOperationTypeRemoteCmd                                        //动作更新操作 || 动作操作 || 动作恢复操作
	ActionOperationTypeAddHost                                          //动作操作
	ActionOperationTypeRmHost                                           //动作操作
	ActionOperationTypeAddToHostGroup                                   //动作操作
	ActionOperationTypeRmFromHostGroup                                  //动作操作
	ActionOperationTypeLinkToTpl                                        //动作操作
	ActionOperationTypeUnlinkFromTpl                                    //动作操作
	ActionOperationTypeEnableHost                                       //动作操作
	ActionOperationTypeDisableHost                                      //动作操作
	ActionOperationTypeSetHostInventoryMode                             //动作操作
	// ActionsOperationRecoveryTypeNotifyAllInvolved 仅解决问题 通知相关所有人  CM里应该是始终通知一个
	ActionsOperationRecoveryTypeNotifyAllInvolved //动作恢复操作
	// ActionsOperationUpdateTypeNotifyAllInvolved 更新问题   通知相关所有人  CM里应该是始终通知一个
	ActionsOperationUpdateTypeNotifyAllInvolved // 动作更新操作
	ActionOperationConditionEventAcknowledged   = 14
)

type ActionOperationEvalType int

const (
	ActionOperationEvalTypeAndOR ActionOperationEvalType = iota
	ActionOperationEvalTypeAnd                           = 1
	ActionOperationEvalTypeOr                            = 2
)

type ActionOperationCommandType int

const (
	ActionOperationCommandTypeCustomScript ActionOperationCommandType = 0
	ActionOperationCommandTypeIPMI                                    = 1
	ActionOperationCommandTypeSSH                                     = 2
	ActionOperationCommandTypeTelnet                                  = 3
	ActionOperationCommandTypeGlobalScript                            = 4
)

type ActionOperationCommandAuthType int

const (
	ActionOperationCommandAuthTypePassword ActionOperationCommandAuthType = 0
	ActionOperationCommandAuthTypePubKey                                  = 1
)

type ActionOperationCommandExecuteOnType int

const (
	ActionOperationCommandExecuteOnAgent  = 0
	ActionOperationCommandExecuteOnServer = 1
	ActionOperationCommandExecuteOnProxy  = 2
)

type ActionOperationMessageDefaultMsgFromType int

const (
	ActionOperationMessageDefaultMsgFromOperation ActionOperationMessageDefaultMsgFromType = 0
	ActionOperationMessageDefaultMsgFromMediaType                                          = 1
)

type ActionObject struct {
	ActionID int `json:"actionid,omitempty"`
	// 默认步骤持续时间  若当前步骤时间走完 则进行下一步骤（升级概念）   > 60s  也可以是分钟单位 1m
	EscPeriod int `json:"esc_period"`
	// Eventsource eventSourceType 事件源类型
	Eventsource int    `json:"eventsource"`
	Name        string `json:"name"`
	// Status actionStatusType 事件状态
	Status int `json:"status,omitempty"`
	// PauseSuppressed ActionPauseSuppressedType 是否在维护期间暂停升级
	PauseSuppressed ActionPauseSuppressedType `json:"pause_suppressed,omitempty"`

	Operations            []ActionOperationObject         `json:"operations,omitempty"`
	Filter                ActionFilterObject              `json:"filter,omitempty"`
	RecoveryOperations    []ActionRecoveryOperationObject `json:"recovery_operations,omitempty"`
	AcknowledgeOperations []ActionRecoveryOperationObject `json:"acknowledge_operations,omitempty"`
}

type ActionOperationObject struct {
	OperationID int `json:"operationid,omitempty"`
	// OperationType ActionsOperationType 操作类型
	OperationType ActionsOperationType `json:"operationtype"`
	ActionID      int                  `json:"actionid,omitempty"`
	// 操作持续时间 触发器操作中，问题恢复和更新操作不支持升级
	// 如果设置为0或0s，将使用默认的操作升级周期
	EscPeriod int `json:"esc_period,omitempty"`
	// 当前步骤(轮次)开始编号 默认1
	EscStepFrom int `json:"esc_step_from,omitempty"`
	// 步骤结束(轮次)结束编号 默认1
	EscStepTo int `json:"esc_step_to,omitempty"`
	// EvalType ActionOperationEvalType 条件运算符类型
	EvalType     ActionOperationEvalType          `json:"evaltype,omitempty"`
	OpCommand    ActionOperationCommandObject     `json:"opcommand,omitempty"`
	OpCommandGrp []ActionOpCommandGrpObject       `json:"opcommand_grp,omitempty"`
	OpCommandHst []ActionOpCommandHostObject      `json:"opcommand_hst,omitempty"`
	OpConditions []ActionOperationConditionObject `json:"opconditions,omitempty"`
	OpGroup      []ActionOpGroupObject            `json:"opgroup,omitempty"`
	OpMessage    ActionOperationMessageObject     `json:"opmessage,omitempty"`
	// 消息发往哪一个用户组 若未设置 则opmessage_usr required
	OpMessageGrp []ActionOpMessageGrpObject `json:"opmessage_grp,omitempty"`
	// 消息发往哪一个用户 # todo 官方文档未提供接口范例 仅仅参数表格提到 确认下可用性
	OpMessageUsr []ActionOpMessageUsrObject `json:"opmessage_usr,omitempty"`
	OpTemplate   []ActionOpTemplateObject   `json:"optemplate,omitempty"`
	OpInventory  ActionOpInventoryObject    `json:"opinventory,omitempty"`
}

type ActionOperationCommandObject struct {
	Command    string                              `json:"command"`
	Type       ActionOperationCommandType          `json:"type"`
	AuthType   ActionOperationCommandAuthType      `json:"authtype,omitempty"`
	ExecuteOn  ActionOperationCommandExecuteOnType `json:"execute_on,omitempty"`
	Password   string                              `json:"password,omitempty"`
	Port       string                              `json:"port,omitempty"`
	PrivateKey string                              `json:"privatekey,omitempty"`
	PublicKey  string                              `json:"publickey,omitempty"`
	ScriptID   string                              `json:"scriptid,omitempty"`
	UserName   string                              `json:"username,omitempty"`
}

type ActionOperationMessageObject struct {
	// DefaultMsg  actionsDefaultMsgType  消息来源类型
	DefaultMsg  ActionOperationMessageDefaultMsgFromType `json:"default_msg,omitempty"`
	MediaTypeID string                                   `json:"mediatypeid,omitempty"`
	Message     string                                   `json:"message,omitempty"`
	Subject     string                                   `json:"subject,omitempty"`
}
type ActionOperationConditionOperatorType int

const (
	ActionOperationConditionOperatorEq ActionOperationConditionOperatorType = 0
)

type ActionOperationConditionObject struct {
	OpConditionID int `json:"opconditionid,omitempty"`
	// ConditionType  actionsCondition 过滤条件类型
	ConditionType ActionsOperationType `json:"conditiontype"`
	// Value  过滤值
	Value       string `json:"value"`
	OperationID int    `json:"operationid,omitempty"`
	// Operator ActionsOperatorActionOperationConditionOperatorType 过滤比较运算符类型
	// 过滤Filter和操作Operation是共用枚举 但是取值不一样  操作只有=
	Operator ActionOperationConditionOperatorType `json:"operator,omitempty"`
}

type ActionRecoveryOperationType int

const (
	ActionRecoveryOperationTypeSendMsg           ActionRecoveryOperationType = 0
	ActionRecoveryOperationTypeRemoteCmd                                     = 1
	ActionRecoveryOperationTypeNotifyAllInvolved                             = 11
)

type ActionRecoveryOperationObject struct {
	OperationID   int                          `json:"operationid"`
	OperationType ActionRecoveryOperationType  `json:"operationtype,omitempty"`
	ActionID      int                          `json:"actionid,omitempty"`
	OpCommand     ActionOperationCommandObject `json:"opcommand,omitempty"`
	OpCommandGrp  []ActionOpCommandGrpObject   `json:"opcommand_grp,omitempty"`
	OpCommandHst  []ActionOpCommandHostObject  `json:"opcommand_hst,omitempty"`
	OpMessage     ActionOperationMessageObject `json:"opmessage,omitempty"`
	OpMessageGrp  []ActionOpMessageGrpObject   `json:"opmessage_grp,omitempty"`
	OpMessageUsr  []ActionOpMessageUsrObject   `json:"opmessage_usr,omitempty"`
}

type ActionUpdateOperationType int

const (
	ActionUpdateOperationTypeSendMsg           ActionUpdateOperationType = 0
	ActionUpdateOperationTypeRemoteCmd                                   = 1
	ActionUpdateOperationTypeNotifyAllInvolved                           = 12
)

type ActionUpdateOperationObject struct {
	OperationID   int                          `json:"operationid"`
	OperationType ActionUpdateOperationType    `json:"operationtype,omitempty"`
	OpCommand     ActionOperationCommandObject `json:"opcommand,omitempty"`
	OpCommandGrp  []ActionOpCommandGrpObject   `json:"opcommand_grp,omitempty"`
	OpCommandHst  []ActionOpCommandHostObject  `json:"opcommand_hst,omitempty"`
	OpMessage     ActionOperationMessageObject `json:"opmessage,omitempty"`
	OpMessageGrp  []ActionOpMessageGrpObject   `json:"opmessage_grp,omitempty"`
	OpMessageUsr  []ActionOpMessageUsrObject   `json:"opmessage_usr,omitempty"`
}

//  ############################################### ActionFilterType ##############################################

type ActionFilterEvalType int

const (
	ActionFilterEvalTypeAndOr  ActionFilterEvalType = iota
	ActionFilterEvalTypeAnd                         = 1
	ActionFilterEvalTypeOr                          = 2
	ActionFilterEvalTypeCustom                      = 3
)

// ActionsFilterCondition 条件类型
type ActionsFilterCondition int

const (
	ActionFilterConditionTypeHostGroup       ActionsFilterCondition = iota //trigger actions ||  internal actions:
	ActionFilterConditionTypeHost                                          // trigger actions ||  internal actions:
	ActionFilterConditionTypeTrigger                                       //trigger actions
	ActionFilterConditionTypeTriggerName                                   //trigger actions
	ActionFilterConditionTypeTriggerSeverity                               //trigger actions
	ActionFilterConditionTypeTriggerValue
	ActionFilterConditionTypeTimePeriod                      //trigger actions
	ActionFilterConditionTypeHostIP                          //discovery actions
	ActionFilterConditionTypeDiscoveryServiceType            //discovery actions
	ActionFilterConditionTypeDiscoveryServicePort            //discovery actions
	ActionFilterConditionTypeDiscoveryStatus                 //discovery actions
	ActionFilterConditionTypeUpDownTimeDuration              //discovery actions
	ActionFilterConditionTypeRcvValue                        //discovery actions
	ActionFilterConditionTypeHostTemplate                    //trigger actions ||  internal actions:
	ActionFilterConditionTypeApplication          = iota + 1 //trigger actions ||  internal actions:
	ActionFilterConditionTypeProblemIsSuppressed             //trigger actions
	ActionFilterConditionTypeDiscRule             = iota + 2 //discovery actions
	ActionFilterConditionTypeDiscCheck                       //discovery actions
	ActionFilterConditionTypeProxy                           //discovery actions || autoregistration actions
	ActionFilterConditionTypeDiscObject                      //discovery actions
	ActionFilterConditionTypeHostName                        //autoregistration actions:
	ActionFilterConditionTypeEventType                       //internal actions
	ActionFilterConditionTypeHostMetadata                    //autoregistration actions:
	ActionFilterConditionTypeTag                             //trigger actions
	ActionFilterConditionTypeTagValue                        //trigger actions
)

type ActionFilterConditionOperatorType int

const (
	ActionFilterConditionOperatorEQ         ActionFilterConditionOperatorType = iota // =
	ActionFilterConditionOperatorNE                                                  // <>
	ActionFilterConditionOperatorContains                                            // contains
	ActionFilterConditionOperatorNotContain                                          // does not contain
	ActionFilterConditionOperatorIN                                                  // in
	ActionFilterConditionOperatorGE                                                  // >=
	ActionFilterConditionOperatorLE                                                  // <=
	ActionFilterConditionOperatorNotIn                                               // not in
	ActionFilterConditionOperatorMatches                                             // matches
	ActionFilterConditionOperatorNotMatches                                          // does not match
	ActionFilterConditionOperatorYes                                                 // yes
	ActionFilterConditionOperatorNo                                                  // no
)

type ActionFilterObject struct {
	Conditions  []ActionFilterConditionObject `json:"conditions"`
	EvalType    ActionFilterEvalType          `json:"evaltype"`
	EvalFormula string                        `json:"eval_formula,omitempty"`
	Formula     string                        `json:"formula,omitempty"`
}

type ActionFilterConditionObject struct {
	ConditionID   int                               `json:"conditionid,omitempty"`
	ConditionType ActionsFilterCondition            `json:"conditiontype"`
	Value         string                            `json:"value"`
	Value2        string                            `json:"value2,omitempty"`
	ActionID      int                               `json:"actionid,omitempty"`
	FormulaID     string                            `json:"formulaid,omitempty"`
	Operator      ActionFilterConditionOperatorType `json:"operator,omitempty"`
}

// ############################################################## ActionOpCommandType ###############################################

type ActionOpCommandGrpObject struct {
	OpCommandGrpID int `json:"opcommand_grpid,omitempty"`
	OperationID    int `json:"operationid,omitempty"`
	GroupID        int `json:"groupid,omitempty"`
}

type ActionOpCommandHostObject struct {
	OpCommandHstID int `json:"opcommand_hstid,omitempty"`
	OperationID    int `json:"operationid,omitempty"`
	HostID         int `json:"hostid,omitempty"`
}

type ActionOpGroupObject struct {
	OperationID int `json:"operationid,omitempty"`
	GroupID     int `json:"groupid,omitempty"`
}

type ActionOpMessageGrpObject struct {
	// 操作id # todo query needed
	OperationID int `json:"operationid,omitempty"`
	// 用户组id # todo query needed
	UsrGrpID int `json:"usrgrpid,omitempty"`
}

type ActionOpMessageUsrObject struct {
	OperationID int `json:"operationid,omitempty"`
	UserID      int `json:"userid,omitempty"`
}

type ActionOpTemplateObject struct {
	OperationID int `json:"operationid,omitempty"`
	TemplateID  int `json:"templateid,omitempty"`
}

type ActionOpInventoryObject struct {
	OperationID   int `json:"operationid,omitempty"`
	InventoryMode int `json:"inventory_mode,omitempty"`
}

type ActionGetParams struct {
	GetParameters

	ActionIDs    []int `json:"actionids,omitempty"`
	GroupIDs     []int `json:"groupids,omitempty"`
	HostIDs      []int `json:"hostids,omitempty"`
	TriggerIDs   []int `json:"triggerids,omitempty"`
	MediaTypeIDs []int `json:"mediatypeids,omitempty"`
	UsrGrpIDs    []int `json:"usrgrpids,omitempty"`
	UserIDs      []int `json:"userids,omitempty"`
	ScriptIDs    []int `json:"scriptids,omitempty"`

	SelectFilter             SelectQuery `json:"selectFilter,omitempty"`
	SelectOperations         SelectQuery `json:"selectOperations,omitempty"`
	SelectRecoveryOperations SelectQuery `json:"selectRecoveryOperations,omitempty"`
	SelectUpdateOperations   SelectQuery `json:"selectUpdateOperations,omitempty"`
}

type ActionCreateResult struct {
	ActionIDs []int `json:"actionids"`
}

type ActionDeleteResult struct {
	ActionIDs []int `json:"actionids"`
}

func (z *Context) ActionGet(params ActionGetParams) ([]ActionObject, int, error) {

	var result []ActionObject

	status, err := z.request("action.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Context) ActionCreate(params []ActionObject) ([]int, int, error) {

	var result ActionCreateResult

	status, err := z.request("action.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ActionIDs, status, nil
}

func (z *Context) ActionDelete(actionIDs []int) ([]int, int, error) {

	var result ActionDeleteResult

	status, err := z.request("action.delete", actionIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.ActionIDs, status, nil
}
