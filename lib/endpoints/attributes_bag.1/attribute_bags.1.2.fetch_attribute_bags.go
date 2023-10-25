package attributesbag

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"io"
	"net/http"
)

func FetchAttributeBags(jwtToken string) (any, error) {
	var resAttributeBags any
	route, err := helpers.GetRoute(lib.RouteAttributeBagsGetAttributeBags)
	if err != nil {
		fmt.Println(err)
		return resAttributeBags, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttributeBags, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributeBags, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributeBags, err
	}

	err = json.Unmarshal(body, &resAttributeBags)
	if err != nil {
		fmt.Println(err)
		return resAttributeBags, err
	}
	return resAttributeBags, nil
}
