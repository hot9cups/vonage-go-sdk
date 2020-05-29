/*
 * Voice API
 *
 * The Voice API lets you create outbound calls, control in-progress calls and get information about historical calls. More information about the Voice API can be found at <https://developer.nexmo.com/voice/voice-api/overview>.
 *
 * API version: 1.2.6
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetCallsResponse struct for GetCallsResponse
type GetCallsResponse struct {
	Count int32 `json:"count,omitempty"`
	PageSize int32 `json:"page_size,omitempty"`
	RecordIndex int32 `json:"record_index,omitempty"`
	Links GetCallResponseLinks `json:"_links,omitempty"`
	Embedded GetCallsResponseEmbedded `json:"_embedded,omitempty"`
}
