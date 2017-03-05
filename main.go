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

	if evRes.Status == "error_cooldown" {
		return "Please, wait before taking another " + med + ", you already took one very recently"
	}

	return "Please, take a " + med + " for your " + symptom
}

func GetPillsHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	symptom, err := echoReq.GetSlotValue("Symptom")
	if err != nil {
		fmt.Println("Error", err)
		echoResp.OutputSpeech("I couldn't understand your symptoms!")
		return
	}

	switch symptom {
	case "headache", "head ache", "head":
		echoResp.OutputSpeech(callMedbayApi("paracetamol", "headache"))
		break

	case "stomachache", "stomach ache", "stomach":
		echoResp.OutputSpeech(callMedbayApi("omeprazol", "stomachache"))
		break

	default:
		fmt.Println("Not recognized symptom", symptom)
		echoResp.OutputSpeech("I couldn't understand your symptoms!")
	}
}
