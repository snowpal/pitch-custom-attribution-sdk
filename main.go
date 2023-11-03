package main

import (
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/config"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/recipes"

	log "github.com/sirupsen/logrus"
	blockattrs "github.com/snowpal/pitch-custom-attribution-sdk/lib/recipes/recipe.2.create_block_with_custom_attrs"
)

func main() {
	var err error
	if config.Files, err = config.InitConfigFiles(); err != nil {
		log.Fatal(err.Error())
		return
	}

	recipeID := 1
	switch recipeID {
	case 1:
		log.Info("Run Recipe1")
		recipes.RegisterFewUsers()
	case 2:
		log.Info("Run Recipe2")
		blockattrs.CreateBlockWithCustomAttrs()
	default:
		log.Info("pick a specific recipe to run")
	}
}
