package customattributes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
)

func FetchAttributes(jwtToken string) (any, error) {
	var resAttributes any
	route, err := helpers.GetRoute(lib.RouteAttributesGetAttributes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, &resAttributes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return nil, nil
}
