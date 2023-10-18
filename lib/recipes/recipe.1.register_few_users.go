package recipes

import (
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"

	log "github.com/sirupsen/logrus"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	user, err := recipes.RegisterUser(lib.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.AddNewProfile(user, "krish12345", "krish", "123-234-5566")
	if err != nil {
		return
	}

	user, err = recipes.RegisterUser(lib.ActiveUser)
	if err != nil {
		return
	}

	_, err = recipes.AddNewProfile(user, "varun12345", "varun", "234-443-4456")
	if err != nil {
		return
	}

	user, err = recipes.RegisterUser(lib.SecondUser)
	if err != nil {
		return
	}

	_, err = recipes.AddNewProfile(user, "raj1234567", "raj", "234-443-4456")
	if err != nil {
		return
	}

	user, err = recipes.RegisterUser(lib.ThirdUser)
	if err != nil {
		return
	}

	_, err = recipes.AddNewProfile(user, "ben1234567", "ben", "234-443-4456")
	if err != nil {
		return
	}

	user, err = recipes.RegisterUser(lib.FourthUser)
	if err != nil {
		return
	}

	_, err = recipes.AddNewProfile(user, "scott12345", "scott", "234-443-4456")
	if err != nil {
		return
	}
}
