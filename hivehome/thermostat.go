package hivehome

import (
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

func (c *Client) GetThermostatIDForZone(zoneName string) (string, error) {

	allNodes, err := c.GetAllNodes()

	if err != nil {
		return "", errors.Wrap(err, "Error getting thermostat ID")
	}

	thermostatParentID := gjson.Get(allNodes, "nodes.#[attributes.zoneName.reportedValue=="+zoneName+"].id")
	thermostatID := gjson.Get(allNodes, "nodes.#[parentNodeId=="+thermostatParentID.String()+"].id")

	return thermostatID.String(), nil
}
