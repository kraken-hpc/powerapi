/*
 * Power API v1.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

//Response return a ImplResponse struct filled
func Response(code int, body interface{}) ImplResponse {
	return ImplResponse{Code: code, Body: body}
}

