package auth

import (
	"github.com/ChicagoDSA/DSA-Events/payloads"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var (
	githubAuth *oauth2.Config = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Scopes: []string{
			"user:organizations",
		},
		Endpoint: github.Endpoint,
	}
)

func GithubInit(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger).WithField("auth", "GithubInit")

	// Read credentials from json file
	credsFile, err := ioutil.ReadFile("./creds.json")
	if err != nil {
		log.WithError(err).Fatal("Credentials find error.")
	}

	var creds struct {
		Id     string `json:"githubClientId"`
		Secret string `json:"githubClientSecret"`
	}

	if err = json.Unmarshal(credsFile, &creds); err != nil {
		log.WithError(err).Fatal("Credentials parsing error.")
	}

	githubAuth.ClientID = creds.Id
	githubAuth.ClientSecret = creds.Secret

	githubAuth.RedirectURL = "http://" + c.MustGet("host").(string) + ":" + c.MustGet("port").(string) + "/account/github/callback"

	// TODO: Replace "state" with rand gen string (CSRF)


	url := githubAuth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(302, url)
}

func GithubCallback(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger).WithField("auth", "GithubCallback")

	code := c.Query("code")

	// TODO: Verify "state" (rand gen string) before exchanging (should be query parameter)

	token, err := githubAuth.Exchange(oauth2.NoContext, code)
	if err != nil {
		//log.WithError(err).Fatal("Failed to exchange initial GitHub auth information")
		c.String(http.StatusUnauthorized, "GitHub failed to authenticate you as a user!")
	}

	client := githubAuth.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		log.WithError(err).Fatal("Error authenticating user with github")
	}
	body, err := ioutil.ReadAll(resp.Body)

	var githubUserResponse payloads.GithubUser

	err = json.Unmarshal(body, &githubUserResponse)
	if err != nil {
		log.WithError(err).Fatal("Error unmarshalling GitHub user response.")
	}

	resp, err = client.Get(githubUserResponse.OrganizationsUrl)
	if err != nil {
		log.WithError(err).Fatal("Error retrieving github org information for user")
	}
	body, err = ioutil.ReadAll(resp.Body)

	var githubOrgResponse []payloads.GithubOrg

	err = json.Unmarshal(body, &githubOrgResponse)
	if err != nil {
		log.WithError(err).Fatal("Error unmarshalling GitHub org response.")
	}

	var isCDSAMember = false
	for i := 0; i < len(githubOrgResponse); i++ {
		if githubOrgResponse[i].Login == "ChicagoDSA" {
			isCDSAMember = true
			break
		}
	}

	if isCDSAMember {
		c.String(http.StatusOK, token.AccessToken)
	} else {
		c.String(http.StatusUnauthorized, "You're not a CDSA GitHub member!\nContact: https://chicagodsa.org/contact/")
	}
}

func ValidateOAuthToken(authHeader string) bool {
	tokenIndex := strings.Index(authHeader, "token ")
	if tokenIndex < 0 {
		return false;
	}

	request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return false
	}

	request.Header.Set("Authorization", authHeader)
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)

	var githubUserResponse payloads.GithubUser

	err = json.Unmarshal(body, &githubUserResponse)
	if err != nil {
		return false
	}

	resp, err = client.Get(githubUserResponse.OrganizationsUrl)
	if err != nil {
		return false
	}
	body, err = ioutil.ReadAll(resp.Body)

	var githubOrgResponse []payloads.GithubOrg

	err = json.Unmarshal(body, &githubOrgResponse)
	if err != nil {
		return false
	}

	var isCDSAMember = false
	for i := 0; i < len(githubOrgResponse); i++ {
		if githubOrgResponse[i].Login == "ChicagoDSA" {
			isCDSAMember = true
			break
		}
	}

	return isCDSAMember
}
