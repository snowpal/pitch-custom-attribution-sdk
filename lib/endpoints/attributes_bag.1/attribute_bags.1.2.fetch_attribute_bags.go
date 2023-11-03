package attributesbag

import (
	"encoding/json"
	"fmt"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
	"io"
	"net/http"
)

func FetchAttributeBags(jwtToken string) (response.AttrBags, error) {
	var resAttributeBags response.AttrBags
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
