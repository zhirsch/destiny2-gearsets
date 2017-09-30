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




// Information about the category and items currently sold in that category.
type DestinyEntitiesVendorsDestinyVendorCategory struct {
	// An index into the DestinyVendorDefinition.categories property, so you can grab the display data for this category.
	CategoryIndex int32 `json:"categoryIndex,omitempty"`
	// An ordered list of indexes into items being sold in this category (DestinyVendorDefinition.itemList) which will contain more information about the items being sold themselves. Can also be used to index into DestinyVendorSaleItemComponent data, if you asked for that data to be returned.
	ItemIndexes []int32 `json:"itemIndexes,omitempty"`
}
