package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/lafin/clubhouseapi"
)

func login() {
	phoneNumber := os.Getenv("PHONE_NUMBER")
	response, err := clubhouseapi.StartPhoneNumberAuth(phoneNumber)
	if err != nil {
		fmt.Println("login", err.Error())
		return
	}
	if !response.Success {
		return
	}
}

func auth(verificationCode string) {
	phoneNumber := os.Getenv("PHONE_NUMBER")
	response, err := clubhouseapi.CompletePhoneNumberAuth(phoneNumber, verificationCode)
	if err != nil {
		fmt.Println("auth", err.Error())
		return
	}
	if !response.Success {
		return
	}
	env, _ := godotenv.Read()
	env["ACCESS_TOKEN"] = response.AccessToken
	env["REFRESH_TOKEN"] = response.RefreshToken
	env["USER_ID"] = strconv.Itoa(response.UserProfile.UserID)
	_ = godotenv.Write(env, ".env")
}

func channels() (clubhouseapi.GetChannelsResponse, error) {
	userID := os.Getenv("USER_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")
	var credentials = map[string]string{
		"CH-UserID":     userID,
		"Authorization": fmt.Sprintf(`Bearer %s`, accessToken),
	}
	clubhouseapi.AddCredentials(credentials)
	return clubhouseapi.GetChannels()
}

func refresh() {
	refreshToken := os.Getenv("REFRESH_TOKEN")
	response, err := clubhouseapi.RefreshToken(refreshToken)
	if err != nil {
		fmt.Println("refresh", err.Error())
		return
	}
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

func autoRefresh() {
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			refresh()
			fmt.Print("@")
			_ = godotenv.Overload()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func main() {
	_ = godotenv.Load()
	refresh()
	go autoRefresh()

	// login()
	// auth("1234")

	for {
		response, err := channels()
		if err != nil {
			fmt.Println("channels", err.Error())
			time.Sleep(10 * time.Second)
			continue
		}
		for _, channel := range response.Channels {
			fmt.Println(channel.ChannelID, channel.Channel, channel.Topic, channel.Club.Name, channel.NumAll, channel.NumSpeakers)
		}
		for _, channel := range response.Channels {
			userID, _ := strconv.ParseInt(os.Getenv("USER_ID"), 10, 32)
			if !userIsAlreadyInChannel(channel, int(userID)) {
				_, err := clubhouseapi.JoinChannel(channel.Channel)
				if err != nil {
					fmt.Println("join", err.Error())
					break
				}
				fmt.Print("+")
			} else {
				_, err := clubhouseapi.ActivePing(channel.Channel)
				if err != nil {
					fmt.Println("ping", err.Error())
					break
				}
				fmt.Print(".")
			}
			time.Sleep(2 * time.Second)
		}
		time.Sleep(10 * time.Second)
	}
}
