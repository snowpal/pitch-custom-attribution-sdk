package request

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type ProfileReqBody struct {
	UserID         string                `json:"userID"`
	Name           common.ProfileName    `json:"Name"`
	PhoneNumbers   []common.PhoneNumber  `json:"phoneNumbers"`
	EmailAddresses []common.EmailAddress `json:"emailAddresses"`
}
