package main

import (
	"os"

	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var Applications = map[string]interface{}{
	"/echo/getpills": alexa.EchoApplication{ // Route
		AppID:    "FakeEcho", // Echo App ID from Amazon Dashboard
		OnIntent: GetPillsHandler,
		OnLaunch: GetPillsHandler,
	},
}

func main() {
	alexa.Run(Applications, os.Getenv("PORT"))
}

func GetPillsHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoResp.OutputSpeech("What's up Adri mother fog r").Card("Wtf", "Is this")
}
