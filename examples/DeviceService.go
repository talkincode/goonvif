package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/talkincode/goonvif"
	"github.com/talkincode/goonvif/Device"
)

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main1() {
	// Getting an camera instance
	dev, err := goonvif.NewDevice("192.168.0.100:80")
	if err != nil {
		panic(err)
	}
	// Authorization
	dev.Authenticate("dev", "dev")

	// Preparing commands
	// systemDateAndTyme := Device.GetSystemDateAndTime{}
	// getCapabilities := Device.GetCapabilities{Category:"All"}
	getDns := Device.GetDNS{}
	// createUser := Device.CreateUsers{User:
	// 		onvif.User{
	// 			Username:  "TestUser",
	// 			Password:  "TestPassword",
	// 			UserLevel: "User",
	// 		},
	// 	}

	// Commands execution
	// systemDateAndTymeResponse, err := dev.CallMethod(systemDateAndTyme)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(readResponse(systemDateAndTymeResponse))
	// }
	// getCapabilitiesResponse, err := dev.CallMethod(getCapabilities)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	r := readResponse(getCapabilitiesResponse)
	// 	fmt.Println(r)
	// }
	getDnsResp, err := dev.CallMethod(getDns)
	if err != nil {
		log.Println(err)
	} else {
		r := readResponse(getDnsResp)
		fmt.Println(r)
	}
	// createUserResponse, err := dev.CallMethod(createUser)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	/*
	// 	You could use https://github.com/talkincode/gosoap for pretty printing response
	// 	 */
	// 	fmt.Println(gosoap.SoapMessage(readResponse(createUserResponse)).StringIndent())
	// }

}
