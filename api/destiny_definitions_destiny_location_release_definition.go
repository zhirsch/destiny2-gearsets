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




// A specific \"spot\" referred to by a location. Only one of these can be active at a time for a given Location.
type DestinyDefinitionsDestinyLocationReleaseDefinition struct {
	// Sadly, these don't appear to be populated anymore (ever?)
	DisplayProperties *interface{} `json:"displayProperties,omitempty"`
	// If we had map information, this spawnPoint would be interesting. But sadly, we don't have that info.
	SpawnPoint int32 `json:"spawnPoint,omitempty"`
	// The Destination being pointed to by this location.
	DestinationHash int32 `json:"destinationHash,omitempty"`
	// The Activity being pointed to by this location.
	ActivityHash int32 `json:"activityHash,omitempty"`
	// The Activity Graph being pointed to by this location.
	ActivityGraphHash int32 `json:"activityGraphHash,omitempty"`
	// The Activity Graph Node being pointed to by this location. (Remember that Activity Graph Node hashes are only unique within an Activity Graph: so use the combination to find the node being spoken of)
	ActivityGraphNodeHash int32 `json:"activityGraphNodeHash,omitempty"`
	// The Activity Bubble within the Destination. Look this up in the DestinyDestinationDefinition's bubbles and bubbleSettings properties.
	ActivityBubbleName int32 `json:"activityBubbleName,omitempty"`
	// If we had map information, this would tell us something cool about the path this location wants you to take. I wish we had map information.
	ActivityPathBundle int32 `json:"activityPathBundle,omitempty"`
	// If we had map information, this would tell us about path information related to destination on the map. Sad. Maybe you can do something cool with it. Go to town man.
	ActivityPathDestination int32 `json:"activityPathDestination,omitempty"`
	// The type of Nav Point that this represents. See the enumeration for more info.
	NavPointType *interface{} `json:"navPointType,omitempty"`
	// Looks like it should be the position on the map, but sadly it does not look populated... yet?
	WorldPosition []int32 `json:"worldPosition,omitempty"`
}
