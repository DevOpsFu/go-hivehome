package hivehome

import (
	"io/ioutil"
	"net/url"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

func (c *Client) GetAllNodes() (string, error) {
	path := &url.URL{Path: "/omnia/nodes"}
	url := c.BaseURL.ResolveReference(path)

	c.checkSession()

	status, _, rbody, err := c.httpClient.Do("GET", url.String(), c.commonHeaders, nil)

	if err != nil {
		return "", errors.Wrap(err, "Error calling Hivehome nodes method")
	}

	if status.Code != 200 {
		return "", errors.New("Non-200 response received when retrieving nodes: " + status.String())
	}

	body, _ := ioutil.ReadAll(rbody)

	return string(body), nil
}

func (c *Client) GetNodeAttributes(nodeID string) (string, error) {
	path := &url.URL{Path: "/omnia/nodes/" + nodeID}
	url := c.BaseURL.ResolveReference(path)

	c.checkSession()

	status, _, rbody, err := c.httpClient.Do("GET", url.String(), c.commonHeaders, nil)

	if err != nil {
		return "", errors.Wrap(err, "Error calling Hivehome nodes method")
	}

	if status.Code != 200 {
		return "", errors.New("Non-200 response received when retrieving nodes: " + status.String())
	}

	body, _ := ioutil.ReadAll(rbody)
	nodeAttributes := gjson.Get(string(body), "nodes.0.attributes")

	return nodeAttributes.String(), nil
}
