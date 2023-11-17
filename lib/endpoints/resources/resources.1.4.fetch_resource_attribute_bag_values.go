package resources

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
)

func FetchResourceAttrBagValues(jwtToken string, resourceID string,
	bagID string) (response.ResourceAttrBagValues, error) {
	var resAttributeBagValues response.ResourceAttrBagValues
	route, err := helpers.GetRoute(lib.RouteResourcesGetResourceAttributeBagValues, resourceID, bagID)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagValues, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagValues, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagValues, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagValues, err
	}

	err = json.Unmarshal(body, &resAttributeBagValues)
	if err != nil {
		fmt.Println(err)
		return resAttributeBagValues, err
	}
	return resAttributeBagValues, nil
}
