/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.3
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices struct for PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices
type PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices struct {
	BasePrice    float32  `json:"base_price,omitempty"`
	CurrentPrice float32  `json:"current_price,omitempty"`
	FriendlyName string   `json:"friendly_name,omitempty"`
	Prefixes     []string `json:"prefixes,omitempty"`
}
