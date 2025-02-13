/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.
 *
 * API version:
 * Contact: office@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateFedcmFlowResponse Contains a list of all available FedCM providers.
type CreateFedcmFlowResponse struct {
	CsrfToken *string    `json:"csrf_token,omitempty"`
	Providers []Provider `json:"providers,omitempty"`
}

// NewCreateFedcmFlowResponse instantiates a new CreateFedcmFlowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFedcmFlowResponse() *CreateFedcmFlowResponse {
	this := CreateFedcmFlowResponse{}
	return &this
}

// NewCreateFedcmFlowResponseWithDefaults instantiates a new CreateFedcmFlowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFedcmFlowResponseWithDefaults() *CreateFedcmFlowResponse {
	this := CreateFedcmFlowResponse{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *CreateFedcmFlowResponse) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFedcmFlowResponse) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *CreateFedcmFlowResponse) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *CreateFedcmFlowResponse) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *CreateFedcmFlowResponse) GetProviders() []Provider {
	if o == nil || o.Providers == nil {
		var ret []Provider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFedcmFlowResponse) GetProvidersOk() ([]Provider, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *CreateFedcmFlowResponse) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []Provider and assigns it to the Providers field.
func (o *CreateFedcmFlowResponse) SetProviders(v []Provider) {
	o.Providers = v
}

func (o CreateFedcmFlowResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFedcmFlowResponse struct {
	value *CreateFedcmFlowResponse
	isSet bool
}

func (v NullableCreateFedcmFlowResponse) Get() *CreateFedcmFlowResponse {
	return v.value
}

func (v *NullableCreateFedcmFlowResponse) Set(val *CreateFedcmFlowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFedcmFlowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFedcmFlowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFedcmFlowResponse(val *CreateFedcmFlowResponse) *NullableCreateFedcmFlowResponse {
	return &NullableCreateFedcmFlowResponse{value: val, isSet: true}
}

func (v NullableCreateFedcmFlowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFedcmFlowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
