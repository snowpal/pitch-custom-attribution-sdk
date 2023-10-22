package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type AttrBag struct {
	ID         string             `json:"ID"`
	Name       string             `json:"name"`
	Attributes []common.Attribute `json:"attributes"`
	common.UserMetadata
}

type AttrBags struct {
	AttrBags []AttrBag `json:"bags"`
}
