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





type SearchResultOfCommunityLiveStatus struct {
	
	Results []CommunityCommunityLiveStatus `json:"results,omitempty"`
	
	TotalResults int32 `json:"totalResults,omitempty"`
	
	HasMore bool `json:"hasMore,omitempty"`
	
	Query *QueriesPagedQuery `json:"query,omitempty"`
	
	ReplacementContinuationToken string `json:"replacementContinuationToken,omitempty"`
	// If useTotalResults is true, then totalResults represents an accurate count.  If False, it does not, and may be estimated/only the size of the current page.  Either way, you should probably always only trust hasMore.  This is a long-held historical throwback to when we used to do paging with known total results. Those queries toasted our database, and we were left to hastily alter our endpoints and create backward- compatible shims, of which useTotalResults is one.
	UseTotalResults bool `json:"useTotalResults,omitempty"`
}

