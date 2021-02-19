// Package clubhouseapi handle work with clubhouse
package clubhouseapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lafin/http"
)

var url = "https://www.clubhouseapi.com/api"
var headers = map[string]string{
	"Connection":      "keep-alive",
	"Content-Type":    "application/json",
	"Host":            "www.clubhouseapi.com",
	"Accept":          "application/json",
	"Accept-Language": "en-US;q=1",
	"User-Agent":      "clubhouse/304 (iPhone; iOS 14.4; Scale/2.00)",
	"CH-Languages":    "en-US",
	"CH-Locale":       "en_US",
	"CH-AppVersion":   "0.1.28",
	"CH-AppBuild":     "304",
	"CH-DeviceId":     "(null)",
	"CH-UserID":       "(null)",
}

// StartPhoneNumberAuth is a method of logging in by phone
func StartPhoneNumberAuth(phoneNumber string) (StartPhoneNumberAuthResponse, error) {
	var data StartPhoneNumberAuthResponse
	response, err := http.Post(fmt.Sprintf("%s/start_phone_number_auth", url), strings.NewReader(fmt.Sprintf(`{"phone_number":"%s"}`, phoneNumber)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// CompletePhoneNumberAuth is a method of completing phone authentication
func CompletePhoneNumberAuth(phoneNumber, verificationCode string) (CompletePhoneNumberAuthResponse, error) {
	var data CompletePhoneNumberAuthResponse
	response, err := http.Post(fmt.Sprintf("%s/complete_phone_number_auth", url), strings.NewReader(fmt.Sprintf(`{"verification_code":"%s","phone_number":"%s"}`, verificationCode, phoneNumber)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// GetChannels is a method of getting all channels
func GetChannels() (GetChannelsResponse, error) {
	var data GetChannelsResponse
	response, err := http.Get(fmt.Sprintf("%s/get_channels", url), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// JoinChannel is a method to join a channel
func JoinChannel(channel string) (JoinChannelResponse, error) {
	var data JoinChannelResponse
	response, err := http.Post(fmt.Sprintf("%s/join_channel", url), strings.NewReader(fmt.Sprintf(`{"channel":"%s"}`, channel)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// LeaveChannel is a method to leave a channel
func LeaveChannel(channel string) (LeaveChannelResponse, error) {
	var data LeaveChannelResponse
	response, err := http.Post(fmt.Sprintf("%s/leave_channel", url), strings.NewReader(fmt.Sprintf(`{"channel":"%s"}`, channel)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// ActivePing is a method to ping
func ActivePing(channel string) (ActivePingResponse, error) {
	var data ActivePingResponse
	response, err := http.Post(fmt.Sprintf("%s/active_ping", url), strings.NewReader(fmt.Sprintf(`{"channel":"%s"}`, channel)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// RefreshToken is a method for getting a new access token
func RefreshToken(refreshToken string) (RefreshTokenResponse, error) {
	var data RefreshTokenResponse
	response, err := http.Post(fmt.Sprintf("%s/refresh_token", url), strings.NewReader(fmt.Sprintf(`{"refresh":"%s"}`, refreshToken)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// Follow is a user follow method
func Follow(userID int) (FollowResponse, error) {
	var data FollowResponse
	response, err := http.Post(fmt.Sprintf("%s/follow", url), strings.NewReader(fmt.Sprintf(`{"source":9,"source_topic_id":null,"user_id":%d,"user_ids":null}`, userID)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// Unfollow is a user unfollow method
func Unfollow(userID int) (UnfollowResponse, error) {
	var data UnfollowResponse
	response, err := http.Post(fmt.Sprintf("%s/unfollow", url), strings.NewReader(fmt.Sprintf(`{"user_id":%d}`, userID)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// GetFollowing is a get following method
func GetFollowing(userID, pageSize, page int) (GetFollowingResponse, error) {
	if pageSize == 0 {
		pageSize = 400
	}
	if page == 0 {
		page = 1
	}
	var data GetFollowingResponse
	response, err := http.Post(fmt.Sprintf("%s/get_following?page_size=%d&page=%d", url, pageSize, page), strings.NewReader(fmt.Sprintf(`{"user_id":%d}`, userID)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// GetFollowers is a get followers method
func GetFollowers(userID, pageSize, page int) (GetFollowersResponse, error) {
	if pageSize == 0 {
		pageSize = 400
	}
	if page == 0 {
		page = 1
	}
	var data GetFollowersResponse
	response, err := http.Post(fmt.Sprintf("%s/get_followers?page_size=%d&page=%d", url, pageSize, page), strings.NewReader(fmt.Sprintf(`{"user_id":%d}`, userID)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// GetProfileResponse is a get profile method
func GetProfile(userID int) (GetProfileResponse, error) {
	var data GetProfileResponse
	response, err := http.Post(fmt.Sprintf("%s/get_profile", url), strings.NewReader(fmt.Sprintf(`{"user_id":%d}`, userID)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// AudienceReplyResponse is a raise hand method
func AudienceReply(channel string, raiseHand bool) (AudienceReplyResponse, error) {
	var data AudienceReplyResponse
	response, err := http.Post(fmt.Sprintf("%s/audience_reply", url), strings.NewReader(fmt.Sprintf(`{"channel":"%s","raise_hands":%t,"unraise_hands":%t}`, channel, raiseHand, !raiseHand)), headers)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(response, &data); err != nil {
		return data, err
	}
	return data, nil
}

// AddCredentials is a method of adding credentials such as User ID or Access Token
func AddCredentials(credentials map[string]string) {
	for k, v := range credentials {
		headers[k] = v
	}
}
