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
)

// AdministrativeRole struct for AdministrativeRole
type AdministrativeRole struct {
	BaseEntity
	// Administrative privilege list.
	Privileges []AdministrativePrivilege `json:"privileges"`
}

// NewAdministrativeRole instantiates a new AdministrativeRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministrativeRole(privileges []AdministrativePrivilege, ) *AdministrativeRole {
	this := AdministrativeRole{}
	this.Privileges = privileges
	return &this
}

// NewAdministrativeRoleWithDefaults instantiates a new AdministrativeRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministrativeRoleWithDefaults() *AdministrativeRole {
	this := AdministrativeRole{}
	return &this
}

// GetPrivileges returns the Privileges field value
func (o *AdministrativeRole) GetPrivileges() []AdministrativePrivilege {
	if o == nil  {
		var ret []AdministrativePrivilege
		return ret
	}

	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *AdministrativeRole) GetPrivilegesOk() (*[]AdministrativePrivilege, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value
func (o *AdministrativeRole) SetPrivileges(v []AdministrativePrivilege) {
	o.Privileges = v
}

func (o AdministrativeRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEntity, errBaseEntity := json.Marshal(o.BaseEntity)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	errBaseEntity = json.Unmarshal([]byte(serializedBaseEntity), &toSerialize)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	if true {
		toSerialize["privileges"] = o.Privileges
	}
	return json.Marshal(toSerialize)
}

type NullableAdministrativeRole struct {
	value *AdministrativeRole
	isSet bool
}

func (v NullableAdministrativeRole) Get() *AdministrativeRole {
	return v.value
}

func (v *NullableAdministrativeRole) Set(val *AdministrativeRole) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministrativeRole) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministrativeRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministrativeRole(val *AdministrativeRole) *NullableAdministrativeRole {
	return &NullableAdministrativeRole{value: val, isSet: true}
}

func (v NullableAdministrativeRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministrativeRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
