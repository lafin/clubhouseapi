package clubhouseapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lafin/http"
)

var URL = "https://www.clubhouseapi.com/api"
var headers = map[string]string{
	"Content-Type": "application/json",
	"User-Agent":   "clubhouse/269 (iPhone; iOS 14.1; Scale/3.00)",
	"CH-AppBuild":  "269",
}

// StartPhoneNumberAuth is ...
func StartPhoneNumberAuth(phoneNumber string) (startPhoneNumberAuthResponse, error) {
	var data startPhoneNumberAuthResponse
	response, err := http.Post(fmt.Sprintf("%s/start_phone_number_auth", URL), strings.NewReader(fmt.Sprintf(`{"phone_number":"%s"}`, phoneNumber)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// CompletePhoneNumberAuth is ...
func CompletePhoneNumberAuth(phoneNumber, verificationCode string) (completePhoneNumberAuthResponse, error) {
	var data completePhoneNumberAuthResponse
	response, err := http.Post(fmt.Sprintf("%s/complete_phone_number_auth", URL), strings.NewReader(fmt.Sprintf(`{"verification_code":"%s","phone_number":"%s"}`, verificationCode, phoneNumber)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// GetChannels is ...
func GetChannels() (getChannelsResponse, error) {
	var data getChannelsResponse
	response, err := http.Post(fmt.Sprintf("%s/get_channels", URL), strings.NewReader("{}"), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

func AddCredentials(credentials map[string]string) {
	for k, v := range credentials {
		headers[k] = v
	}
}
