package request

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type PrimitiveAttrReq struct {
	Name  string `json:"name"`
	Label *bool  `json:"label,omitempty"`
}

type SelectAttrReq struct {
	Name    string                `json:"name"`
	Options []common.SelectOption `json:"options"`
}

type MultiselectAttrReq = SelectAttrReq

type NSSAttrReq struct {
	Name    string             `json:"name"`
	Options []common.NSSOption `json:"options"`
}

type AttributeReq struct {
	*PrimitiveAttrReq
	*SelectAttrReq
	*MultiselectAttrReq
	*NSSAttrReq
}
