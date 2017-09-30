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




// Defines a specific stat value on an item, and the minimum/maximum range that we could compute for the item based on our heuristics for how the item might be generated.  Not guaranteed to match real-world instances of the item, but should hopefully at least be close. If it's not close, let us know on the Bungie API forums.
type DestinyDefinitionsDestinyInventoryItemStatDefinition struct {
	// The hash for the DestinyStatDefinition representing this stat.
	StatHash int32 `json:"statHash,omitempty"`
	// This value represents the stat value assuming the minimum possible roll but accounting for any mandatory bonuses that should be applied to the stat on item creation.  In Destiny 1, this was different from the \"minimum\" value because there were certain conditions where an item could be theoretically lower level/value than the initial roll.   In Destiny 2, this is not possible unless Talent Grids begin to be used again for these purposes or some other system change occurs... thus in practice, value and minimum should be the same in Destiny 2. Good riddance.
	Value int32 `json:"value,omitempty"`
	// The minimum possible value for this stat that we think the item can roll.
	Minimum int32 `json:"minimum,omitempty"`
	// The maximum possible value for this stat that we think the item can roll.
	Maximum int32 `json:"maximum,omitempty"`
}
