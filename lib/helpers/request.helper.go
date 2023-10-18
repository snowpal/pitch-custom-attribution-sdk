package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"golang.org/x/exp/slices"
)

func AddUserHeaders(jwtToken string, req *http.Request) {
	req.Header.Add("User-Authorization", jwtToken)
	addHeaders(req)
}

func AddAppHeaders(req *http.Request) {
	addHeaders(req)
}

func GetRoute(route string, args ...string) (string, error) {
	var res string
	if len(strings.Split(route, "%s"))-1 != len(args) {
		return res, errors.New("substitution does not match args count")
	}
	for _, arg := range args {
		route = strings.Replace(route, "%s", arg, 1)
	}
	res = lib.GatewayHost + route
	return res, nil
}

func GetRequestPayload(obj interface{}) (*strings.Reader, error) {
	var payload *strings.Reader
	marshaled, err := json.Marshal(obj)
	if err != nil {
		return payload, err
	}
	payload = strings.NewReader(string(marshaled))
	return payload, nil
}

// Private Methods

func addHeaders(req *http.Request) {
	req.Header.Add("x-api-key", lib.XApiKey)
	req.Header.Add("x-snowpal-product-code", lib.XProductCode)
	req.Header.Add("Content-Type", "application/json")
}

func MakeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	successCodes := []int{200, 201, 202, 204}
	if err != nil || !slices.Contains(successCodes, res.StatusCode) {
		fmt.Println(err)
		return res, errors.New("API Request Failed")
	}
	return res, nil
}
