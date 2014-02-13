// Wraps the fibit api calls and handles the oauth authorization construction.
package main

import (
	"bytes"
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/dgrijalva/jwt-go"
	"github.com/kurrik/oauth1a"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Asks for the JWT from the authorization header, decodes the jwt
// and returns the access token parameters required for the request.
func getAccessToken(r *http.Request) (string, string, error) {
	token, err := jwt.ParseFromRequest(r, func(t *jwt.Token) ([]byte, error) {
			return []byte(configuration.JWT_secret), nil
		})

	if err != nil {
		log.Println("error getting access token: ", err)
		return "", "", err
	}
	at := token.Claims["access_token"].(string)
	ats := token.Claims["access_token_secret"].(string)

	return at, ats, nil
}
// Builds a valid oAuth 1.0a HttpRequest to send to the fitbit api.
func buildHttpRequest(method string, url string, contentType string, token string, secret string, body io.Reader) (*http.Request, error) {
	r, _ := http.NewRequest(method, url, body)

	if contentType != "" {
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	usercfg := oauth1a.NewAuthorizedConfig(token, secret)
	authsvc.Sign(r, usercfg)

	return r, nil
}

// Retrieves a list of user devices.
func getDevices(w http.ResponseWriter, r *http.Request) {
	token, secret, _ := getAccessToken(r)
	rr, _ := buildHttpRequest("GET", "http://api.fitbit.com/1/user/-/devices.json", "", token, secret, nil)

	client := new(http.Client)
	resp, err := client.Do(rr)

	if err != nil {
		log.Println("get devices error: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	io.Copy(w, resp.Body)
}

// Retrieves a list of the user's current alarms for a given device.
func getAlarms(w http.ResponseWriter, r *http.Request, params martini.Params) {
	token, secret, _ := getAccessToken(r)
	rr, _ := buildHttpRequest("GET", "http://api.fitbit.com/1/user/-/devices/tracker/"+params["deviceId"]+"/alarms.json", "", token, secret, nil)

	client := new(http.Client)
	resp, err := client.Do(rr)

	if err != nil {
		log.Println("get alarms error: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	io.Copy(w, resp.Body)
}

// Sets a new alarm for the device.  Requires the user to sync (manually or automatically by dongle).
func setAlarm(w http.ResponseWriter, r *http.Request) {
	token, secret, _ := getAccessToken(r)

	var a alarm;
	body, _ := ioutil.ReadAll(r.Body)
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.Decode(&a)

	data := url.Values{}
	data.Add("time", a.Time)
	data.Add("enabled", "true")
	data.Add("recurring", "false")
	data.Add("weekDays", "[]")
	data.Add("label", a.Label)

	rr, _ := buildHttpRequest("POST", "http://api.fitbit.com/1/user/-/devices/tracker/"+a.DeviceId+"/alarms.json", "application/x-www-form-urlencoded", token, secret, strings.NewReader(data.Encode()))

	client := new(http.Client)
	resp, err := client.Do(rr)

	if err != nil {
		log.Println("set alarm error: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	io.Copy(w, resp.Body)
}
// Deletes a device alarm.  Requires the user to sync (manually or automatically by dongle).
func deleteAlarm(w http.ResponseWriter, r *http.Request, params martini.Params) {
	token, secret, _ := getAccessToken(r)

	rr, _ := buildHttpRequest("DELETE", "http://api.fitbit.com/1/user/-/devices/tracker/"+params["deviceId"]+"/alarms/"+params["alarmId"]+".json", "", token, secret, nil)

	client := new(http.Client)

	resp, err := client.Do(rr)

	if err != nil {
		log.Println("delete alarm error: ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	io.Copy(w, resp.Body)
}
