package customattributes

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
)

func FetchAttributes(jwtToken string) (response.Attributes, error) {
	var resAttributes response.Attributes
	route, err := helpers.GetRoute(lib.RouteAttributesGetAttributes)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	err = json.Unmarshal(body, &resAttributes)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	return resAttributes, nil
}
