package hivehome

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

type loginResponse struct {
	Sessions []*sessionResponse `json:"sessions"`
}

type sessionResponse struct {
	SessionID string `json:"sessionId"`
}

func (c *client) Login() error {

	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(c.sessionInfo)

	path := &url.URL{Path: "/omnia/auth/sessions"}

	url := c.BaseURL.ResolveReference(path)

	status, _, rbody, err := c.httpClient.Do("POST", url.String(), c.commonHeaders, b)

	if err != nil {
		return errors.Wrap(err, "Error calling Hivehome session API")
	}

	if status.Code != 200 {
		return errors.Wrap(errors.New("Non-200 response received when creating new Hivehome session"), status.String())
	}

	lr := new(loginResponse)

	err = json.NewDecoder(rbody).Decode(lr)

	if err != nil {
		return errors.Wrap(err, "Error when decoding new Session response body")
	}

	c.commonHeaders["X-Omnia-Access-Token"] = []string{lr.Sessions[0].SessionID}

	return nil
}
