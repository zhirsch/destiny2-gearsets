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

import (
	"time"
)




type DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup struct {
	// Period for the group. If the stat periodType is day, then this will have a specific day. If the type is monthly, then this value will be the first day of the applicable month. This value is not set when the periodType is 'all time'.
	Period time.Time `json:"period,omitempty"`
	// If the period group is for a specific activity, this property will be set.
	ActivityDetails *interface{} `json:"activityDetails,omitempty"`
	// Collection of stats for the period.
	Values map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"values,omitempty"`
}

