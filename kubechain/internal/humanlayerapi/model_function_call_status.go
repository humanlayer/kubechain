/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package humanlayerapi

import (
	"encoding/json"
	"time"
)

// checks if the FunctionCallStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FunctionCallStatus{}

// FunctionCallStatus struct for FunctionCallStatus
type FunctionCallStatus struct {
	RequestedAt      NullableTime           `json:"requested_at,omitempty"`
	RespondedAt      NullableTime           `json:"responded_at,omitempty"`
	Approved         NullableBool           `json:"approved,omitempty"`
	Comment          NullableString         `json:"comment,omitempty"`
	UserInfo         map[string]interface{} `json:"user_info,omitempty"`
	SlackContext     map[string]interface{} `json:"slack_context,omitempty"`
	RejectOptionName NullableString         `json:"reject_option_name,omitempty"`
	SlackMessageTs   NullableString         `json:"slack_message_ts,omitempty"`
}

// NewFunctionCallStatus instantiates a new FunctionCallStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionCallStatus() *FunctionCallStatus {
	this := FunctionCallStatus{}
	return &this
}

// NewFunctionCallStatusWithDefaults instantiates a new FunctionCallStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionCallStatusWithDefaults() *FunctionCallStatus {
	this := FunctionCallStatus{}
	return &this
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetRequestedAt() time.Time {
	if o == nil || IsNil(o.RequestedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt.Get()
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedAt.Get(), o.RequestedAt.IsSet()
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasRequestedAt() bool {
	if o != nil && o.RequestedAt.IsSet() {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given NullableTime and assigns it to the RequestedAt field.
func (o *FunctionCallStatus) SetRequestedAt(v time.Time) {
	o.RequestedAt.Set(&v)
}

// SetRequestedAtNil sets the value for RequestedAt to be an explicit nil
func (o *FunctionCallStatus) SetRequestedAtNil() {
	o.RequestedAt.Set(nil)
}

// UnsetRequestedAt ensures that no value is present for RequestedAt, not even an explicit nil
func (o *FunctionCallStatus) UnsetRequestedAt() {
	o.RequestedAt.Unset()
}

// GetRespondedAt returns the RespondedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetRespondedAt() time.Time {
	if o == nil || IsNil(o.RespondedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RespondedAt.Get()
}

// GetRespondedAtOk returns a tuple with the RespondedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetRespondedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RespondedAt.Get(), o.RespondedAt.IsSet()
}

// HasRespondedAt returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasRespondedAt() bool {
	if o != nil && o.RespondedAt.IsSet() {
		return true
	}

	return false
}

// SetRespondedAt gets a reference to the given NullableTime and assigns it to the RespondedAt field.
func (o *FunctionCallStatus) SetRespondedAt(v time.Time) {
	o.RespondedAt.Set(&v)
}

// SetRespondedAtNil sets the value for RespondedAt to be an explicit nil
func (o *FunctionCallStatus) SetRespondedAtNil() {
	o.RespondedAt.Set(nil)
}

// UnsetRespondedAt ensures that no value is present for RespondedAt, not even an explicit nil
func (o *FunctionCallStatus) UnsetRespondedAt() {
	o.RespondedAt.Unset()
}

// GetApproved returns the Approved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetApproved() bool {
	if o == nil || IsNil(o.Approved.Get()) {
		var ret bool
		return ret
	}
	return *o.Approved.Get()
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetApprovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approved.Get(), o.Approved.IsSet()
}

// HasApproved returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasApproved() bool {
	if o != nil && o.Approved.IsSet() {
		return true
	}

	return false
}

// SetApproved gets a reference to the given NullableBool and assigns it to the Approved field.
func (o *FunctionCallStatus) SetApproved(v bool) {
	o.Approved.Set(&v)
}

// SetApprovedNil sets the value for Approved to be an explicit nil
func (o *FunctionCallStatus) SetApprovedNil() {
	o.Approved.Set(nil)
}

// UnsetApproved ensures that no value is present for Approved, not even an explicit nil
func (o *FunctionCallStatus) UnsetApproved() {
	o.Approved.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *FunctionCallStatus) SetComment(v string) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *FunctionCallStatus) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *FunctionCallStatus) UnsetComment() {
	o.Comment.Unset()
}

// GetUserInfo returns the UserInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetUserInfo() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetUserInfoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserInfo) {
		return map[string]interface{}{}, false
	}
	return o.UserInfo, true
}

// HasUserInfo returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasUserInfo() bool {
	if o != nil && !IsNil(o.UserInfo) {
		return true
	}

	return false
}

// SetUserInfo gets a reference to the given map[string]interface{} and assigns it to the UserInfo field.
func (o *FunctionCallStatus) SetUserInfo(v map[string]interface{}) {
	o.UserInfo = v
}

