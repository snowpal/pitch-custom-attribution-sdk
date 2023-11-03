package blockattrs

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func CreateAttributeBagForBlock(user response.User, attrIDs []primitive.ObjectID) (response.AttrBag,
	error) {
	log.Info("TODO(1): Creating attribute bag for a block with the given attribute IDs.")
	var attrBag response.AttrBag

	return attrBag, nil
}

func AssiciateBagWithBlock(user response.User, blockID primitive.ObjectID,
	attrBag response.AttrBag) error {
	log.Info("TODO(1): Associate the attribute bag with block.")

	return nil
}
