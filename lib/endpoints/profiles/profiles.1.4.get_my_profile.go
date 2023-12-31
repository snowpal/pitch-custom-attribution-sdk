package profiles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-custom-attribution-sdk/lib"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/helpers"
	"github.com/snowpal/pitch-custom-attribution-sdk/lib/structs/response"
)

func GetMyProfile(jwtToken string) (response.Profile, error) {
	var resProfile response.Profile

	route, err := helpers.GetRoute(lib.RouteProfilesGetMyProfile)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}

	err = json.Unmarshal(body, &resProfile)
	if err != nil {
		fmt.Println(err)
		return resProfile, err
	}
	return resProfile, nil
}
