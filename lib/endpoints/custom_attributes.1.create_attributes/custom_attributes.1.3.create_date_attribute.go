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

func CreateDateAttribute(jwtToken string, reqBody request.PrimitiveAttrReq) (response.PrimitiveAttr, error) {
	var resDateAttr response.PrimitiveAttr

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resDateAttr, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteAttributesCreateDateAttribute)
	if err != nil {
		fmt.Println(err)
		return resDateAttr, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resDateAttr, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resDateAttr, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resDateAttr, err
	}

	err = json.Unmarshal(body, &resDateAttr)
	if err != nil {
		fmt.Println(err)
		return resDateAttr, err
	}

	return resDateAttr, nil
}
