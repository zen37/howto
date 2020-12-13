/*
The oauth2 package requires that you set parameters in the Config structure to use the OAuth2 flow. Here are the fields with descriptions relating to Salesforce:

ClientID - this will be your Salesforce Connected App consumer key.
ClientSecret - this will be your Salesforce Connected App consumer secret.
Scopes - these are the scopes that will be added to the access token.
		Try to use the lowest privilege scopes whenever possible as it helps mitigate any risks to the user.
RedirectURL - this is the URL that will receive the authorization code after the request has been authorized. For development purposes we will use our local machine.
Endpoint:
AuthURL - this is the URL used to authorize requests. This URL will validate the user and the Connected App's consumer key, secret, and scopes. Once validated, an authorization code will be sent to the RedirectURL.
TokenURL - this is the URL that will exchange an authorization code for an access token.

Take the code from OAuth2 Configuration component and review all the fields in the Config structure.
You will have to fill in your own consumer key and secret.
Instead of hardcoding your key and secret, come up with another way to ingest these values into your code.
You can try environment variables, reading from a file, or using a secrets manager.
*/
package main

import (
	"fmt"

	"golang.org/x/oauth2"
)

const oauth2URL = "https://login.salesforce.com/services/oauth2/"

var (
	cfg oauth2.Config
)

func main() {
	cfg = oauth2.Config{
		ClientID:     "xxxxx",
		ClientSecret: "xxxxx",
		Scopes:       []string{"full"},
		RedirectURL:  "http://localhost:9094/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  oauth2URL + "authorize",
			TokenURL: oauth2URL + "token",
		},
	}

	fmt.Printf("%+v\n", cfg)
}
