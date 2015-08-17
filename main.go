package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/google/google-api-go-client/plus/v1"
)

const key = "service_account_key.json"

func main() {
	data, err := ioutil.ReadFile(key)
	if err != nil {
		log.Fatal(err)
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/plus.profile.emails.read")
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext)
	plusService, err := plus.New(client)
	if err != nil {
		log.Fatal(err)
	}

	peopleService := plus.NewPeopleService(plusService)
	call := peopleService.Get("105303214497629424637")
	person, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(person)
}
