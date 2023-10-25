package attributesbag

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"io"
	"net/http"
)

func CreateAttributeBag(jwtToken string, reqBody any) (any, error) {
	var resAttributeBag any

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAttributeBag, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteAttributeBagsCreateAttributeBag)
	if err != nil {
		fmt.Println(err)
		return resAttributeBag, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resAttributeBag, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributeBag, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributeBag, err
	}

	err = json.Unmarshal(body, &resAttributeBag)
	if err != nil {
		fmt.Println(err)
		return resAttributeBag, err
	}
	return resAttributeBag, nil
}
