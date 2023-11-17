package blockattrs

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func CreateBlockAttributes(user response.User) (map[int]primitive.ObjectID, error) {
	log.Info("Test Case: Creating custtom attributes for a block.")
	attrIDs := make(map[int]primitive.ObjectID)

	var err error
	var attrID primitive.ObjectID
	if attrID, err = recipes.CreateTextAttribute(user, BlockAttrs[BlockID].Name, false); err != nil {
		return attrIDs, err
	}
	attrIDs[BlockID] = attrID

	if attrID, err = recipes.CreateTextAttribute(user, BlockAttrs[Name].Name, false); err != nil {
		return attrIDs, err
	}
	attrIDs[Name] = attrID

	if attrID, err = recipes.CreateTextAttribute(
		user,
		BlockAttrs[Description].Name,
		false,
	); err != nil {
		return attrIDs, err
	}
	attrIDs[Description] = attrID

	if attrID, err = recipes.CreateTextAttribute(
		user,
		BlockAttrs[SimpleDescription].Name,
		false,
	); err != nil {
		return attrIDs, err
	}
	attrIDs[SimpleDescription] = attrID

	if attrID, err = recipes.CreateTextAttribute(user, BlockAttrs[Tags].Name, true); err != nil {
		return attrIDs, err
	}
	attrIDs[Tags] = attrID

	if attrID, err = recipes.CreateTextAttribute(user, BlockAttrs[Flag].Name, false); err != nil {
		return attrIDs, err
	}
	attrIDs[Flag] = attrID

	if attrID, err = createBlockType(user); err != nil {
		return attrIDs, err
	}
	attrIDs[BlockType] = attrID

	if attrID, err = createScaleAndScaleValue(user); err != nil {
		return attrIDs, err
	}
	attrIDs[ScaleAndScaleValue] = attrID

	if attrID, err = recipes.CreateDateAttribute(user, BlockAttrs[DueDate].Name); err != nil {
		return attrIDs, err
	}
	attrIDs[DueDate] = attrID

	if attrID, err = recipes.CreateDateAttribute(user, BlockAttrs[StartDate].Name); err != nil {
		return attrIDs, err
	}
	attrIDs[StartDate] = attrID

	if attrID, err = recipes.CreateDateAttribute(user, BlockAttrs[EndDate].Name); err != nil {
		return attrIDs, err
	}
	attrIDs[EndDate] = attrID

	return attrIDs, nil
}

func createBlockType(user response.User) (primitive.ObjectID, error) {
	log.Info("Creating block type for a block.")
	var attributeID primitive.ObjectID
	var err error

	attributeID, err = recipes.CreateSingleSelectAttribute(
		user,
		BlockAttrs[BlockType].Name,
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

	attributeID, err = recipes.CreateNSSAttribute(
		user,
		BlockAttrs[ScaleAndScaleValue].Name,
		[]common.NSSOption{
			{
				ID:    1,
				Value: "Grades",
				Child: &common.NSSChild{
					Name:      BlockAttrs[ScaleAndScaleValue].ChildName,
					FieldType: "SingleSelect",
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
							Value: "C",
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
