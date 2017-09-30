/* 
 * Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@bungie.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package api




// Used in a number of Destiny contracts to return data about an item stack and its quantity. Can optionally return an itemInstanceId if the item is instanced - in which case, the quantity returned will be 1. If it's not... uh, let me know okay? Thanks.
type DestinyDestinyItemQuantity struct {
	// The hash identifier for the item in question. Use it to look up the item's DestinyInventoryItemDefinition.
	ItemHash int32 `json:"itemHash,omitempty"`
	// If this quantity is referring to a specific instance of an item, this will have the item's instance ID. Normally, this will be null.
	ItemInstanceId int64 `json:"itemInstanceId,omitempty"`
	// The amount of the item needed/available depending on the context of where DestinyItemQuantity is being used.
	Quantity int32 `json:"quantity,omitempty"`
}

