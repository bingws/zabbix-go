package zabbix

import "fmt"

const (
	HistoryObjectTypeFloat           = 0
	HistoryObjectTypeCharacter       = 1
	HistoryObjectTypeLog             = 2
	HistoryObjectTypeNumericUnsigned = 3
	HistoryObjectTypeText            = 4
)

type HistoryFloatObject struct {
	Clock  int     `json:"clock,omitempty"`
	ItemID int     `json:"itemid,omitempty"`
	NS     int     `json:"ns,omitempty"`
	Value  float64 `json:"value,omitempty"`
}

type HistoryIntegerObject struct {
	Clock  int `json:"clock,omitempty"`
	ItemID int `json:"itemid,omitempty"`
	NS     int `json:"ns,omitempty"`
	Value  int `json:"value,omitempty"`
}

type HistoryStringObject struct {
	Clock  int    `json:"clock,omitempty"`
	ItemID int    `json:"itemid,omitempty"`
	NS     int    `json:"ns,omitempty"`
	Value  string `json:"value,omitempty"`
}

type HistoryTextObject struct {
	ID     int    `json:"id,omitempty"`
	Clock  int    `json:"clock,omitempty"`
	ItemID int    `json:"itemid,omitempty"`
	NS     int    `json:"ns,omitempty"`
	Value  string `json:"value,omitempty"`
}

type HistoryLogObject struct {
	ID         int    `json:"id,omitempty"`
	Clock      int    `json:"clock,omitempty"`
	ItemID     int    `json:"itemid,omitempty"`
	LogeventID int    `json:"logeventid,omitempty"`
	NS         int    `json:"ns,omitempty"`
	Severity   int    `json:"severity,omitempty"`
	Source     int    `json:"source,omitempty"`
	Timestamp  int    `json:"timestamp,omitempty"`
	Value      string `json:"value,omitempty"`
}

type HistoryGetParams struct {
	GetParameters

	History  int   `json:"history"`
	HostIDs  []int `json:"hostids,omitempty"`
	ItemIDs  []int `json:"itemids,omitempty"`
	TimeFrom int   `json:"time_from,omitempty"`
	TimeTill int   `json:"time_till,omitempty"`

	Sortfield string `json:"sortfield,omitempty"`
}

func (z *Context) HistoryGet(params HistoryGetParams) (interface{}, int, error) {

	var result interface{}

	switch params.History {
	case HistoryObjectTypeFloat:
		result = &([]HistoryFloatObject{})
	case HistoryObjectTypeCharacter:
		result = &([]HistoryStringObject{})
	case HistoryObjectTypeLog:
		result = &([]HistoryLogObject{})
	case HistoryObjectTypeNumericUnsigned:
		result = &([]HistoryIntegerObject{})
	case HistoryObjectTypeText:
		result = &([]HistoryTextObject{})
	default:
		return nil, 0, fmt.Errorf("Unknown history type")
	}

	status, err := z.request("history.get", params, result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
