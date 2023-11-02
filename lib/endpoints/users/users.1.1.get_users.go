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

func GetUsers(jwtToken string) ([]response.User, error) {
	resUsers := response.Users{}
	route, err := helpers.GetRoute(lib.RouteUserGetUsers)
	if err != nil {
		fmt.Println(err)
		return resUsers.Users, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUsers.Users, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUsers.Users, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUsers.Users, err
	}

	err = json.Unmarshal(body, &resUsers)
	if err != nil {
		fmt.Println(err)
		return resUsers.Users, err
	}
	return resUsers.Users, nil
}
