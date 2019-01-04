package hivehome

import (
	"io/ioutil"
	"net/url"

	"github.com/tidwall/gjson"
)

func (c *client) GetAllNodes() string {
	path := &url.URL{Path: "/omnia/nodes"}

	url := c.BaseURL.ResolveReference(path)

	_, _, rbody, _ := c.httpClient.Do("GET", url.String(), c.commonHeaders, nil)

	body, _ := ioutil.ReadAll(rbody)

	return string(body)
}

func (c *client) GetNodeAttributes(nodeID string) string {
	path := &url.URL{Path: "/omnia/nodes/" + nodeID}

	url := c.BaseURL.ResolveReference(path)

	_, _, rbody, _ := c.httpClient.Do("GET", url.String(), c.commonHeaders, nil)

	body, _ := ioutil.ReadAll(rbody)

	nodeAttributes := gjson.Get(string(body), "nodes.0.attributes")

	return nodeAttributes.String()
}
