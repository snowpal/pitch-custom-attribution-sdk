/*
Check out the recipes to get a sense of what the underlying API has to offer. Here's a simple example:

	package main

	import (
		"github.com/snowpal/pitch-go-status-sdk/lib/recipes"

		log "github.com/sirupsen/logrus"
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
		default:
			log.Info("pick a specific recipe to run")
		}
	}

For a full guide, visit https://developers.snowpal.com
*/
package main
