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




// DestinyManifest is the external-facing contract for just the properties needed by those calling the Destiny Platform.
type DestinyConfigDestinyManifest struct {
	
	Version string `json:"version,omitempty"`
	
	MobileAssetContentPath string `json:"mobileAssetContentPath,omitempty"`
	
	MobileGearAssetDataBases []DestinyConfigGearAssetDataBaseDefinition `json:"mobileGearAssetDataBases,omitempty"`
	
	MobileWorldContentPaths map[string]string `json:"mobileWorldContentPaths,omitempty"`
	
	MobileClanBannerDatabasePath string `json:"mobileClanBannerDatabasePath,omitempty"`
	
	MobileGearCDN map[string]string `json:"mobileGearCDN,omitempty"`
}

