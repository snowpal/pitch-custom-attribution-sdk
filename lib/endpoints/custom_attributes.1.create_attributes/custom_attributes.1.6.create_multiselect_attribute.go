package customattributes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
)

func CreateMultiselectAttribute(jwtToken string, reqBody any) (any, error) {
	var resMultiselectAttr any

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resMultiselectAttr, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteAttributesCreateMultiselectAttribute)
	if err != nil {
		fmt.Println(err)
		return resMultiselectAttr, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resMultiselectAttr, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resMultiselectAttr, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resMultiselectAttr, err
	}

	err = json.Unmarshal(body, &resMultiselectAttr)
	if err != nil {
		fmt.Println(err)
		return resMultiselectAttr, err
	}
	return resMultiselectAttr, nil
}
