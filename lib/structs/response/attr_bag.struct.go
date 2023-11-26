package response

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttrBag struct {
	ID         primitive.ObjectID    `json:"ID"`
	Name       string                `json:"name"`
	Attributes []common.BagAttribute `json:"attributes"`
	common.UserMetadata
}

type AttrBags struct {
	AttrBags []AttrBag `json:"bags"`
}
