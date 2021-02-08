package main

import (
	"fmt"
	"os"

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
		return
	}
	if !response.Success {
		return
	}
	fmt.Println(response)
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
