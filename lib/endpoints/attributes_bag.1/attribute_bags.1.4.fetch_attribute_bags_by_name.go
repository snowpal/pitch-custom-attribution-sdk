package attributesbag

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
	"io"
	"net/http"
)

func FetchAttributeBagsByName(jwtToken string) (response.AttrBag, error) {
	var resAttributeBagsByName response.AttrBag
	route, err := helpers.GetRoute(lib.RouteAttributeBagsGetAttributeBagByName)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagsByName, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagsByName, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagsByName, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagsByName, err
	}

	err = json.Unmarshal(body, &resAttributeBagsByName)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagsByName, err
	}

	return resAttributeBagsByName, nil
}
