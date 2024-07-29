package zabbix

// 允许自动登录。  可能的值: 0 - *(default)* 禁止自动登录; 1 -允许自动登录。
const (
	UserAutoLoginDisabled = 0
	UserAutoLoginEnabled  = 1
)

// 用户的主题。  可能的值: `default` - *(default)* system default; `blue-theme` - Blue; `dark-theme` - Dark.
const (
	UserThemeDefault = "default"
	UserThemeBlue    = "blue-theme"
	UserThemeDark    = "dark-theme"
)

// 用户类型。  可能的值: 1 - *(default)* Zabbix user; 2 - Zabbix admin; 3 - Zabbix super admin.
const (
	UserTypeUser       = 1
	UserTypeAdmin      = 2
	UserTypeSuperAdmin = 3
)

// 是否启用媒体。  可能的值: 0 - *(默认)* enabled; 1 - disabled.
const (
	MediaActiveEnabled  = 0
	MediaActiveDisabled = 1
)

type UserObject struct {
	//用户的ID。
	UserID int `json:"userid,omitempty"`
	//用户别名。
	Alias string `json:"alias,omitempty"`
	//最近一次登录失败的时间。
	AttemptClock int `json:"attempt_clock,omitempty"`
	//最近失败的登录尝试次数。
	AttemptFailed int `json:"attempt_failed,omitempty"`
	//最近一次失败的登录来源IP地址。
	AttemptIP string `json:"attempt_ip,omitempty"`
	//允许自动登录。  可能的值: 0 - *(default)* 禁止自动登录; 1 -允许自动登录。
	AutoLogin int `json:"autologin,omitempty"`
	//会话过期时间。 接受具有后缀的秒或时间单位。 如果设置为 0s, 用户登录会话永远不会过期。  默认: 15m.
	AutoLogout string `json:"autologout"`
	//用户默认语言代码  默认: `en_GB`。
	Lang string `json:"lang,omitempty"`
	//用户名.
	Name string `json:"name,omitempty"`
	//自动刷新时间间隔. 接受具有后缀的秒或时间单位。  默认: 30s.
	Refresh string `json:"refresh,omitempty"`
	//每页显示的对象行数。  默认: 50.
	RowsPerPage int `json:"rows_per_page,omitempty"`
	//姓。
	Surname string `json:"surname,omitempty"`
	//用户的主题。  可能的值: `default` - *(default)* system default; `blue-theme` - Blue; `dark-theme` - Dark.
	Theme string `json:"theme,omitempty"`
	//用户类型。  可能的值: 1 - *(default)* Zabbix user; 2 - Zabbix admin; 3 - Zabbix super admin.
	Type int `json:"type,omitempty"`
	//在登录后将用户重定向到页面的URL。
	URL string `json:"url,omitempty"`

	//  user.login
	UserDataObject

	//用户媒介
	Medias []MediaObject `json:"medias,omitempty"`
	//媒介类型对象
	MediaTypes []MediaTypeObject `json:"mediatypes,omitempty"`
	//所属用户组
	UsrGrps []UserGroupObject `json:"usrgrps,omitempty"`

	// 新用户创建后使用
	UserMedias []MediaObject `json:"user_medias,omitempty"`
	Passwd     string        `json:"passwd,omitempty"`
}

// 媒介对象
type MediaObject struct {
	//媒介ID
	MediaID int `json:"mediaid,omitempty"`
	//用于媒介的媒介类型ID
	MediaTypeID int `json:"mediatypeid,omitempty"`
	//地址, 用户名或者接收方的其他标识符。  如果类型是电子邮件, 值被设置为数组。 其他类型值被设置为字符串。
	SendTo []string `json:"sendto,omitempty"`
	//是否启用媒体。  可能的值: 0 - *(默认)* enabled; 1 - disabled.
	Active int `json:"active,omitempty"`
	//触发发送通知告警级别。  严重性以二进制形式存储，每一位表示相应的严重性。例如，12在二进制中等于1100，这意味着通知将由具有警告和平均级别的触发器发送
	Severity int `json:"severity,omitempty"`
	//当通知可以作为时间段发送或者用分号隔开用户宏。  默认: 1-7,00:00-24:00
	Period string `json:"period,omitempty"`
}

type UserLoginParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
	UserData string `json:"userData,omitempty"`
}

type UserDataObject struct {
	DebugMode bool   `json:"debug_mode,omitempty"`
	GUIAccess int    `json:"gui_access,omitempty"`
	SessionID string `json:"sessionid,omitempty"`
	UserIP    string `json:"userip,omitempty"`
}

type UserGetParams struct {
	GetParameters

	MediaIDs     []int `json:"mediaids,omitempty"`
	NediaTypeIDs []int `json:"mediatypeids,omitempty"`
	UserIDs      []int `json:"userids,omitempty"`
	UsrGrpIDs    []int `json:"usrgrpids,omitempty"`

	GetAccess        bool        `json:"getAccess,omitempty"`
	SelectMedias     SelectQuery `json:"selectMedias,omitempty"`
	SelectMediaTypes SelectQuery `json:"selectMediatypes,omitempty"`
	SelectUsrGrps    SelectQuery `json:"selectUsrgrps,omitempty"`
}

type userCreateResult struct {
	UserIDs []int `json:"userids"`
}

type userDeleteResult struct {
	UserIDs []int `json:"userids"`
}

func (z *Context) UserGet(params UserGetParams) ([]UserObject, int, error) {

	var result []UserObject

	status, err := z.request("user.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Context) UserCreate(params []UserObject) ([]int, int, error) {

	var result userCreateResult

	status, err := z.request("user.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UserIDs, status, nil
}

func (z *Context) UserDelete(userIDs []int) ([]int, int, error) {

	var result userDeleteResult

	status, err := z.request("user.delete", userIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UserIDs, status, nil
}

func (z *Context) userLogin(params UserLoginParams) (string, int, error) {

	var result string

	status, err := z.request("user.login", params, &result)
	if err != nil {
		return "", status, err
	}

	return result, status, nil
}

func (z *Context) userLogout() (bool, int, error) {

	var result bool

	status, err := z.request("user.logout", []string{}, &result)
	if err != nil {
		return false, status, err
	}

	return result, status, nil
}
