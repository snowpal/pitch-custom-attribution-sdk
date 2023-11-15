package blockattrs

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/endpoints/resources"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/request"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
	attributesbag "github.com/snowpal/pitch-custom-attribution-sdk/lib/endpoints/attributes_bag.1"
)

func getAttrbitueData(attrID primitive.ObjectID, attr int) common.BagAttribute {
	return common.BagAttribute{
		ID:       attrID.Hex(),
		Required: BlockAttrs[attr].Required,
		List:     BlockAttrs[attr].List,
		Metadata: BlockAttrs[attr].Metadata,
	}
}

func CreateAttributeBagForBlock(user response.User,
	attrIDs map[int]primitive.ObjectID) (response.AttrBag, error) {
	log.Info("Test case: Creating attribute bag for a block with the attribute IDs.")
	recipes.SleepBefore()
	var attrBag response.AttrBag

	attrBagReq := request.AttrBagReq{
		Name: BlockAttrsBagName,
		AttributesReq: []common.BagAttribute{
			getAttrbitueData(attrIDs[BlockID], BlockID),
			getAttrbitueData(attrIDs[Name], Name),
			getAttrbitueData(attrIDs[Description], Description),
			getAttrbitueData(attrIDs[SimpleDescription], SimpleDescription),
			getAttrbitueData(attrIDs[Tags], Tags),
			getAttrbitueData(attrIDs[Flag], Flag),
			getAttrbitueData(attrIDs[BlockType], BlockType),
			getAttrbitueData(attrIDs[ScaleAndScaleValue], ScaleAndScaleValue),
			getAttrbitueData(attrIDs[DueDate], DueDate),
			getAttrbitueData(attrIDs[StartDate], StartDate),
			getAttrbitueData(attrIDs[EndDate], EndDate),
		},
	}

	attrBag, err := attributesbag.CreateAttributeBag(user.JwtToken, attrBagReq)
	if err != nil {
		return attrBag, err
	}

	recipes.SleepAfter()
	log.Infof(".attribute bag has for a block been created, ID: %s", attrBag.ID.Hex())
	return attrBag, nil
}

func AssociateBagWithBlock(user response.User, blockID primitive.ObjectID,
	attrBag response.AttrBag) error {
	recipes.SleepBefore()
	log.Info("Test case: Associate the attribute bag with block.")

	attrBagsData := request.AttrBagsReq{
		BagIDs: []string{attrBag.ID.Hex()},
	}

	_, err := resources.AssociateAttrBagToResource(user.JwtToken, attrBagsData)
	if err != nil {
		return err
	}

	recipes.SleepAfter()
	log.Info(".attribute bag has been associated with block")
	return nil
}
