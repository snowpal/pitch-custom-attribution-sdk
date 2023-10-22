package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type PrimitiveAttr struct {
	ID        string `json:"ID"`
	Name      string `json:"name"`
	FieldType string `json:"fieldType"`
	Label     *bool  `json:"label,omitempty"`
	common.UserMetadata
}

type SingleSelectAttr struct {
	ID        string                `json:"ID"`
	Name      string                `json:"name"`
	FieldType string                `json:"fieldType"`
	Options   []common.SelectOption `json:"options"`
	common.UserMetadata
}

type MultiselectAttr = SingleSelectAttr

type NSSAttr struct {
	ID        string             `json:"ID"`
	Name      string             `json:"name"`
	FieldType string             `json:"fieldType"`
	Options   []common.NSSOption `json:"options"`
	common.UserMetadata
}

type Attribute struct {
	*PrimitiveAttr
	*SingleSelectAttr
	*MultiselectAttr
	*NSSAttr
}

type Attributes struct {
	Attributes []Attribute `json:"attributes"`
}
