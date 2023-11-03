package request

import "github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"

type AttrBagReq struct {
	Name          string                `json:"name"`
	AttributesReq []common.BagAttribute `json:"attributes"`
}
