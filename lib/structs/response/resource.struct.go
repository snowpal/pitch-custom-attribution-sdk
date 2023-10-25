package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

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
