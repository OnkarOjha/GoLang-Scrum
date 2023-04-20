// Download the helper library from https://www.twilio.com/docs/go/install
package main

import (
	"fmt"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

const TWILIO_ACCOUNT_SID= "AC5869424c6ae2d0b27f66a5d5b9b90485"
const TWILIO_AUTH_TOKEN= "fc2aa066a47915f9019d0e8e4ea35f9b"

var TwilioClient *twilio.RestClient



func main() {
	
	TwilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
	// client := twilio.NewRestClient()

	params := &api.CreateCallParams{}
	params.SetUrl("https://4019-112-196-113-2.ngrok-free.app/answer")
	params.SetTo("+919877370350")
	params.SetFrom("+16204009493") //number brought from twilio

	resp, err := TwilioClient.Api.CreateCall(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}