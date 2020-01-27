package hivehome

import (
	"bytes"
	"encoding/json"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type loginResponse struct {
	Sessions []*sessionResponse `json:"sessions"`
}

type sessionResponse struct {
	SessionID string `json:"sessionId"`
}

func (c *Client) checkSession() error {

	if _, ok := c.commonHeaders["X-Omnia-Access-Token"]; !ok {
		return c.getToken()
	}

	t := c.commonHeaders["X-Omnia-Access-Token"][0]

	token, _, err := new(jwt.Parser).ParseUnverified(t, jwt.MapClaims{})
	if err != nil {
		return errors.Wrap(err, "Error parsing JWT")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.Wrap(err, "Error mapping token claims")
	}

	var tm time.Time
	expires := int64(claims["exp"].(float64))
	tm = time.Unix(expires, 0)

	timeRemaining := tm.Sub(time.Now())

	if timeRemaining.Minutes() < 5 {
		return c.getToken()
	}

	return nil
}

func (c *Client) getToken() error {
	b := new(bytes.Buffer)
	path := &url.URL{Path: "/omnia/auth/sessions"}
	url := c.BaseURL.ResolveReference(path)

	delete(c.commonHeaders, "X-Omnia-Access-Token")
	json.NewEncoder(b).Encode(c.sessionInfo)
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
