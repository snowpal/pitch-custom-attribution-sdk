package response

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PrimitiveAttr struct {
	ID        primitive.ObjectID `json:"ID"`
	Name      string             `json:"name"`
	FieldType string             `json:"fieldType"`
	Label     *bool              `json:"label,omitempty"`
	Active    bool               `json:"active"`
	common.UserMetadata
}

type SingleSelectAttr struct {
	ID        primitive.ObjectID    `json:"ID"`
	Name      string                `json:"name"`
	FieldType string                `json:"fieldType"`
	Options   []common.SelectOption `json:"options"`
	common.UserMetadata
}

type MultiselectAttr = SingleSelectAttr

type NSSAttr struct {
	ID        primitive.ObjectID `json:"ID"`
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
	ID            primitive.ObjectID `json:"ID"`
	Name          string             `json:"name"`
	FieldType     string             `json:"fieldType"`
	LeafFieldType *string            `json:"leafFieldType,omitempty"`
	Label         *bool              `json:"label,omitempty"`
	Active        bool               `json:"active"`
	Options       []AttributeOption  `json:"options,omitempty"`
	common.UserMetadata
}

type Attributes struct {
	Attributes []Attribute `json:"attributes"`
}
