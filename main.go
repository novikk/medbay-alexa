package main

import (
	"os"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var Applications = map[string]interface{}{
	"/echo/getpills": alexa.EchoApplication{ // Route
		AppID:    "amzn1.ask.skill.254cb708-27ac-4d2b-98da-5763aff261b6", // Echo App ID from Amazon Dashboard
		OnIntent: GetPillsHandler,
		OnLaunch: GetPillsHandler,
	},
}

func main() {
	alexa.Run(Applications, os.Getenv("PORT"))
}

func GetPillsHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	symptom, err := echoReq.GetSlotValue("Symptom")
	if err != nil {
		echoResp.OutputSpeech("I couldn't understand your symptoms!")
		return
	}
	echoResp.OutputSpeech("What's up Adri mother fog r, you got a " + symptom + ", right?")
}
