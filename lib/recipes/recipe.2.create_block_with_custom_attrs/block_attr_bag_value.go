package blockattrs

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func AddValuesToBlockAttributes(user response.User, blockID primitive.ObjectID) error {
	log.Info("TODO(1): Add values to the block attributes.")

	return nil
}

func FetchBlockAttributes(user response.User) error {
	log.Info("TODO(1): Fetch all block custom attributes along with their values.")

	return nil
}
