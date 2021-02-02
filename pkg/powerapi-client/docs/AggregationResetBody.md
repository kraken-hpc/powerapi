# AggregationResetBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResetType** | Pointer to [**ResetType**](ResetType.md) |  | [optional] 
**TargetURIs** | Pointer to **[]string** | A list of system URIs to apply the reset to | [optional] 

## Methods

### NewAggregationResetBody

`func NewAggregationResetBody() *AggregationResetBody`

NewAggregationResetBody instantiates a new AggregationResetBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregationResetBodyWithDefaults

`func NewAggregationResetBodyWithDefaults() *AggregationResetBody`

NewAggregationResetBodyWithDefaults instantiates a new AggregationResetBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResetType

`func (o *AggregationResetBody) GetResetType() ResetType`

GetResetType returns the ResetType field if non-nil, zero value otherwise.

### GetResetTypeOk

`func (o *AggregationResetBody) GetResetTypeOk() (*ResetType, bool)`

GetResetTypeOk returns a tuple with the ResetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetType

`func (o *AggregationResetBody) SetResetType(v ResetType)`

SetResetType sets ResetType field to given value.

### HasResetType

`func (o *AggregationResetBody) HasResetType() bool`

HasResetType returns a boolean if a field has been set.

### GetTargetURIs

`func (o *AggregationResetBody) GetTargetURIs() []string`

GetTargetURIs returns the TargetURIs field if non-nil, zero value otherwise.

### GetTargetURIsOk

`func (o *AggregationResetBody) GetTargetURIsOk() (*[]string, bool)`

GetTargetURIsOk returns a tuple with the TargetURIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetURIs

`func (o *AggregationResetBody) SetTargetURIs(v []string)`

SetTargetURIs sets TargetURIs field to given value.

### HasTargetURIs

`func (o *AggregationResetBody) HasTargetURIs() bool`

HasTargetURIs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


