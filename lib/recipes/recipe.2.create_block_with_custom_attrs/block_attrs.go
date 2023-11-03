package blockattrs

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func CreateBlockAttributes(user response.User) ([]primitive.ObjectID, error) {
	log.Info("TODO(1): Creating custtom attributes for a block.")
	var attributeIDs []primitive.ObjectID

	createTextAttribute(user, "Block ID", false)
	createTextAttribute(user, "Name", false)
	createTextAttribute(user, "Description", false)
	createTextAttribute(user, "Simple Description", false)
	createTextAttribute(user, "Tags", true)
	createTextAttribute(user, "Flag", false)
	createBlockType(user)
	createScaleAndScaleValue(user)
	createDateAttribute(user, "Due Date")
	createDateAttribute(user, "Start Date")
	createDateAttribute(user, "End Date")

	return attributeIDs, nil
}

func createBlockType(user response.User) (primitive.ObjectID, error) {
	log.Info("Creating block type for a block.")
	var attributeID primitive.ObjectID
	var err error

	attributeID, err = createSingleSelectAttribute(
		user,
		"Block Type",
		[]common.SelectOption{
			{
				ID:    1,
				Value: "Basic",
			},
			{
				ID:    2,
				Value: "Advanced Placement",
			},
			{
				ID:    3,
				Value: "Honors",
			},
		})
	if err != nil {
		return attributeID, err
	}

	return attributeID, nil
}

func createScaleAndScaleValue(user response.User) (primitive.ObjectID, error) {
	log.Info("Creating scale and scale value for a block.")
	var attributeID primitive.ObjectID
	var err error

	attributeID, err = createNSSAttribute(user, "Scale", []common.NSSOption{
		{
			ID:    1,
			Value: "Grades",
			Child: &common.NSSChild{
				Name:      "Scale Value",
				FieldType: "Single Select",
				Options: &[]common.NSSOption{
					{
						ID:    1,
						Value: "A",
					},
					{
						ID:    2,
						Value: "B",
					},
					{
						ID:    3,
						Value: "F",
					},
				},
			},
		},
	})
	if err != nil {
		return attributeID, err
	}

	return attributeID, nil
}

func createTextAttribute(user response.User, attrName string, label bool) (primitive.ObjectID,
	error) {
	log.Info("Creating a text attribute for a block.")
	var attributeID primitive.ObjectID

	return attributeID, nil
}

func createSingleSelectAttribute(user response.User, singleSelectName string,
	options []common.SelectOption) (primitive.ObjectID, error) {
	log.Info("Creating single select for a block.")
	var attributeID primitive.ObjectID

	return attributeID, nil
}

func createDateAttribute(user response.User, attrName string) (primitive.ObjectID,
	error) {
	log.Info("Creating a date attribute for a block.")
	var attributeID primitive.ObjectID

	return attributeID, nil
}

func createNSSAttribute(user response.User, attrName string,
	options []common.NSSOption) (primitive.ObjectID,
	error) {
	log.Info("Creating a number attribute for a block.")
	var attributeID primitive.ObjectID

	return attributeID, nil
}
