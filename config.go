package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ServiceSettings struct {
	Url                        string `json:"url"`
	Authorization              string `json:"authorization"`
	RequireValidSslCertificate bool   `json:"requireValidSslCertificate"`
}

type Params struct {
	Status                int             `json:"status"`
	ServiceType           string          `json:"serviceType"`
	Servicesettings       ServiceSettings `json:"serviceSettings"`
	SubscribeToEventTypes json.RawMessage `json:"subscribeToEventTypes"`
}

type ParamsTest struct {
	EventType string `json:"eventType"`
}

type TemplateConfigSetPush struct {
	PARAMS  Params `json:"params"`
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      string `json:"id"`
}

type TemplateConfigGetPush struct {
	PARAMS  json.RawMessage `json:"params"`
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	ID      string          `json:"id"`
}

type TemplateTestPush struct {
	PARAMS  ParamsTest `json:"params"`
	JSONRPC string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	ID      string     `json:"id"`
}

func TemplateSetPushEventSettings(auth string) []byte {

	byteTemplate := TemplateConfigSetPush{
		PARAMS: Params{
			Status:      1,
			ServiceType: "cef",
			Servicesettings: ServiceSettings{
				Url:                        getenv("BDGZ_URL"),
				Authorization:              auth,
				RequireValidSslCertificate: false,
			},
			SubscribeToEventTypes: []byte(`{
				"adcloud" : true ,
				"antiexploit" : true ,
				"aph" : true ,
				"av" : true ,
				"avc" : true ,
				"dp" : true ,
				"endpoint-moved-in" : true ,
				"endpoint-moved-out" : true ,
				"exchange-malware" : true ,
				"exchange-user-credentials" : true ,
				"fw" : true ,
				"hwid-change" : true ,
				"install" : false ,
				"modules" : false ,
				"network-monitor" : true ,
				"network-sandboxing" : true ,
				"new-incident" : true ,
				"ransomware-mitigation" : true ,
				"registration" : false ,
				"security-container-update-available" : true ,
				"supa-update-status" : true ,
				"sva" : true ,
				"sva-load" : true ,
				"task-status" : true ,
				"troubleshooting-activity" : true ,
				"uc" : true ,
				"uninstall" : false
			}`),
		},
		JSONRPC: "2.0",
		Method:  "setPushEventSettings",
		ID:      "1",
	}

	body, err := json.Marshal(byteTemplate)
	if err != nil {
		panic(err)
	}
	return body

}

func TemplateGetPushEventSettings() []byte {
	byteTemplate := TemplateConfigGetPush{
		PARAMS:  []byte(`{}`),
		JSONRPC: "2.0",
		Method:  "getPushEventSettings",
		ID:      "3",
	}
	body, err := json.Marshal(byteTemplate)
	if err != nil {
		panic(err)
	}
	return body
}

func TemplateSendTestPushEvent() []byte {
	byteTemplate := TemplateTestPush{
		PARAMS: ParamsTest{
			EventType: "av",
		},
		JSONRPC: "2.0",
		Method:  "sendTestPushEvent",
		ID:      "4",
	}
	body, err := json.Marshal(byteTemplate)
	if err != nil {
		panic(err)
	}
	return body
}

func sendRequest(body []byte, authorizationHeader string) *http.Response {
	postUrl := getenv("BDGZ_ACCESS_URL")
	// Create a HTTP post request
	r, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	//Adding the headers
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", authorizationHeader)

	//Create a client
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	return resp

}

// function to obtain the Bitdefender API configuration
func getPushEventSettings() *http.Response {
	//Defining the authentication code
	authorizationHeader := generateAuthCode(getenv("BDGZ_API_KEY"))

	//Generating request body
	body := TemplateGetPushEventSettings()

	//Sending configuration
	resp := sendRequest(body, authorizationHeader)
	return resp

}

func setPushEventSettings() *http.Response {
	//Defining the authentication code
	authorizationHeader := generateAuthCode(getenv("BDGZ_API_KEY"))

	//Generating request body
	body := TemplateSetPushEventSettings(authorizationHeader)

	//Sending configuration
	return sendRequest(body, authorizationHeader)
}

func sendTestPushEvent() *http.Response {
	//Defining the authentication code
	authorizationHeader := generateAuthCode(getenv("BDGZ_API_KEY"))

	//Generating request body
	body := TemplateSendTestPushEvent()

	//Sending configuration
	return sendRequest(body, authorizationHeader)
}
