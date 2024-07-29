package zabbix

// For `TemplateGetParams` field: `EvalType`
const (
	TemplateEvalTypeAndOr = 0
	TemplateEvalTypeOr    = 2
)

// For `TemplateTag` field: `Operator`
const (
	TemplateTagOperatorContains = 0
	TemplateTagOperatorEquals   = 1
)

// TemplateObject 告警模板对象
type TemplateObject struct {
	//模板ID。
	TemplateID int `json:"templateid,omitempty"`
	//模板的正式名称。
	Host string `json:"host,omitempty"`
	//模板说明。
	Description string `json:"description,omitempty"`
	//主机的可见名称。默认：主机的属性值。
	Name string `json:"name,omitempty"`
	//模板关联的主机组
	Groups []HostGroupObject `json:"groups,omitempty"`
	//模板标签
	Tags      []TemplateTagObject `json:"tags,omitempty"`
	Templates []TemplateObject    `json:"templates,omitempty"`
	//链接的父模板
	ParentTemplates []TemplateObject `json:"parentTemplates,omitempty"`
	//宏
	Macros []UserMacroObject `json:"macros,omitempty"`
	//模板关联的主机
	Hosts []HostObject `json:"hosts,omitempty"`
}

// TemplateTagObject struct is used to store template tag data
type TemplateTagObject struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// TemplateGetParams struct is used for template get requests
type TemplateGetParams struct {
	GetParameters

	TemplateIDs       []int `json:"templateids,omitempty"`
	GroupIDs          []int `json:"groupids,omitempty"`
	ParentTemplateIDs []int `json:"parentTemplateids,omitempty"`
	HostIDs           []int `json:"hostids,omitempty"`
	GraphIDs          []int `json:"graphids,omitempty"`
	ItemIDs           []int `json:"itemids,omitempty"`
	TriggerIDs        []int `json:"triggerids,omitempty"`

	WithItems     bool                `json:"with_items,omitempty"`
	WithTriggers  bool                `json:"with_triggers,omitempty"`
	WithGraphs    bool                `json:"with_graphs,omitempty"`
	WithHttpTests bool                `json:"with_httptests,omitempty"`
	EvalType      int                 `json:"evaltype,omitempty"` // has defined consts, see above
	Tags          []TemplateTagObject `json:"tags,omitempty"`

	SelectGroups          SelectQuery `json:"selectGroups,omitempty"`
	SelectTags            SelectQuery `json:"selectTags,omitempty"`
	SelectHosts           SelectQuery `json:"selectHosts,omitempty"`
	SelectTemplates       SelectQuery `json:"selectTemplates,omitempty"`
	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	SelectMacros          SelectQuery `json:"selectMacros,omitempty"`

	// SelectHttpTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	// SelectItems           SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	// SelectApplications    SelectQuery `json:"selectApplications,omitempty"` // not implemented yet
	// SelectScreens         SelectQuery `json:"selectScreens,omitempty"` // not implemented yet
}

// Structure to store creation result
type templateCreateResult struct {
	TemplateIDs []int `json:"templateids"`
}

// Structure to update deletion result
type templateUpdateResult struct {
	TemplateIDs []int `json:"templateids"`
}

// Structure to store deletion result
type templateDeleteResult struct {
	TemplateIDs []int `json:"templateids"`
}

// TemplateGet gets templates
func (z *Context) TemplateGet(params TemplateGetParams) ([]TemplateObject, int, error) {

	var result []TemplateObject

	status, err := z.request("template.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

// TemplateCreate creates templates
func (z *Context) TemplateCreate(params []TemplateObject) ([]int, int, error) {

	var result templateCreateResult

	status, err := z.request("template.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}

// TemplateUpdate update templates
func (z *Context) TemplateUpdate(params []TemplateObject) ([]int, int, error) {

	var result templateUpdateResult

	status, err := z.request("template.update", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}

// TemplateDelete deletes templates
func (z *Context) TemplateDelete(templateIDs []int) ([]int, int, error) {

	var result templateDeleteResult

	status, err := z.request("template.delete", templateIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.TemplateIDs, status, nil
}
