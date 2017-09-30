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




// If an item is a Plug, its DestinyInventoryItemDefinition.plug property will be populated with an instance of one of these bad boys.  This gives information about when it can be inserted, what the plug's category is (and thus whether it is compatible with a socket... see DestinySocketTypeDefinition for information about Plug Categories and socket compatibility), whether it is enabled and other Plug info.
type DestinyDefinitionsItemsDestinyItemPlugDefinition struct {
	// The rules around when this plug can be inserted into a socket, aside from the socket's individual restrictions.  The live data DestinyItemPlugComponent.insertFailIndexes will be an index into this array, so you can pull out the failure strings appropriate for the user.
	InsertionRules []DestinyDefinitionsItemsDestinyPlugRuleDefinition `json:"insertionRules,omitempty"`
	// The string identifier for the plug's category. Use the socket's DestinySocketTypeDefinition.plugWhitelist to determine whether this plug can be inserted into the socket.
	PlugCategoryIdentifier string `json:"plugCategoryIdentifier,omitempty"`
	// The hash for the plugCategoryIdentifier. You can use this instead if you wish: I put both in the definition for debugging purposes.
	PlugCategoryHash int32 `json:"plugCategoryHash,omitempty"`
	// If you successfully socket the item, this will determine whether or not you get \"refunded\" on the plug.
	OnActionRecreateSelf bool `json:"onActionRecreateSelf,omitempty"`
	// If inserting this plug requires materials, this is the hash identifier for looking up the DestinyMaterialRequirementSetDefinition for those requirements.
	InsertionMaterialRequirementHash int32 `json:"insertionMaterialRequirementHash,omitempty"`
	// In the game, if you're inspecting a plug item directly, this will be the item shown with the plug attached. Look up the DestinyInventoryItemDefinition for this hash for the item.
	PreviewItemOverrideHash int32 `json:"previewItemOverrideHash,omitempty"`
	// It's not enough for the plug to be inserted. It has to be enabled as well. For it to be enabled, it may require materials. This is the hash identifier for the DestinyMaterialRequirementSetDefinition for those requirements, if there is one.
	EnabledMaterialRequirementHash int32 `json:"enabledMaterialRequirementHash,omitempty"`
	// The rules around whether the plug, once inserted, is enabled and providing its benefits.  The live data DestinyItemPlugComponent.enableFailIndexes will be an index into this array, so you can pull out the failure strings appropriate for the user.
	EnabledRules []DestinyDefinitionsItemsDestinyPlugRuleDefinition `json:"enabledRules,omitempty"`
}

