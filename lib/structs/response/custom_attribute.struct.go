package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type TextAttr struct {
	ID    string `json:"ID"`
	Name  string `json:"name"`
	Label bool   `json:"label"`
	common.UserMetadata
}
