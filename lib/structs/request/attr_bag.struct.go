package request

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type AttrBagReq struct {
	Name          string                `json:"name"`
	AttributesReq []common.BagAttribute `json:"attributes"`
}
