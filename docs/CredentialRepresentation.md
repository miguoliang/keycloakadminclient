# CredentialRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserLabel** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **int64** |  | [optional] 
**SecretData** | Pointer to **string** |  | [optional] 
**CredentialData** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 

## Methods

### NewCredentialRepresentation

`func NewCredentialRepresentation() *CredentialRepresentation`

NewCredentialRepresentation instantiates a new CredentialRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialRepresentationWithDefaults

`func NewCredentialRepresentationWithDefaults() *CredentialRepresentation`

NewCredentialRepresentationWithDefaults instantiates a new CredentialRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CredentialRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CredentialRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserLabel

`func (o *CredentialRepresentation) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *CredentialRepresentation) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *CredentialRepresentation) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *CredentialRepresentation) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *CredentialRepresentation) GetCreatedDate() int64`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CredentialRepresentation) GetCreatedDateOk() (*int64, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CredentialRepresentation) SetCreatedDate(v int64)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CredentialRepresentation) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSecretData

`func (o *CredentialRepresentation) GetSecretData() string`

GetSecretData returns the SecretData field if non-nil, zero value otherwise.

### GetSecretDataOk

`func (o *CredentialRepresentation) GetSecretDataOk() (*string, bool)`

GetSecretDataOk returns a tuple with the SecretData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretData

`func (o *CredentialRepresentation) SetSecretData(v string)`

SetSecretData sets SecretData field to given value.

### HasSecretData

`func (o *CredentialRepresentation) HasSecretData() bool`

HasSecretData returns a boolean if a field has been set.

### GetCredentialData

`func (o *CredentialRepresentation) GetCredentialData() string`

GetCredentialData returns the CredentialData field if non-nil, zero value otherwise.

### GetCredentialDataOk

`func (o *CredentialRepresentation) GetCredentialDataOk() (*string, bool)`

GetCredentialDataOk returns a tuple with the CredentialData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialData

`func (o *CredentialRepresentation) SetCredentialData(v string)`

SetCredentialData sets CredentialData field to given value.

### HasCredentialData

`func (o *CredentialRepresentation) HasCredentialData() bool`

HasCredentialData returns a boolean if a field has been set.

### GetPriority

`func (o *CredentialRepresentation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CredentialRepresentation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CredentialRepresentation) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CredentialRepresentation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetValue

`func (o *CredentialRepresentation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CredentialRepresentation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CredentialRepresentation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CredentialRepresentation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTemporary

`func (o *CredentialRepresentation) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *CredentialRepresentation) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *CredentialRepresentation) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *CredentialRepresentation) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


