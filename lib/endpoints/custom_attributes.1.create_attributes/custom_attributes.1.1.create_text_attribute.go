package customattributes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
)

func CreateTextAttribute(jwtToken string, reqBody request.PrimitiveAttrReq) (response.PrimitiveAttr, error) {
	var resTextAttr response.PrimitiveAttr

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTextAttr, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteAttributesCreateTextAttribute)
	if err != nil {
		fmt.Println(err)
		return resTextAttr, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTextAttr, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTextAttr, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTextAttr, err
	}

	err = json.Unmarshal(body, &resTextAttr)
	if err != nil {
		fmt.Println(err)
		return resTextAttr, err
	}

	return resTextAttr, nil
}
