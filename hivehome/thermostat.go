package hivehome

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func (c *client) GetThermostatIDForZone(zoneName string) (string, error) {

	allNodes, err := c.GetAllNodes()

	if err != nil {
		return "", fmt.Errorf("Error getting Thermostat ID: %+v", err)
	}

	thermostatParentID := gjson.Get(allNodes, "nodes.#[attributes.zoneName.reportedValue=="+zoneName+"].id")

	thermostatID := gjson.Get(allNodes, "nodes.#[parentNodeId=="+thermostatParentID.String()+"].id")

	return thermostatID.String(), nil
}
