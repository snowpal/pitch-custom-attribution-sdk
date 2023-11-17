package blockattrs

import (
	"fmt"

	"github.com/snowpal/pitch-custom-attribution-sdk/lib/endpoints/resources"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/request"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func getTextAttributeValue(attrID primitive.ObjectID, value string) *request.BagTextAttrReq {
	return &request.BagTextAttrReq{
		ID:   attrID.Hex(),
		Text: common.BagPrimitiveAttr{Value: &value},
	}
}

func getDateAttributeValue(attrID primitive.ObjectID, value string) *request.BagDateAttrReq {
	return &request.BagDateAttrReq{
		ID:   attrID.Hex(),
		Date: common.BagPrimitiveAttr{Value: &value},
	}
}

func getSingleSelectAttributeValue(attrID primitive.ObjectID, selectedID int) *request.BagSelectAttrReq {
	return &request.BagSelectAttrReq{
		ID:     attrID.Hex(),
		Option: common.BagSelectAttr{SelectedID: &selectedID},
	}
}

func getNestedSingleSelectAttributeValue(attrID primitive.ObjectID, selectedID int,
	childSelectedID int) *request.BagNSSAttrReq {
	return &request.BagNSSAttrReq{
		Option: common.BagNSSAttr{
			SelectedID: selectedID,
			Child: &common.NSSValueChild{
				Option: &common.BagSelectOption{
					BagSelectAttr: &common.BagSelectAttr{SelectedID: &childSelectedID},
				},
			},
		},
	}
}

func AddValuesToBlockAttributes(user response.User, blockID primitive.ObjectID,
	attributeBag response.AttrBag, attrIDs map[int]primitive.ObjectID) error {
	log.Info("Test Case: Add values to block attributes.")
	recipes.SleepBefore()

	attrsReq := request.BagAttrsReq{
		Attributes: []request.BagAttrReq{
			{BagTextAttrReq: getTextAttributeValue(attrIDs[BlockID], BlockIDValue)},
			{BagTextAttrReq: getTextAttributeValue(attrIDs[Name], BlockNameValue)},
			{BagTextAttrReq: getTextAttributeValue(attrIDs[Description], BlockDescValue)},
			{BagTextAttrReq: getTextAttributeValue(attrIDs[SimpleDescription], BlockSimpleDescValue)},
			{BagTextAttrReq: getTextAttributeValue(attrIDs[Tags], BlockTagsValue)},
			{BagTextAttrReq: getTextAttributeValue(attrIDs[Flag], BlockFlagValue)},
			{BagSelectAttrReq: getSingleSelectAttributeValue(attrIDs[BlockType], BlockTypeID)},
			{BagNSSAttrReq: getNestedSingleSelectAttributeValue(attrIDs[ScaleAndScaleValue],
				BlockScaleID, BlockScaleValueID)},
			{BagDateAttrReq: getDateAttributeValue(attrIDs[DueDate], BlockDueDateValue)},
			{BagDateAttrReq: getDateAttributeValue(attrIDs[StartDate], BlockStartDateValue)},
			{BagDateAttrReq: getDateAttributeValue(attrIDs[EndDate], BlockEndDateValue)},
		},
	}

	if _, err := resources.AddAttrValueToResource(user.JwtToken, blockID.Hex(),
		attributeBag.ID.Hex(), attrsReq); err != nil {
		return err
	}

	recipes.SleepAfter()
	log.Info(".values has been added to block attributes.")
	return nil
}

func FetchBlockAttributes(user response.User, blockID primitive.ObjectID,
	bagID primitive.ObjectID) (response.ResourceAttrBagValues, error) {
	log.Info("Test Case: Fetch all block custom attributes along with their values.")
	recipes.SleepBefore()

	attrBagValues, err := resources.FetchResourceAttrBagValues(user.JwtToken, blockID.Hex(),
		bagID.Hex())
	if err != nil {
		return attrBagValues, err
	}

	recipes.SleepAfter()
	log.Info(".values has been added to block attributes.")
	return attrBagValues, nil
}

func DisplayBlockAttributes(user response.User, attrBagValues response.ResourceAttrBagValues) {
	blockAttrs := attrBagValues.Attributes
	fmt.Printf("\nBlockID: %s\n", *blockAttrs[BlockID].Text.Value)
	fmt.Printf("\nName: %s\n", *blockAttrs[Name].Text.Value)
	fmt.Printf("\nDescription: %s\n", *blockAttrs[Description].Text.Value)
	fmt.Printf("\nSimple Description: %s\n", *blockAttrs[SimpleDescription].Text.Value)
	fmt.Printf("\nTags: %s\n", *blockAttrs[Tags].Text.Value)
	fmt.Printf("\nFlag: %s\n", *blockAttrs[Flag].Text.Value)
	fmt.Printf("\nBlock Type: %d\n", *blockAttrs[BlockType].Option.BagSelectAttr.SelectedID)
	fmt.Printf("\nScale: %d\n", blockAttrs[ScaleAndScaleValue].Option.BagNSSAttr.SelectedID)
	fmt.Printf("\nScale Value: %d\n",
		blockAttrs[ScaleAndScaleValue].Option.BagNSSAttr.Child.Option.BagSelectAttr.SelectedID)
	fmt.Printf("\nDue Date: %s\n", *blockAttrs[DueDate].Date.Value)
	fmt.Printf("\nStart Date: %s\n", *blockAttrs[StartDate].Date.Value)
	fmt.Printf("\nEnd Date: %s\n", *blockAttrs[EndDate].Date.Value)
}
