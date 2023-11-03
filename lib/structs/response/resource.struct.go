package response

import "github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"

type Resource struct {
	ID         string            `json:"ID"`
	ResourceID string            `json:"resourceID"`
	Bags       []ResourceAttrBag `json:"bags"`
	common.UserMetadata
}

type ResourceAttrBag struct {
	ID     string `json:"ID"`
	Active bool   `json:"active"`
}

type ResourceAttrBagValue struct {
	ID     string                   `json:"ID"`
	Text   *common.BagPrimitiveAttr `json:"text"`
	Number *common.BagPrimitiveAttr `json:"number"`
	Date   *common.BagPrimitiveAttr `json:"date"`
	File   *common.BagPrimitiveAttr `json:"file"`
	Option *common.BagSelectOption  `json:"option"`
	common.UserMetadata
}

type ResourceAttrBagValues struct {
	Attributes []ResourceAttrBagValue `json:"attributes"`
}
