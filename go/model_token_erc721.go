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
	"time"
)

// TokenERC721 struct for TokenERC721
type TokenERC721 struct {
	Id string `json:"id"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	TokenStandardType TokenStandardType `json:"tokenStandardType"`
	ContractERC721Id string `json:"contractERC721Id"`
	TokenId float32 `json:"tokenId"`
	TokenURI string `json:"tokenURI"`
	Metadata map[string]interface{} `json:"metadata"`
	MintTransactionHash string `json:"mintTransactionHash"`
	InitialOwnerAddress string `json:"initialOwnerAddress"`
	CurrentOwnerAddress string `json:"currentOwnerAddress"`
	TransferHistory []TransferData `json:"transferHistory"`
}

// NewTokenERC721 instantiates a new TokenERC721 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenERC721(id string, createAt time.Time, updateAt time.Time, tokenStandardType TokenStandardType, contractERC721Id string, tokenId float32, tokenURI string, metadata map[string]interface{}, mintTransactionHash string, initialOwnerAddress string, currentOwnerAddress string, transferHistory []TransferData) *TokenERC721 {
	this := TokenERC721{}
	this.Id = id
	this.CreateAt = createAt
	this.UpdateAt = updateAt
	this.TokenStandardType = tokenStandardType
	this.ContractERC721Id = contractERC721Id
	this.TokenId = tokenId
	this.TokenURI = tokenURI
	this.Metadata = metadata
	this.MintTransactionHash = mintTransactionHash
	this.InitialOwnerAddress = initialOwnerAddress
	this.CurrentOwnerAddress = currentOwnerAddress
	this.TransferHistory = transferHistory
	return &this
}

// NewTokenERC721WithDefaults instantiates a new TokenERC721 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenERC721WithDefaults() *TokenERC721 {
	this := TokenERC721{}
	var tokenStandardType TokenStandardType = ERC721
	this.TokenStandardType = tokenStandardType
	return &this
}

// GetId returns the Id field value
func (o *TokenERC721) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TokenERC721) SetId(v string) {
	o.Id = v
}

// GetCreateAt returns the CreateAt field value
func (o *TokenERC721) GetCreateAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreateAt
}

// GetCreateAtOk returns a tuple with the CreateAt field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetCreateAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreateAt, true
}

// SetCreateAt sets field value
func (o *TokenERC721) SetCreateAt(v time.Time) {
	o.CreateAt = v
}

// GetUpdateAt returns the UpdateAt field value
func (o *TokenERC721) GetUpdateAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdateAt
}

// GetUpdateAtOk returns a tuple with the UpdateAt field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetUpdateAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateAt, true
}

// SetUpdateAt sets field value
func (o *TokenERC721) SetUpdateAt(v time.Time) {
	o.UpdateAt = v
}

// GetTokenStandardType returns the TokenStandardType field value
func (o *TokenERC721) GetTokenStandardType() TokenStandardType {
	if o == nil {
		var ret TokenStandardType
		return ret
	}

	return o.TokenStandardType
}

// GetTokenStandardTypeOk returns a tuple with the TokenStandardType field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetTokenStandardTypeOk() (*TokenStandardType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenStandardType, true
}

// SetTokenStandardType sets field value
func (o *TokenERC721) SetTokenStandardType(v TokenStandardType) {
	o.TokenStandardType = v
}

// GetContractERC721Id returns the ContractERC721Id field value
func (o *TokenERC721) GetContractERC721Id() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractERC721Id
}

// GetContractERC721IdOk returns a tuple with the ContractERC721Id field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetContractERC721IdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContractERC721Id, true
}

// SetContractERC721Id sets field value
func (o *TokenERC721) SetContractERC721Id(v string) {
	o.ContractERC721Id = v
}

// GetTokenId returns the TokenId field value
func (o *TokenERC721) GetTokenId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetTokenIdOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *TokenERC721) SetTokenId(v float32) {
	o.TokenId = v
}

// GetTokenURI returns the TokenURI field value
func (o *TokenERC721) GetTokenURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenURI
}

// GetTokenURIOk returns a tuple with the TokenURI field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetTokenURIOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenURI, true
}

// SetTokenURI sets field value
func (o *TokenERC721) SetTokenURI(v string) {
	o.TokenURI = v
}

// GetMetadata returns the Metadata field value
func (o *TokenERC721) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *TokenERC721) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMintTransactionHash returns the MintTransactionHash field value
func (o *TokenERC721) GetMintTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MintTransactionHash
}

// GetMintTransactionHashOk returns a tuple with the MintTransactionHash field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetMintTransactionHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MintTransactionHash, true
}

// SetMintTransactionHash sets field value
func (o *TokenERC721) SetMintTransactionHash(v string) {
	o.MintTransactionHash = v
}

// GetInitialOwnerAddress returns the InitialOwnerAddress field value
func (o *TokenERC721) GetInitialOwnerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialOwnerAddress
}

// GetInitialOwnerAddressOk returns a tuple with the InitialOwnerAddress field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetInitialOwnerAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InitialOwnerAddress, true
}

// SetInitialOwnerAddress sets field value
func (o *TokenERC721) SetInitialOwnerAddress(v string) {
	o.InitialOwnerAddress = v
}

// GetCurrentOwnerAddress returns the CurrentOwnerAddress field value
func (o *TokenERC721) GetCurrentOwnerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentOwnerAddress
}

// GetCurrentOwnerAddressOk returns a tuple with the CurrentOwnerAddress field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetCurrentOwnerAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentOwnerAddress, true
}

// SetCurrentOwnerAddress sets field value
func (o *TokenERC721) SetCurrentOwnerAddress(v string) {
	o.CurrentOwnerAddress = v
}

// GetTransferHistory returns the TransferHistory field value
func (o *TokenERC721) GetTransferHistory() []TransferData {
	if o == nil {
		var ret []TransferData
		return ret
	}

	return o.TransferHistory
}

// GetTransferHistoryOk returns a tuple with the TransferHistory field value
// and a boolean to check if the value has been set.
func (o *TokenERC721) GetTransferHistoryOk() (*[]TransferData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferHistory, true
}

// SetTransferHistory sets field value
func (o *TokenERC721) SetTransferHistory(v []TransferData) {
	o.TransferHistory = v
}

func (o TokenERC721) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["createAt"] = o.CreateAt
	}
	if true {
		toSerialize["updateAt"] = o.UpdateAt
	}
	if true {
		toSerialize["tokenStandardType"] = o.TokenStandardType
	}
	if true {
		toSerialize["contractERC721Id"] = o.ContractERC721Id
	}
	if true {
		toSerialize["tokenId"] = o.TokenId
	}
	if true {
		toSerialize["tokenURI"] = o.TokenURI
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["mintTransactionHash"] = o.MintTransactionHash
	}
	if true {
		toSerialize["initialOwnerAddress"] = o.InitialOwnerAddress
	}
	if true {
		toSerialize["currentOwnerAddress"] = o.CurrentOwnerAddress
	}
	if true {
		toSerialize["transferHistory"] = o.TransferHistory
	}
	return json.Marshal(toSerialize)
}

type NullableTokenERC721 struct {
	value *TokenERC721
	isSet bool
}

func (v NullableTokenERC721) Get() *TokenERC721 {
	return v.value
}

func (v *NullableTokenERC721) Set(val *TokenERC721) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenERC721) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenERC721) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenERC721(val *TokenERC721) *NullableTokenERC721 {
	return &NullableTokenERC721{value: val, isSet: true}
}

func (v NullableTokenERC721) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenERC721) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

