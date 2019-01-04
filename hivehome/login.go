package hivehome

import (
	"bytes"
	"encoding/json"
	"net/url"
)

type loginResponse struct {
	Sessions []*sessionResponse `json:"sessions"`
}

type sessionResponse struct {
	SessionID string `json:"sessionId"`
}

func (c *client) Login() {

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(c.sessionInfo)

	path := &url.URL{Path: "/omnia/auth/sessions"}

	url := c.BaseURL.ResolveReference(path)

	_, _, rbody, _ := c.httpClient.Do("POST", url.String(), c.commonHeaders, b)

	lr := new(loginResponse)

	json.NewDecoder(rbody).Decode(lr)

	c.commonHeaders["X-Omnia-Access-Token"] = []string{lr.Sessions[0].SessionID}
}
