// Handles program initialization.
package main

import (
	"github.com/kurrik/oauth1a"
)

var (
	configuration *config
	authsvc *oauth1a.Service
	authcfg *oauth1a.UserConfig
)

func init() {
	configuration = loadConfig("../conf/conf.json", true);

	authsvc = &oauth1a.Service{
		RequestURL:
		configuration.Request_token_url,
		AuthorizeURL: configuration.Authorize_url,
		AccessURL:    configuration.Access_token_url,
		ClientConfig: &oauth1a.ClientConfig{
			ConsumerKey:	configuration.Consumer_key,
			ConsumerSecret: configuration.Consumer_secret,
			CallbackURL:    configuration.Callback_url,
		},
		Signer: new(oauth1a.HmacSha1Signer),
	}

	authcfg = &oauth1a.UserConfig{}
}
