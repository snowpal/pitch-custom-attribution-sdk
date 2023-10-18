package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type Profile struct {
	ID               string                `json:"ID"`
	UserID           string                `json:"userId"`
	Name             common.ProfileName    `json:"Name"`
	PhoneNumbers     []common.PhoneNumber  `json:"PhoneNumbers"`
	EmailAddresses   []common.EmailAddress `json:"EmailAddresses"`
	UserEmailAddress string                `json:"UserEmailAddress"`
	LastModifiedOn   string                `json:"lastModifiedOn"`
}
