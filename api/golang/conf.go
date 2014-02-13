// Loads the app configuration from a json file.
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type config struct{
	// oAuth1a parameters
	Consumer_key      string `json:"consumer_key"`
	Consumer_secret   string `json:"consumer_secret"`
	Request_token_url string `json:"request_token_url"`
	Access_token_url  string `json:"access_token_url"`
	Authorize_url     string `json:"authorize_url"`
	Callback_url      string `json:"callback_url"`
	// JWT secret for hash
	JWT_secret        string `json:jwt_secret`
	// local application settings
	App_url			  string `json:app_url`
	App_port		  string `json:app_port`
}

// Attempts to load the configuration file and may safely exit the program if the conf file is missing or corrupt.
func loadConfig(path string, fail bool) *config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("open config: ", err)
		if fail {
			os.Exit(1)
		}
	}

	temp := new(config)
	if err = json.Unmarshal(file, temp); err != nil {
		log.Println("parse config: ", err)
		if fail {
			os.Exit(1)
		}
	}
	return temp
}
