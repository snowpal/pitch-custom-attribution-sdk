package customattributes

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
)

func UpdateAttribute(jwtToken string, reqBody request.AttributeReq) (response.Attribute, error) {
	var resAttribute response.Attribute

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAttribute, err
	}

	route, err := helpers.GetRoute(lib.RouteAttributesUpdateAttribute)
	if err != nil {
		fmt.Println(err)
		return resAttribute, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPut, route, payload)
	if err != nil {
		fmt.Println(err)
		return resAttribute, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttribute, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttribute, err
	}

	err = json.Unmarshal(body, &resAttribute)
	if err != nil {
		fmt.Println(err)
		return resAttribute, err
	}

	return resAttribute, nil
}
