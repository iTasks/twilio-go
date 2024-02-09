/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// FlexV1InsightsConversationalAiReportInsights struct for FlexV1InsightsConversationalAiReportInsights
type FlexV1InsightsConversationalAiReportInsights struct {
	// The Instance SID of the instance for which report insights is fetched
	InstanceSid *string `json:"instance_sid,omitempty"`
	// The type of report insights required to fetch.Like gauge,channel-metrics,queue-metrics
	ReportId *string `json:"report_id,omitempty"`
	// The start date from which report insights data is included
	PeriodStart *time.Time `json:"period_start,omitempty"`
	// The end date till report insights data is included
	PeriodEnd *time.Time `json:"period_end,omitempty"`
	// Updated time of the report insights
	Updated *time.Time `json:"updated,omitempty"`
	// List of report insights breakdown
	Insights *[]interface{} `json:"insights,omitempty"`
	// The URL of this resource
	Url *string `json:"url,omitempty"`
}