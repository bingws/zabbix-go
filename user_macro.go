package zabbix

// For `UserMacroObject` field: `Type`
const (
	UserMacroTypeText   = 0
	UserMacroTypeSecret = 1
)

// UserMacroObject 用户宏对象
type UserMacroObject struct {

	// global macro fields only
	GlobalMacroID int `json:"globalMacroid,omitempty"`

	// Host macro fields only
	HostMacroID int `json:"hostMacroid,omitempty"`
	HostID      int `json:"hostid,omitempty"`

	// Common fields
	Macro       string `json:"macro,omitempty"`
	Value       string `json:"value,omitempty"`
	Type        int    `json:"type,omitempty"` // has defined consts, see above
	Description string `json:"description,omitempty"`

	Groups    []HostGroupObject `json:"groups,omitempty"`
	Hosts     []HostObject      `json:"hosts,omitempty"`
	Templates []TemplateObject  `json:"templates,omitempty"`
}

// UserMacroGetParams struct is used for hostMacro get requests
type UserMacroGetParams struct {
	GetParameters

	GlobalMacro    bool  `json:"globalMacro,omitempty"`
	GlobalMacroIDs []int `json:"globalMacroids,omitempty"`
	GroupIDs       []int `json:"groupids,omitempty"`
	HostIDs        []int `json:"hostids,omitempty"`
	HostMacroIDs   []int `json:"hostMacroids,omitempty"`
	TemplateIDs    []int `json:"templateids,omitempty"`

	SelectGroups    SelectQuery `json:"selectGroups,omitempty"`
	SelectHosts     SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates SelectQuery `json:"selectTemplates,omitempty"`
}

// Structure to store creation result
type hostMacroCreateResult struct {
	HostMacroIDs []int `json:"hostmacroids"`
}

// Structure to store creation global macros result
type globalMacroCreateResult struct {
	GlobalMacroIDs []int `json:"globalmacroids"`
}

// Structure to store deletion result
type hostMacroDeleteResult struct {
	HostMacroIDs []int `json:"hostmacroids"`
}

// Structure to store deletion global macros result
type globalMacroDeleteResult struct {
	GlobalMacroIDs []int `json:"globalmacroids"`
}

// UserMacroGet gets global or host macros according to the given parameters
func (z *Context) UserMacroGet(params UserMacroGetParams) ([]UserMacroObject, int, error) {

	var result []UserMacroObject

	status, err := z.request("usermacro.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// HostMacroCreate creates new hostMacros
func (z *Context) HostMacroCreate(params []UserMacroObject) ([]int, int, error) {

	var result hostMacroCreateResult

	status, err := z.request("usermacro.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostMacroIDs, status, nil
}

// GlobalMacroCreate creates new globalMacros
func (z *Context) GlobalMacroCreate(params []UserMacroObject) ([]int, int, error) {

	var result globalMacroCreateResult

	status, err := z.request("usermacro.createglobal", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GlobalMacroIDs, status, nil
}

// HostMacroDelete deletes hostMacros
func (z *Context) HostMacroDelete(hostMacroIDs []int) ([]int, int, error) {

	var result hostMacroDeleteResult

	status, err := z.request("usermacro.delete", hostMacroIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.HostMacroIDs, status, nil
}

// GlobalMacroDelete deletes globalMacros
func (z *Context) GlobalMacroDelete(globalMacroIDs []int) ([]int, int, error) {

	var result globalMacroDeleteResult

	status, err := z.request("usermacro.deleteglobal", globalMacroIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.GlobalMacroIDs, status, nil
}
