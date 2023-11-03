package request

import "github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/common"

type AttrBagsReq struct {
	BagIDs []string `json:"bagIDs"`
}

type BagSelectAttrReq struct {
	ID     string               `json:"ID"`
	Option common.BagSelectAttr `json:"option"`
}

type BagNSSAttrReq struct {
	ID     string            `json:"ID"`
	Option common.BagNSSAttr `json:"option"`
}

type BagTextAttrReq struct {
	ID   string                  `json:"ID"`
	Text common.BagPrimitiveAttr `json:"text"`
}

type BagNumberAttrReq struct {
	ID     string                  `json:"ID"`
	Number common.BagPrimitiveAttr `json:"number"`
}

type BagDateAttrReq struct {
	ID   string                  `json:"ID"`
	Date common.BagPrimitiveAttr `json:"date"`
}

type BagFileAttrReq struct {
	ID   string                  `json:"ID"`
	File common.BagPrimitiveAttr `json:"file"`
}
