package zabbix

const (
	UserGroupDebugModeDisabled = 0
	UserGroupDebugModeEnabled  = 1
)

const (
	UserGroupGuiAccessSystemDefaultAuth = 0
	UserGroupGuiAccessInternalAuth      = 1
	UserGroupGuiAccessLDAPAuth          = 2
	UserGroupGuiAccessDisableFrontend   = 3
)

const (
	UserGroupUsersStatusEnabled  = 0
	UserGroupUsersStatusDisabled = 1
)

const (
	UserGroupPermissionDenied = 0
	UserGroupPermissionRO     = 2
	UserGroupPermissionRW     = 3
)

type UserGroupObject struct {
	UsrGrpID    int    `json:"usrgrpid,omitempty"`
	Name        string `json:"name,omitempty"`
	DebugMode   int    `json:"debug_mode,omitempty"`
	GuiAccess   int    `json:"gui_access,omitempty"`
	UsersStatus int    `json:"users_status,omitempty"`

	Users      []UserObject                        `json:"users,omitempty"`
	Rights     []UserGroupPermissionObject         `json:"rights,omitempty"`
	TagFilters []UserGroupTagBasedPermissionObject `json:"tag_filters,omitempty"`

	// 已有新用户组创建或更新 后使用
	UserIDs []int `json:"userids,omitempty"`
}

type UserGroupPermissionObject struct {
	ID         int `json:"id"`
	Permission int `json:"permission"`
}

type UserGroupTagBasedPermissionObject struct {
	GroupID int    `json:"groupid,omitempty"`
	Tag     string `json:"tag,omitempty"`
	Value   string `json:"value,omitempty"`
}

type UserGroupGetParams struct {
	GetParameters

	Status        []int `json:"status,omitempty"`
	UserIDs       []int `json:"userids,omitempty"`
	UsrGrpIDs     []int `json:"usrgrpids,omitempty"`
	WithGuiAccess []int `json:"with_gui_access,omitempty"`

	SelectTagFilters SelectQuery `json:"selectTagFilters,omitempty"`
	SelectUsers      SelectQuery `json:"selectUsers,omitempty"`
	SelectRights     SelectQuery `json:"selectRights,omitempty"`
}

type userGroupCreateResult struct {
	UsrGrpIDs []int `json:"usrgrpids"`
}

type userGroupUpdateResult struct {
	UsrGrpIDs []int `json:"usrgrpids"`
}

type userGroupDeleteResult struct {
	UsrGrpIDs []int `json:"usrgrpids"`
}

func (z *Context) UserGroupGet(params UserGroupGetParams) ([]UserGroupObject, int, error) {

	var result []UserGroupObject

	status, err := z.request("usergroup.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Context) UserGroupCreate(params []UserGroupObject) ([]int, int, error) {

	var result userGroupCreateResult

	status, err := z.request("usergroup.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrGrpIDs, status, nil
}

func (z *Context) UserGroupUpdate(params []UserGroupObject) ([]int, int, error) {

	var result userGroupUpdateResult

	status, err := z.request("usergroup.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrGrpIDs, status, nil
}

func (z *Context) UserGroupDelete(userGroupIDs []int) ([]int, int, error) {

	var result userGroupDeleteResult

	status, err := z.request("usergroup.delete", userGroupIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.UsrGrpIDs, status, nil
}
