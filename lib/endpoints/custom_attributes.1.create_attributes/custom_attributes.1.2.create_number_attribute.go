package customattributes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
)

func CreateNumberAttribute(jwtToken string, reqBody any) (any, error) {
	var resNumberAttr any

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resNumberAttr, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteAttributesCreateNumberAttribute)
	if err != nil {
		fmt.Println(err)
		return resNumberAttr, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resNumberAttr, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resNumberAttr, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resNumberAttr, err
	}

	err = json.Unmarshal(body, &resNumberAttr)
	if err != nil {
		fmt.Println(err)
		return resNumberAttr, err
	}

	return resNumberAttr, nil
}
