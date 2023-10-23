package resources

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"io"
	"net/http"
)

func AddAttrValueToResource(jwtToken string, reqBody any) (any, error) {
	var resAttributeValue any

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAttributeValue, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteResourcesAddAttributeValuesToResource)
	if err != nil {
		fmt.Println(err)
		return resAttributeValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resAttributeValue, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributeValue, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributeValue, err
	}

	err = json.Unmarshal(body, &resAttributeValue)
	if err != nil {
		fmt.Println(err)
		return resAttributeValue, err
	}
	return resAttributeValue, nil
}