// GetSlackContext returns the SlackContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetSlackContext() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SlackContext
}

// GetSlackContextOk returns a tuple with the SlackContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetSlackContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SlackContext) {
		return map[string]interface{}{}, false
	}
	return o.SlackContext, true
}

// HasSlackContext returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasSlackContext() bool {
	if o != nil && !IsNil(o.SlackContext) {
		return true
	}

	return false
}

// SetSlackContext gets a reference to the given map[string]interface{} and assigns it to the SlackContext field.
func (o *FunctionCallStatus) SetSlackContext(v map[string]interface{}) {
	o.SlackContext = v
}

// GetRejectOptionName returns the RejectOptionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetRejectOptionName() string {
	if o == nil || IsNil(o.RejectOptionName.Get()) {
		var ret string
		return ret
	}
	return *o.RejectOptionName.Get()
}

// GetRejectOptionNameOk returns a tuple with the RejectOptionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetRejectOptionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RejectOptionName.Get(), o.RejectOptionName.IsSet()
}

// HasRejectOptionName returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasRejectOptionName() bool {
	if o != nil && o.RejectOptionName.IsSet() {
		return true
	}

	return false
}

// SetRejectOptionName gets a reference to the given NullableString and assigns it to the RejectOptionName field.
func (o *FunctionCallStatus) SetRejectOptionName(v string) {
	o.RejectOptionName.Set(&v)
}

// SetRejectOptionNameNil sets the value for RejectOptionName to be an explicit nil
func (o *FunctionCallStatus) SetRejectOptionNameNil() {
	o.RejectOptionName.Set(nil)
}

// UnsetRejectOptionName ensures that no value is present for RejectOptionName, not even an explicit nil
func (o *FunctionCallStatus) UnsetRejectOptionName() {
	o.RejectOptionName.Unset()
}

// GetSlackMessageTs returns the SlackMessageTs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FunctionCallStatus) GetSlackMessageTs() string {
	if o == nil || IsNil(o.SlackMessageTs.Get()) {
		var ret string
		return ret
	}
	return *o.SlackMessageTs.Get()
}

// GetSlackMessageTsOk returns a tuple with the SlackMessageTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FunctionCallStatus) GetSlackMessageTsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SlackMessageTs.Get(), o.SlackMessageTs.IsSet()
}

// HasSlackMessageTs returns a boolean if a field has been set.
func (o *FunctionCallStatus) HasSlackMessageTs() bool {
	if o != nil && o.SlackMessageTs.IsSet() {
		return true
	}

	return false
}

// SetSlackMessageTs gets a reference to the given NullableString and assigns it to the SlackMessageTs field.
func (o *FunctionCallStatus) SetSlackMessageTs(v string) {
	o.SlackMessageTs.Set(&v)
}

// SetSlackMessageTsNil sets the value for SlackMessageTs to be an explicit nil
func (o *FunctionCallStatus) SetSlackMessageTsNil() {
	o.SlackMessageTs.Set(nil)
}

// UnsetSlackMessageTs ensures that no value is present for SlackMessageTs, not even an explicit nil
func (o *FunctionCallStatus) UnsetSlackMessageTs() {
	o.SlackMessageTs.Unset()
}

func (o FunctionCallStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FunctionCallStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestedAt.IsSet() {
		toSerialize["requested_at"] = o.RequestedAt.Get()
	}
	if o.RespondedAt.IsSet() {
		toSerialize["responded_at"] = o.RespondedAt.Get()
	}
	if o.Approved.IsSet() {
		toSerialize["approved"] = o.Approved.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.UserInfo != nil {
		toSerialize["user_info"] = o.UserInfo
	}
	if o.SlackContext != nil {
		toSerialize["slack_context"] = o.SlackContext
	}
	if o.RejectOptionName.IsSet() {
		toSerialize["reject_option_name"] = o.RejectOptionName.Get()
	}
	if o.SlackMessageTs.IsSet() {
		toSerialize["slack_message_ts"] = o.SlackMessageTs.Get()
	}
	return toSerialize, nil
}

type NullableFunctionCallStatus struct {
	value *FunctionCallStatus
	isSet bool
}

func (v NullableFunctionCallStatus) Get() *FunctionCallStatus {
	return v.value
}

func (v *NullableFunctionCallStatus) Set(val *FunctionCallStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionCallStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionCallStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionCallStatus(val *FunctionCallStatus) *NullableFunctionCallStatus {
	return &NullableFunctionCallStatus{value: val, isSet: true}
}

func (v NullableFunctionCallStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionCallStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
