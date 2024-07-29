package zabbix

const (
	MediaTypeEmail   = 0
	MediaTypeScript  = 1
	MediaTypeSMS     = 2
	MediaTypeWebhook = 4
)

const (
	MediaTypeSMTPSecurityNone     = 0
	MediaTypeSMTPSecuritySTARTTLS = 1
	MediaTypeSMTPSecuritySSLTLS   = 2
)

const (
	MediaTypeSMTPVerifyHostNo  = 0
	MediaTypeSMTPVerifyHostYes = 1
)

const (
	MediaTypeSMTPVerifyPeerNo  = 0
	MediaTypeSMTPVerifyPeerYes = 1
)

const (
	MediaTypeSMTPAuthenticationNone           = 0
	MediaTypeSMTPAuthenticationNormalPassword = 1
)

const (
	MediaTypeStatusEnabled  = 0
	MediaTypeScriptDisabled = 1
)

const (
	MediaTypeContentTypePlainText = 0
	MediaTypeContentTypeHTML      = 1
)

const (
	MediaTypeProcessTagsNo  = 0
	MediaTypeProcessTagsYes = 1
)

const (
	MediaTypeShowEventMenuNo  = 0
	MediaTypeShowEventMenuYes = 1
)

const (
	MediaTypeMessageTemplateEventSourceTriggers         = 0
	MediaTypeMessageTemplateEventSourceDiscovery        = 1
	MediaTypeMessageTemplateEventSourceAutoregistration = 2
	MediaTypeMessageTemplateEventSourceInternal         = 3
)

const (
	MediaTypeMessageTemplateRecoveryOperations         = 0
	MediaTypeMessageTemplateRecoveryRecoveryOperations = 1
	MediaTypeMessageTemplateRecoveryUpdateOperations   = 2
)

// MediaTypeObject 媒介类型对象
type MediaTypeObject struct {
	MediaTypeID        int                                `json:"mediatypeid,omitempty"`
	Name               string                             `json:"name,omitempty"`
	Type               int                                `json:"type,omitempty"`
	ExecPath           string                             `json:"exec_path,omitempty"`
	GsmModem           string                             `json:"gsm_modem,omitempty"`
	Passwd             string                             `json:"passwd,omitempty"`
	SMTPEmail          string                             `json:"smtp_email,omitempty"`
	SMTPHelo           string                             `json:"smtp_helo,omitempty"`
	SMTPServer         string                             `json:"smtp_server,omitempty"`
	SMTPPort           int                                `json:"smtp_port,omitempty"`
	SMTPSecurity       int                                `json:"smtp_security,omitempty"`
	SMTPVerifyHost     int                                `json:"smtp_verify_host,omitempty"`
	SMTPVerifyPeer     int                                `json:"smtp_verify_peer,omitempty"`
	SMTPAuthentication int                                `json:"smtp_authentication,omitempty"`
	Status             int                                `json:"status,omitempty"`
	Username           string                             `json:"username,omitempty"`
	ExecParams         string                             `json:"exec_params,omitempty"`
	MaxSessions        int                                `json:"maxsessions,omitempty"`
	MaxAttempts        int                                `json:"maxattempts,omitempty"`
	AttemptInterval    string                             `json:"attempt_interval,omitempty"`
	ContentType        int                                `json:"content_type,omitempty"`
	Script             string                             `json:"script,omitempty"`
	Timeout            string                             `json:"timeout,omitempty"`
	ProcessTags        int                                `json:"process_tags,omitempty"`
	ShowEventMenu      int                                `json:"show_event_menu,omitempty"`
	EventMenuURL       string                             `json:"event_menu_url,omitempty"`
	EventMenuName      string                             `json:"event_menu_name,omitempty"`
	Parameters         []MediaTypeWebhookParametersObject `json:"parameters,omitempty"`
	Description        string                             `json:"description,omitempty"`
	MessageTemplates   []MediaTypeMessageTemplateObject   `json:"message_templates,omitempty"`

	Users []UserObject `json:"users,omitempty"`
}

type MediaTypeWebhookParametersObject struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type MediaTypeMessageTemplateObject struct {
	EventSource int    `json:"eventsource"`
	Recovery    int    `json:"recovery"`
	Subject     string `json:"subject,omitempty"`
	Message     string `json:"message,omitempty"`
}

type MediaTypeGetParams struct {
	GetParameters

	MediaTypeIDs []int `json:"mediatypeids,omitempty"`
	MediaIDs     []int `json:"mediaids,omitempty"`
	UserIDs      []int `json:"userids,omitempty"`

	SelectMessageTemplates SelectQuery `json:"selectMessageTemplates,omitempty"`
	SelectUsers            SelectQuery `json:"selectUsers,omitempty"`
}

type mediatypeCreateResult struct {
	MediaTypeIDs []int `json:"mediatypeids"`
}

type mediatypeDeleteResult struct {
	MediaTypeIDs []int `json:"mediatypeids"`
}

func (z *Context) MediaTypeGet(params MediaTypeGetParams) ([]MediaTypeObject, int, error) {

	var result []MediaTypeObject

	status, err := z.request("mediatype.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}

func (z *Context) MediaTypeCreate(params []MediaTypeObject) ([]int, int, error) {

	var result mediatypeCreateResult

	status, err := z.request("mediatype.create", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result.MediaTypeIDs, status, nil
}

func (z *Context) MediaTypeDelete(mediatypeIDs []int) ([]int, int, error) {

	var result mediatypeDeleteResult

	status, err := z.request("mediatype.delete", mediatypeIDs, &result)
	if err != nil {
		return nil, status, err
	}

	return result.MediaTypeIDs, status, nil
}
