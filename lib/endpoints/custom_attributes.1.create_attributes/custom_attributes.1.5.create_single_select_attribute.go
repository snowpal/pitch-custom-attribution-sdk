package customattributes

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/request"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"io"
	"net/http"

	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers"
)

func CreateSingleSelectAttribute(jwtToken string, reqBody request.SelectAttrReq) (response.SingleSelectAttr, error) {
	var resSingleSelectAttr response.SingleSelectAttr

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resSingleSelectAttr, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteAttributesCreateSingleSelectAttribute)
	if err != nil {
		fmt.Println(err)
		return resSingleSelectAttr, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resSingleSelectAttr, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resSingleSelectAttr, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resSingleSelectAttr, err
	}

	err = json.Unmarshal(body, &resSingleSelectAttr)
	if err != nil {
		fmt.Println(err)
		return resSingleSelectAttr, err
	}

	return resSingleSelectAttr, nil
}
