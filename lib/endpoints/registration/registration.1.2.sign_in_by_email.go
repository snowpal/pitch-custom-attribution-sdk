package registration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
)

func SignInByEmail(reqBody request.SignInReqBody) (response.User, error) {
	resUser := response.UserRegistration{}

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resUser.User, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteRegistrationSignInByEmail)
	if err != nil {
		fmt.Println(err)
		return resUser.User, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resUser.User, err
	}

	helpers.AddAppHeaders(req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUser.User, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUser.User, err
	}

	err = json.Unmarshal(body, &resUser)
	if err != nil {
		fmt.Println(err)
		return resUser.User, err
	}
	return resUser.User, nil
}
