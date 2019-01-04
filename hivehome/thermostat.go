package hivehome

import "github.com/tidwall/gjson"

func (c *client) GetThermostatIDForZone(zoneName string) string {

	allNodes := c.GetAllNodes()

	thermostatParentID := gjson.Get(allNodes, "nodes.#[attributes.zoneName.reportedValue=="+zoneName+"].id")

	thermostatID := gjson.Get(allNodes, "nodes.#[parentNodeId=="+thermostatParentID.String()+"].id")

	return thermostatID.String()
}
