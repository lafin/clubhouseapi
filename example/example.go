package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

func channels(userID, accessToken string) clubhouseapi.GetChannelsResponse {
	var credentials = map[string]string{
		"CH-UserID":     userID,
		"Authorization": fmt.Sprintf(`Bearer %s`, accessToken),
	}
	clubhouseapi.AddCredentials(credentials)
	response, err := clubhouseapi.GetChannels()
	if err != nil {
		fmt.Println(err.Error())
		refresh()
		return response
	}
	return response
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

func userIsAlreadyInChannel(channel clubhouseapi.Channel, userID int) bool {
	for _, user := range channel.Users {
		if user.UserID == userID {
			return true
		}
	}
	return false
}

func main() {
	_ = godotenv.Load()
	// phoneNumber := os.Getenv("PHONE_NUMBER")
	// login(phoneNumber)
	// verificationCode := "1234"
	// auth(phoneNumber, verificationCode)

	userID := os.Getenv("USER_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")
	response := channels(userID, accessToken)
	for _, channel := range response.Channels {
		fmt.Println(channel.ChannelID, channel.Channel, channel.Topic, channel.Club.Name, channel.NumAll, channel.NumSpeakers)
		parsedUserID, _ := strconv.ParseInt(userID, 10, 32)
		if !userIsAlreadyInChannel(channel, int(parsedUserID)) {
			fmt.Println("join to channel", channel.Channel)
			_, err := clubhouseapi.JoinChannel(channel.Channel)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			fmt.Println("ping channel", channel.Channel)
			_, err := clubhouseapi.ActivePing(channel.Channel)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		time.Sleep(2 * time.Second)
	}
}
