/*
 * sdk_api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 4.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage Request body data model for CreateOrUpdateItemStockPhysicalShippingInfo API
type CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Country string `json:"country"`
	Email string `json:"email"`
	PostalCode string `json:"postalCode"`
	City string `json:"city"`
	State string `json:"state"`
	Address1 string `json:"address1"`
	PhoneNumber string `json:"phoneNumber"`
	Address2 NullableString `json:"address2"`
	Address3 NullableString `json:"address3"`
	RequestTimestamp string `json:"requestTimestamp"`
}

// NewCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage instantiates a new CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage(firstName string, lastName string, country string, email string, postalCode string, city string, state string, address1 string, phoneNumber string, address2 NullableString, address3 NullableString, requestTimestamp string) *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage {
	this := CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Country = country
	this.Email = email
	this.PostalCode = postalCode
	this.City = city
	this.State = state
	this.Address1 = address1
	this.PhoneNumber = phoneNumber
	this.Address2 = address2
	this.Address3 = address3
	this.RequestTimestamp = requestTimestamp
	return &this
}

// NewCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessageWithDefaults instantiates a new CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessageWithDefaults() *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage {
	this := CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetLastName(v string) {
	o.LastName = v
}

// GetCountry returns the Country field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetCountry(v string) {
	o.Country = v
}

// GetEmail returns the Email field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetEmail(v string) {
	o.Email = v
}

// GetPostalCode returns the PostalCode field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCity returns the City field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetState(v string) {
	o.State = v
}

// GetAddress1 returns the Address1 field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetAddress1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address1, true
}

// SetAddress1 sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetAddress1(v string) {
	o.Address1 = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetPhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetAddress2 returns the Address2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetAddress2() string {
	if o == nil || o.Address2.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address2.Get()
}

// GetAddress2Ok returns a tuple with the Address2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetAddress2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address2.Get(), o.Address2.IsSet()
}

// SetAddress2 sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetAddress2(v string) {
	o.Address2.Set(&v)
}

// GetAddress3 returns the Address3 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetAddress3() string {
	if o == nil || o.Address3.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address3.Get()
}

// GetAddress3Ok returns a tuple with the Address3 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetAddress3Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address3.Get(), o.Address3.IsSet()
}

// SetAddress3 sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetAddress3(v string) {
	o.Address3.Set(&v)
}

// GetRequestTimestamp returns the RequestTimestamp field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetRequestTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestTimestamp
}

// GetRequestTimestampOk returns a tuple with the RequestTimestamp field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) GetRequestTimestampOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestTimestamp, true
}

// SetRequestTimestamp sets field value
func (o *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) SetRequestTimestamp(v string) {
	o.RequestTimestamp = v
}

func (o CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["postalCode"] = o.PostalCode
	}
	if true {
		toSerialize["city"] = o.City
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["address1"] = o.Address1
	}
	if true {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if true {
		toSerialize["address2"] = o.Address2.Get()
	}
	if true {
		toSerialize["address3"] = o.Address3.Get()
	}
	if true {
		toSerialize["requestTimestamp"] = o.RequestTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage struct {
	value *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage
	isSet bool
}

func (v NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) Get() *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage {
	return v.value
}

func (v *NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) Set(val *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage(val *CreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) *NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage {
	return &NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage{value: val, isSet: true}
}

func (v NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateItemStockPhysicalShippingInfoRequestBodyMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


