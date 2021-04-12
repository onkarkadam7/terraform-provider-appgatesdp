/*
 * Appgate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the Appgate SDP Controller.    Please refer to the Integration chapter in the manual or contact Appgate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the Admin Interface (default port 8443) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-admin-interface-configure.html)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Admin Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:8443/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v14+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 14
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"time"
)

// BlacklistEntry struct for BlacklistEntry
type BlacklistEntry struct {
	User
	// The date and time of the blacklisting.
	BlacklistedAt *time.Time `json:"blacklistedAt,omitempty"`
	// The reason for blacklisting. The value is stored and logged.
	Reason *string `json:"reason,omitempty"`
}

// NewBlacklistEntry instantiates a new BlacklistEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlacklistEntry() *BlacklistEntry {
	this := BlacklistEntry{}
	return &this
}

// NewBlacklistEntryWithDefaults instantiates a new BlacklistEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlacklistEntryWithDefaults() *BlacklistEntry {
	this := BlacklistEntry{}
	return &this
}

// GetBlacklistedAt returns the BlacklistedAt field value if set, zero value otherwise.
func (o *BlacklistEntry) GetBlacklistedAt() time.Time {
	if o == nil || o.BlacklistedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.BlacklistedAt
}

// GetBlacklistedAtOk returns a tuple with the BlacklistedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlacklistEntry) GetBlacklistedAtOk() (*time.Time, bool) {
	if o == nil || o.BlacklistedAt == nil {
		return nil, false
	}
	return o.BlacklistedAt, true
}

// HasBlacklistedAt returns a boolean if a field has been set.
func (o *BlacklistEntry) HasBlacklistedAt() bool {
	if o != nil && o.BlacklistedAt != nil {
		return true
	}

	return false
}

// SetBlacklistedAt gets a reference to the given time.Time and assigns it to the BlacklistedAt field.
func (o *BlacklistEntry) SetBlacklistedAt(v time.Time) {
	o.BlacklistedAt = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BlacklistEntry) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlacklistEntry) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BlacklistEntry) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BlacklistEntry) SetReason(v string) {
	o.Reason = &v
}

func (o BlacklistEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedUser, errUser := json.Marshal(o.User)
	if errUser != nil {
		return []byte{}, errUser
	}
	errUser = json.Unmarshal([]byte(serializedUser), &toSerialize)
	if errUser != nil {
		return []byte{}, errUser
	}
	if o.BlacklistedAt != nil {
		toSerialize["blacklistedAt"] = o.BlacklistedAt
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableBlacklistEntry struct {
	value *BlacklistEntry
	isSet bool
}

func (v NullableBlacklistEntry) Get() *BlacklistEntry {
	return v.value
}

func (v *NullableBlacklistEntry) Set(val *BlacklistEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableBlacklistEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableBlacklistEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlacklistEntry(val *BlacklistEntry) *NullableBlacklistEntry {
	return &NullableBlacklistEntry{value: val, isSet: true}
}

func (v NullableBlacklistEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlacklistEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}