package customattributes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/request"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
)

func CreateNestedSingleSelectAttribute(jwtToken string, reqBody request.NSSAttrReq) (response.NSSAttr, error) {
	var resNestedSingleSelectAttr response.NSSAttr
	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resNestedSingleSelectAttr, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteAttributesCreateNestedSingleSelectAttribute)
	if err != nil {
		fmt.Println(err)
		return resNestedSingleSelectAttr, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resNestedSingleSelectAttr, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resNestedSingleSelectAttr, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resNestedSingleSelectAttr, err
	}

	err = json.Unmarshal(body, &resNestedSingleSelectAttr)
	if err != nil {
		fmt.Println(err)
		return resNestedSingleSelectAttr, err
	}

	return resNestedSingleSelectAttr, nil
}
