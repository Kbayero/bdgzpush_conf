package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "sendConfig":
			//Loading environment variables
			loadEnv()

			//Sending configuration
			fmt.Println("Sending configuration...")
			resp := setPushEventSettings()
			fmt.Println(resp.Status)
			myBody, _ := io.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()

			//Check if the configuration was sent correctly
			regex := regexp.MustCompile(`result":true`)
			match := regex.Match([]byte(string(myBody)))
			if !match {
				log.Fatalln("Failed to send configuration")
			}

		case "getConfig":
			//Loading environment variables
			loadEnv()

			fmt.Println("Getting configuration...")
			resp := getPushEventSettings()
			fmt.Println(resp.Status)
			myBody, _ := io.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()

			//Check if the configuration was enabled
			match := regexp.MustCompile(`(status\\"(\\s)?:(\\s)?1)`).Match(myBody)
			if !match {
				log.Fatalln("Failed, configuration disabled")
			}

		case "logTest":
			//Loading environment variables
			loadEnv()

			fmt.Println("Send Event Test...")
			resp := sendTestPushEvent()
			fmt.Println(resp.Status)
			myBody, _ := io.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()
		}
	}

}
