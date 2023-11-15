package recipes

import (
	"time"

	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/request"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	log "github.com/sirupsen/logrus"
	ca "github.com/snowpal/pitch-custom-attribution-sdk/lib/endpoints/custom_attributes.1.create_attributes"
)

func sleepWindow(sleepTime time.Duration) {
	time.Sleep(time.Second * sleepTime)
}

func SleepBefore() {
	sleepWindow(1)
}

func SleepAfter() {
	sleepWindow(2)
}

// ValidateDependencies We require that the first recipe be run before anything else as it registers a bunch of users.
// To verify if it was actually run, we do this "random" check.
func ValidateDependencies() (response.User, error) {
	user, err := SignIn(lib.DefaultEmail, lib.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateTextAttribute(user response.User, attrName string, label bool) (primitive.ObjectID,
	error) {
	SleepBefore()
	log.Infof("Creating %s text attribute.", attrName)

	var attribute response.PrimitiveAttr
	attribute, err := ca.CreateTextAttribute(
		user.JwtToken,
		request.PrimitiveAttrReq{
			Name:  attrName,
			Label: &label,
		},
	)
	if err != nil {
		return attribute.ID, err
	}

	SleepAfter()
	log.Infof(".%s attribute has been created, ID: %s", attrName, attribute.ID.Hex())
	return attribute.ID, nil
}

func CreateSingleSelectAttribute(user response.User, singleSelectName string,
	options []common.SelectOption) (primitive.ObjectID, error) {
	SleepBefore()
	log.Infof("Creating %s single select.", singleSelectName)

	var attribute response.SingleSelectAttr
	attribute, err := ca.CreateSingleSelectAttribute(
		user.JwtToken,
		request.SelectAttrReq{
			Name:    singleSelectName,
			Options: options,
		},
	)
	if err != nil {
		return attribute.ID, err
	}

	SleepAfter()
	log.Infof(".%s attribute has been created, ID: %s", singleSelectName, attribute.ID.Hex())
	return attribute.ID, nil
}

func CreateDateAttribute(user response.User, attrName string) (primitive.ObjectID,
	error) {
	SleepBefore()
	log.Infof("Creating %s date attribute.", attrName)

	var attribute response.PrimitiveAttr
	attribute, err := ca.CreateDateAttribute(
		user.JwtToken,
		request.PrimitiveAttrReq{Name: attrName},
	)
	if err != nil {
		return attribute.ID, err
	}

	SleepAfter()
	log.Infof(".%s attribute has been created, ID: %s", attrName, attribute.ID.Hex())
	return attribute.ID, nil
}

func CreateNSSAttribute(user response.User, attrName string,
	options []common.NSSOption) (primitive.ObjectID,
	error) {
	SleepBefore()
	log.Infof("Creating %s nested single select.", attrName)

	var attributeID primitive.ObjectID
	attribute, err := ca.CreateNestedSingleSelectAttribute(
		user.JwtToken,
		request.NSSAttrReq{Name: attrName},
	)
	if err != nil {
		return attribute.ID, err
	}

	SleepAfter()
	log.Infof(".%s attribute has been created, ID: %s", attrName, attribute.ID.Hex())
	return attributeID, nil
}
