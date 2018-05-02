// Copyright (c) 2018 The Decred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	slackposter "github.com/m0t0k1ch1/go-slack-poster"
)

type AjaxResponse struct {
	Type string
	Text string
}

type Configuration struct {
	SLACK slackConfig
}

type slackConfig struct {
	APIKey  string `json:"apikey"`
	Channel string `json:"channel"`
}

const (
	defaultPort = ":8000"
)

var (
	client    *slackposter.Client
	config    Configuration
	emailForm = template.Must(template.ParseFiles("views/form.html"))

	re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func init() {
	config = LoadConfig("./config.json")
	if len(config.SLACK.APIKey) < 42 {
		log.Fatal("Error - Please check your api key in config file")
	}
	if len(config.SLACK.Channel) < 2 {
		log.Fatal("Error - Please check the channel name config file")
	}

	client = slackposter.NewClient(config.SLACK.APIKey)
}

func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}

	return config
}

func form(w http.ResponseWriter, r *http.Request) {
	err := emailForm.Execute(w, nil)
	if err != nil {
		http.Error(w, "Failed to parse files", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}

func view(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	// Process form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate Email
	// if Email is valid, create json success response
	email := r.FormValue("email")
	if len("email") <= 0 || !re.MatchString(email) {
		http.Error(w, "Invalid email address", http.StatusOK)
		return
	}

	// post the invite to slack channel
	err = client.SendMessage(context.Background(), config.SLACK.Channel, email, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to SendMessage", http.StatusInternalServerError)
		return
	}

	log.Print(email)
	ajaxresponse := AjaxResponse{"success", email}
	js, err := json.Marshal(ajaxresponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	// Routes
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("tpldcrslackinvite"))))
	http.HandleFunc("/", form)
	http.HandleFunc("/view", view)

	log.Printf("Starting server on %v", defaultPort)
	err := http.ListenAndServe(defaultPort, nil) // setting listening port
	if err != nil {
		log.Fatal(err)
	}
}
