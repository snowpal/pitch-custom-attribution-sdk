package blockattrs

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers/recipes"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
)

func CreateBlockWithCustomAttrs() {
	log.Info("Objective: Add a new team, add couple of members to that and report status for each 1 of them")
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

	var blockID primitive.ObjectID
	// Associate the above attribute bag with a block
	err = AssiciateBagWithBlock(user, blockID, attributeBag)
	if err != nil {
		return
	}

	// Add values to the block attributes
	err = AddValuesToBlockAttributes(user, blockID)
	if err != nil {
		return
	}

	// Fetch all block custom attributes
	err = FetchBlockAttributes(user)
	if err != nil {
		return
	}
}
