package main

import (
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
	alexa.Run(Applications, "3000")
}

func GetPillsHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("Hello world from my new Echo test app!").Card("Hello World", "This is a test card.")
}
