package customattributes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
)

func FetchAttributeByName(jwtToken string) (any, error) {
	var resAttributesByName any
	route, err := helpers.GetRoute(lib.RouteAttributesGetAttributeByName)
	if err != nil {
		fmt.Println(err)
		return resAttributesByName, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttributesByName, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributesByName, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributesByName, err
	}

	err = json.Unmarshal(body, &resAttributesByName)
	if err != nil {
		fmt.Println(err)
		return resAttributesByName, err
	}

	return resAttributesByName, nil
}
