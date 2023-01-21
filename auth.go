package main

import "encoding/base64"

func generateAuthCode(apiKey string) string {
	loginString := apiKey + ":"
	encodedBytes := base64.StdEncoding.EncodeToString([]byte(loginString))
	encodedUserPassSequence := string(encodedBytes[:])
	authCode := "Basic " + encodedUserPassSequence

	return authCode
}
