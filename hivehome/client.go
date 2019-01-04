package hivehome

import (
	"net/url"

	"github.com/gorilla/http"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Caller   string `json:"caller"`
}

type sessionInfo struct {
	Sessions []credentials `json:"sessions"`
}

type client struct {
	BaseURL       *url.URL
	sessionID     string
	httpClient    *http.Client
	sessionInfo   sessionInfo
	commonHeaders map[string][]string
}

const (
	baseURL = "https://api-prod.bgchprod.info"
)

var commonHeaders = map[string][]string{
	"Content-Type":   []string{"application/vnd.alertme.zoo-6.1+json"},
	"Accept":         []string{"application/vnd.alertme.zoo-6.1+json"},
	"X-Omnia-Client": []string{"Hive Web Dashboard"},
}

func NewClient(username string, password string) *client {

	clientBaseURL, _ := url.Parse(baseURL)

	creds := credentials{Username: username, Password: password, Caller: "WEB"}

	sessInfo := sessionInfo{Sessions: []credentials{creds}}

	return &client{
		BaseURL:       clientBaseURL,
		httpClient:    &http.DefaultClient,
		sessionInfo:   sessInfo,
		commonHeaders: commonHeaders,
	}
}
