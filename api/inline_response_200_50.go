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





type InlineResponse20050 struct {
	
	Response *SearchResultOfCommunityLiveStatus `json:"Response,omitempty"`
	
	ErrorCode *ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`
	
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
	
	ErrorStatus string `json:"ErrorStatus,omitempty"`
	
	Message string `json:"Message,omitempty"`
	
	MessageData map[string]string `json:"MessageData,omitempty"`
}

