package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"encoding/json"

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

type EventAddResponse struct {
	Status string `json:"status"`
}

func callMedbayApi(med, symptom string) string {
	res, err := http.Get("https://medbay.scalingo.io/api/events/add?event=" + med)
	if err != nil {
		fmt.Println("Error calling medbay API")
		return "Oops, something went wrong, please try again in a few seconds"
	}

	defer res.Body.Close()
	var evRes EventAddResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading API response body")
		return "Oops, something went wrong, please try again in a few seconds"
	}

	err = json.Unmarshal(body, &evRes)
	if err != nil {
		fmt.Println("Error unmarshaling API response")
		return "Oops, something went wrong, please try again in a few seconds"
	}

	return "Please, take an " + med + " for your " + symptom
}

func GetPillsHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	symptom, err := echoReq.GetSlotValue("Symptom")
	if err != nil {
		echoResp.OutputSpeech("I couldn't understand your symptoms!")
		return
	}

	switch symptom {
	case "headache":
		echoResp.OutputSpeech(callMedbayApi("ibuprofen", "headache"))
		break

	default:
		echoResp.OutputSpeech("I couldn't understand your symptoms!")
	}
}
