package blockattrs


import (
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func CreateBlockAttributes(user response.User) ([]primitive.ObjectID, error) {
	log.Info("TODO(1): Creating custtom attributes for a block.")
	var attributeIDs []primitive.ObjectID

	return attributeIDs, nil
}
