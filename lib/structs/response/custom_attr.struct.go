package response

import "github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"

type PrimitiveAttr struct {
	ID        string `json:"ID"`
	Name      string `json:"name"`
	FieldType string `json:"fieldType"`
	Label     *bool  `json:"label,omitempty"`
	Active    bool   `json:"active"`
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

type AttributeOption struct {
	*common.SelectOption
	*common.NSSOption
}

type Attribute struct {
	ID            string            `json:"ID"`
	Name          string            `json:"name"`
	FieldType     string            `json:"fieldType"`
	LeafFieldType *string           `json:"leafFieldType,omitempty"`
	Label         *bool             `json:"label,omitempty"`
	Active        bool              `json:"active"`
	Options       []AttributeOption `json:"options,omitempty"`
	common.UserMetadata
}

type Attributes struct {
	Attributes []Attribute `json:"attributes"`
}
