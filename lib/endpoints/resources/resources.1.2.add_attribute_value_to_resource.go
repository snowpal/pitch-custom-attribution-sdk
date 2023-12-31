package resources

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

func AddAttrValueToResource(jwtToken string, resourceID string, bagID string,
	reqBody request.BagAttrsReq) (response.Attributes, error) {
	var resAttributeValue response.Attributes

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAttributeValue, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteResourcesAddAttributeValuesToResource, resourceID, bagID)
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
