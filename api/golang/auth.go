// Handles the initial authorization of a fitbit user for the app.
package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)


type authorizationUrl struct {
	Url string
}
type jwtToken struct {
	Token string
}
// Generates and returns the Authorization Url so the user can be redirected to fitibit to authorize app.
func authorize(w http.ResponseWriter, r *http.Request) {
	httpClient := new(http.Client)
	err := authcfg.GetRequestToken(authsvc, httpClient)

	if err != nil {
		log.Println("request token error: ", err)
		b, _ := json.Marshal(err)
		w.Write(b)
		return
	}

	url, err := authcfg.GetAuthorizeURL(authsvc)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		log.Println("auth url error: ", err)
		b, _ := json.Marshal(err)
		w.Write(b)
	}else {
		resp := authorizationUrl{url}
		b, _ := json.Marshal(resp)
		w.Write(b)
	}

}
// Handles the fitbit authorization callback.  Retrieves the Access Token and Access Token Secret
// used to create the JWT that will be used for authorization on subsequent requests.
func authorized(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	oauthToken := q.Get("oauth_token")
	oauthVerifier := q.Get("oauth_verifier")

	client := new(http.Client)
	err := authcfg.GetAccessToken(oauthToken, oauthVerifier, authsvc, client)

	if err != nil{
		log.Println("get access token error: ", err)
		return
	}

	accessToken := authcfg.AccessTokenKey;
	accessTokenSecret := authcfg.AccessTokenSecret

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["access_token"] = accessToken
	token.Claims["access_token_secret"] = accessTokenSecret
	sb := []byte(configuration.JWT_secret)
	tokenString, _ := token.SignedString(sb)

	// App will allow the jwt to reside in a cookie while we redirect to the app root.
	// Cookie will be persist. Alternative would be session or url parameter.
	c := http.Cookie{
		Name: "jwt",
		Value: tokenString,
		Path: "/",
	}

	http.SetCookie(w, &c)
	http.Redirect(w, r, configuration.App_url, 302)

}
// TODO: user friendly unauthorized messaging.
func unauthorized(w http.ResponseWriter, r *http.Request) {
	log.Println("unauthorized")
}
