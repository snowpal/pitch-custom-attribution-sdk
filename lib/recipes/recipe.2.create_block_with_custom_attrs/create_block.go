package blockattrs

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func CreateBlockWithCustomAttrs() {
	log.Info("Objective: Create a block with all its custom attributes")
	if _, err := recipes.ValidateDependencies(); err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	// Create all custom attributes you want to add to the block
	attributeIDs, err := CreateBlockAttributes(user)
	if err != nil {
		return
	}

	// Create an attribute bag with the block attributes above
	attributeBag, err := CreateAttributeBagForBlock(user, attributeIDs)
	if err != nil {
		return
	}

	// Associate the above attribute bag with a block
	var blockID primitive.ObjectID
	if blockID, err = primitive.ObjectIDFromHex("26554ee536e4f585fb899f26"); err != nil {
		return
	}
	if err = AssociateBagWithBlock(user, blockID, attributeBag); err != nil {
		return
	}

	// Add values to the block attributes
	if err = AddValuesToBlockAttributes(user, blockID, attributeBag, attributeIDs); err != nil {
		return
	}

	// Fetch all block custom attributes
	var attrBagValues response.ResourceAttrBagValues
	if attrBagValues, err = FetchBlockAttributes(user, blockID, attributeBag.ID); err != nil {
		return
	}

	DisplayBlockAttributes(user, attrBagValues)
}
