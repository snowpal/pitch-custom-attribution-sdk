package field_types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
)

func FetchFieldTypes(jwtToken string) ([]response.FieldType, error) {
	resFieldTypes := response.FieldTypes{}
	route, err := helpers.GetRoute(lib.RouteFieldTypesGetFieldTypes)
	if err != nil {
		fmt.Println(err)
		return resFieldTypes.FieldTypes, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resFieldTypes.FieldTypes, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFieldTypes.FieldTypes, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resFieldTypes.FieldTypes, err
	}

	err = json.Unmarshal(body, &resFieldTypes)
	if err != nil {
		fmt.Println(err)
		return resFieldTypes.FieldTypes, err
	}
	return resFieldTypes.FieldTypes, nil
}
