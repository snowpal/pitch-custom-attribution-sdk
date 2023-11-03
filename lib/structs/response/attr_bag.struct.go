package response

import "github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"

type AttrBag struct {
	ID         string                `json:"ID"`
	Name       string                `json:"name"`
	Attributes []common.BagAttribute `json:"attributes"`
	common.UserMetadata
}

type AttrBags struct {
	AttrBags []AttrBag `json:"bags"`
}
