package hivehome

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/tidwall/gjson"
)

func (c *client) GetAllNodes() (string, error) {
	path := &url.URL{Path: "/omnia/nodes"}

	url := c.BaseURL.ResolveReference(path)

	status, _, rbody, err := c.httpClient.Do("GET", url.String(), c.commonHeaders, nil)

	if err != nil {
		return "", fmt.Errorf("Error calling Hivehome nodes method: %+v", err)
	}

	if status.Code != 200 {
		return "", fmt.Errorf("Non-200 response received when retrieving nodes: %+v", status.String())
	}

	body, _ := ioutil.ReadAll(rbody)

	return string(body), nil
}

func (c *client) GetNodeAttributes(nodeID string) (string, error) {
	path := &url.URL{Path: "/omnia/nodes/" + nodeID}

	url := c.BaseURL.ResolveReference(path)

	status, _, rbody, err := c.httpClient.Do("GET", url.String(), c.commonHeaders, nil)

	if err != nil {
		return "", fmt.Errorf("Error calling Hivehome nodes method: %+v", err)
	}

	if status.Code != 200 {
		return "", fmt.Errorf("Non-200 response received when retrieving nodes: %+v", status.String())
	}

	body, _ := ioutil.ReadAll(rbody)

	nodeAttributes := gjson.Get(string(body), "nodes.0.attributes")

	return nodeAttributes.String(), nil
}
