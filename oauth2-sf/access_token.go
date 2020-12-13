/*

Now we are going to create two new endpoints for our local machine.
Make sure to have your RedirectURL set in your Salesforce Connected App configuration.
Documentation claims you need to use https, but http seems to work.
Understand that by not using https you put your data at risk. If this concerns you, you can look into setting up TLS for your local machine.
For development and learning purposes, you should be fine using a throwaway Connected App and http.

Please reference the Access Token component file.

Our two endpoints will be:
/ - our "homepage" where users are redirected to authorize with Salesforce.
/oauth2 - our redirect endpoint that will take an authorization code and exchange it for an access token.

Note that these endpoints could be named anything, what matters is how they interact.

We tie the / endpoint to the HomePage function (represented by http.HandleFunc("/", HomePage)). This function creates a URL the AuthURL defined in the oauth2.Config with the addition state string. In practice the state will be a random string that is generated per session. This state string is used to help correlating the authorization code to the session as well as confirming the legitimacy of any requests to the /oauth2 endpoint. The /oauth2 endpoint is tied to the Authorize function. The Authorize function does the following:
Parses the sent form.
Validates the returned state string is the expected state string.
Validate that we received an authorization code.
Exchanges the code for an access token. This exchange uses the TokenURL specified in the oauth2.Config.

From here you will have a valid access token.
*/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

const oauth2URL = "https://login.salesforce.com/services/oauth2/"

var (
	cfg oauth2.Config
)

// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Hit!")
	u := cfg.AuthCodeURL("my_random_state")
	http.Redirect(w, r, u, http.StatusFound)
}

// Authorize
func Authorize(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	state := r.Form.Get("state")
	if state != "my_random_state" {
		http.Error(w, "State invalid", http.StatusBadRequest)
		return
	}

	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	tokenData, err := cfg.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v\n", tokenData)
}

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

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/oauth2", Authorize)

	// We start up our Client on port 9094
	log.Println("Client is running at 9094 port.")
	log.Fatal(http.ListenAndServe(":9094", nil))
}
