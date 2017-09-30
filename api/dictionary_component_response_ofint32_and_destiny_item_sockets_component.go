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





type DictionaryComponentResponseOfint32AndDestinyItemSocketsComponent struct {
	
	Data map[string]DestinyEntitiesItemsDestinyItemSocketsComponent `json:"data,omitempty"`
	
	Privacy *ComponentsComponentPrivacySetting `json:"privacy,omitempty"`
}

