package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	sns "github.com/robbiet480/go.sns"
)

func snsEndpoint(w http.ResponseWriter, r *http.Request) {
	var notificationPayload sns.Payload

	requestPayload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body")
		log.Print(err)
		w.WriteHeader(400)
		return
	}

	err = json.Unmarshal(requestPayload, &notificationPayload)
	if err != nil {
		log.Println("Error unmarshalling request payload")
		log.Print(err)
		w.WriteHeader(400)
		return
	}

	err = notificationPayload.VerifyPayload()
	if err != nil {
		log.Println("Error verifying request payload")
		log.Print(err)
		w.WriteHeader(400)
		return
	}

	if notificationPayload.SubscribeURL != "" {
		_, err := notificationPayload.Subscribe()
		if err != nil {
			log.Print(err)
			w.WriteHeader(400)
			return
		}
		log.Println("Subscribe successful")
	} else {
		putRecord([]byte(notificationPayload.Message), notificationPayload.MessageId)
		log.Println("Request processed")
	}

	w.WriteHeader(200)
	return
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/sns", snsEndpoint)

	portstr, ok := os.LookupEnv("PORT")
	if !ok {
		portstr = "8080"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", portstr), serverMux)
	if err != nil {
		log.Fatal(err)
	}
	return
}
