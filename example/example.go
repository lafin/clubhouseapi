package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lafin/clubhouseapi"
)

func login(phoneNumber string) {
	response, err := clubhouseapi.StartPhoneNumberAuth(phoneNumber)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !response.Success {
		return
	}
}

func auth(phoneNumber, verificationCode string) {
	response, err := clubhouseapi.CompletePhoneNumberAuth(phoneNumber, verificationCode)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !response.Success {
		return
	}
	fmt.Println("AccessToken", response.AccessToken)
	fmt.Println("RefreshToken", response.RefreshToken)
	fmt.Println("UserID", response.UserProfile.UserID)
	env, _ := godotenv.Read()
	env["ACCESS_TOKEN"] = response.AccessToken
	env["REFRESH_TOKEN"] = response.RefreshToken
	env["USER_ID"] = strconv.Itoa(response.UserProfile.UserID)
	_ = godotenv.Write(env, ".env")
}

func channels(userID, accessToken string) {
	var credentials = map[string]string{
		"CH-UserID":     userID,
		"Authorization": fmt.Sprintf(`Bearer %s`, accessToken),
	}
	clubhouseapi.AddCredentials(credentials)
	response, err := clubhouseapi.GetChannels()
	if err != nil {
		fmt.Println(err.Error())
		refresh()
		return
	}
	if !response.Success {
		return
	}
	for _, channel := range response.Channels {
		fmt.Println(channel.ChannelID, channel.Channel, channel.Topic, channel.Club.Name)
	}
	for _, event := range response.Events {
		fmt.Println(event.EventID, event.Name, event.Description, event.Club.Name)
	}
}

func refresh() {
	refreshToken := os.Getenv("REFRESH_TOKEN")
	response, err := clubhouseapi.RefreshToken(refreshToken)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("AccessToken", response.Access)
	fmt.Println("RefreshToken", response.Refresh)
	env, _ := godotenv.Read()
	env["ACCESS_TOKEN"] = response.Access
	env["REFRESH_TOKEN"] = response.Refresh
	_ = godotenv.Write(env, ".env")
}

func main() {
	_ = godotenv.Load()
	// phoneNumber := os.Getenv("PHONE_NUMBER")
	// login(phoneNumber)
	// verificationCode := "1234"
	// auth(phoneNumber, verificationCode)

	userID := os.Getenv("USER_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")
	channels(userID, accessToken)
}
