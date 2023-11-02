package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
)

func GetUserByUuid(jwtToken string, userUuid string) (response.User, error) {
	resUser := response.User{}
	route, err := helpers.GetRoute(lib.RouteUserGetUserByUuid, userUuid)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}

	err = json.Unmarshal(body, &resUser)
	if err != nil {
		fmt.Println(err)
		return resUser, err
	}
	return resUser, nil
}
