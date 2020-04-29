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

// ApplianceAllOfIotConnectorClients struct for ApplianceAllOfIotConnectorClients
type ApplianceAllOfIotConnectorClients struct {
	// Name for the client. It will be mapped to the user claim 'clientName'.
	Name string `json:"name"`
	// The device ID to assign to this client. It will be used to generate device distinguished name.
	DeviceId *string `json:"deviceId,omitempty"`
	// Source configuration to allow via iptables.
	Sources *[]map[string]interface{} `json:"sources,omitempty"`
	// Use Source NAT for IoT client tunnel.
	Snat *bool `json:"snat,omitempty"`
}

// NewApplianceAllOfIotConnectorClients instantiates a new ApplianceAllOfIotConnectorClients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceAllOfIotConnectorClients(name string, ) *ApplianceAllOfIotConnectorClients {
	this := ApplianceAllOfIotConnectorClients{}
	this.Name = name
	var snat bool = true
	this.Snat = &snat
	return &this
}

// NewApplianceAllOfIotConnectorClientsWithDefaults instantiates a new ApplianceAllOfIotConnectorClients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceAllOfIotConnectorClientsWithDefaults() *ApplianceAllOfIotConnectorClients {
	this := ApplianceAllOfIotConnectorClients{}
	var snat bool = true
	this.Snat = &snat
	return &this
}

// GetName returns the Name field value
func (o *ApplianceAllOfIotConnectorClients) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfIotConnectorClients) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplianceAllOfIotConnectorClients) SetName(v string) {
	o.Name = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ApplianceAllOfIotConnectorClients) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfIotConnectorClients) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ApplianceAllOfIotConnectorClients) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ApplianceAllOfIotConnectorClients) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *ApplianceAllOfIotConnectorClients) GetSources() []map[string]interface{} {
	if o == nil || o.Sources == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfIotConnectorClients) GetSourcesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *ApplianceAllOfIotConnectorClients) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []map[string]interface{} and assigns it to the Sources field.
func (o *ApplianceAllOfIotConnectorClients) SetSources(v []map[string]interface{}) {
	o.Sources = &v
}

// GetSnat returns the Snat field value if set, zero value otherwise.
func (o *ApplianceAllOfIotConnectorClients) GetSnat() bool {
	if o == nil || o.Snat == nil {
		var ret bool
		return ret
	}
	return *o.Snat
}

// GetSnatOk returns a tuple with the Snat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceAllOfIotConnectorClients) GetSnatOk() (*bool, bool) {
	if o == nil || o.Snat == nil {
		return nil, false
	}
	return o.Snat, true
}

// HasSnat returns a boolean if a field has been set.
func (o *ApplianceAllOfIotConnectorClients) HasSnat() bool {
	if o != nil && o.Snat != nil {
		return true
	}

	return false
}

// SetSnat gets a reference to the given bool and assigns it to the Snat field.
func (o *ApplianceAllOfIotConnectorClients) SetSnat(v bool) {
	o.Snat = &v
}

func (o ApplianceAllOfIotConnectorClients) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Snat != nil {
		toSerialize["snat"] = o.Snat
	}
	return json.Marshal(toSerialize)
}

type NullableApplianceAllOfIotConnectorClients struct {
	value *ApplianceAllOfIotConnectorClients
	isSet bool
}

func (v NullableApplianceAllOfIotConnectorClients) Get() *ApplianceAllOfIotConnectorClients {
	return v.value
}

func (v *NullableApplianceAllOfIotConnectorClients) Set(val *ApplianceAllOfIotConnectorClients) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceAllOfIotConnectorClients) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceAllOfIotConnectorClients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceAllOfIotConnectorClients(val *ApplianceAllOfIotConnectorClients) *NullableApplianceAllOfIotConnectorClients {
	return &NullableApplianceAllOfIotConnectorClients{value: val, isSet: true}
}

func (v NullableApplianceAllOfIotConnectorClients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceAllOfIotConnectorClients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
