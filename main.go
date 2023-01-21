package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "sentConfig":
			//Loading environment variables
			loadEnv()

			//Sending configuration
			fmt.Println("Sending configuration...")
			resp := setPushEventSettings()
			fmt.Println(resp.Status)
			myBody, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()

			//Check if the configuration was sent correctly
			match := regexp.MustCompile(`(result\\"(\\s)?:(\\s)?true)`).Match(myBody)
			if !match {
				log.Fatalln("Failed to send configuration")
			}

		case "getConfig":
			//Loading environment variables
			loadEnv()

			fmt.Println("Getting configuration...")
			resp := getPushEventSettings()
			fmt.Println(resp.Status)
			myBody, _ := ioutil.ReadAll(resp.Body)
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
			myBody, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(myBody))
			defer resp.Body.Close()
		}
	}

}
