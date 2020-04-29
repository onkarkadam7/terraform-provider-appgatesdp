/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v12+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information abot the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 12
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"time"
)

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// The javascript expression to evaluate.
	Expression string `json:"expression"`
	UserClaims *map[string]map[string]interface{} `json:"userClaims,omitempty"`
	DeviceClaims *map[string]map[string]interface{} `json:"deviceClaims,omitempty"`
	SystemClaims *map[string]map[string]interface{} `json:"systemClaims,omitempty"`
	Time *time.Time `json:"time,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2(expression string, ) *InlineObject2 {
	this := InlineObject2{}
	this.Expression = expression
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// GetExpression returns the Expression field value
func (o *InlineObject2) GetExpression() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetExpressionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *InlineObject2) SetExpression(v string) {
	o.Expression = v
}

// GetUserClaims returns the UserClaims field value if set, zero value otherwise.
func (o *InlineObject2) GetUserClaims() map[string]map[string]interface{} {
	if o == nil || o.UserClaims == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.UserClaims
}

// GetUserClaimsOk returns a tuple with the UserClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetUserClaimsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.UserClaims == nil {
		return nil, false
	}
	return o.UserClaims, true
}

// HasUserClaims returns a boolean if a field has been set.
func (o *InlineObject2) HasUserClaims() bool {
	if o != nil && o.UserClaims != nil {
		return true
	}

	return false
}

// SetUserClaims gets a reference to the given map[string]map[string]interface{} and assigns it to the UserClaims field.
func (o *InlineObject2) SetUserClaims(v map[string]map[string]interface{}) {
	o.UserClaims = &v
}

// GetDeviceClaims returns the DeviceClaims field value if set, zero value otherwise.
func (o *InlineObject2) GetDeviceClaims() map[string]map[string]interface{} {
	if o == nil || o.DeviceClaims == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.DeviceClaims
}

// GetDeviceClaimsOk returns a tuple with the DeviceClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetDeviceClaimsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.DeviceClaims == nil {
		return nil, false
	}
	return o.DeviceClaims, true
}

// HasDeviceClaims returns a boolean if a field has been set.
func (o *InlineObject2) HasDeviceClaims() bool {
	if o != nil && o.DeviceClaims != nil {
		return true
	}

	return false
}

// SetDeviceClaims gets a reference to the given map[string]map[string]interface{} and assigns it to the DeviceClaims field.
func (o *InlineObject2) SetDeviceClaims(v map[string]map[string]interface{}) {
	o.DeviceClaims = &v
}

// GetSystemClaims returns the SystemClaims field value if set, zero value otherwise.
func (o *InlineObject2) GetSystemClaims() map[string]map[string]interface{} {
	if o == nil || o.SystemClaims == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.SystemClaims
}

// GetSystemClaimsOk returns a tuple with the SystemClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetSystemClaimsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.SystemClaims == nil {
		return nil, false
	}
	return o.SystemClaims, true
}

// HasSystemClaims returns a boolean if a field has been set.
func (o *InlineObject2) HasSystemClaims() bool {
	if o != nil && o.SystemClaims != nil {
		return true
	}

	return false
}

// SetSystemClaims gets a reference to the given map[string]map[string]interface{} and assigns it to the SystemClaims field.
func (o *InlineObject2) SetSystemClaims(v map[string]map[string]interface{}) {
	o.SystemClaims = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineObject2) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineObject2) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineObject2) SetTime(v time.Time) {
	o.Time = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["expression"] = o.Expression
	}
	if o.UserClaims != nil {
		toSerialize["userClaims"] = o.UserClaims
	}
	if o.DeviceClaims != nil {
		toSerialize["deviceClaims"] = o.DeviceClaims
	}
	if o.SystemClaims != nil {
		toSerialize["systemClaims"] = o.SystemClaims
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
