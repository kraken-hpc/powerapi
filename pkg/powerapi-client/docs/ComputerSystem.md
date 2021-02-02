# ComputerSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An id is a URI-reference for the object | [readonly] 
**Name** | **string** | The name is the unique name identifier for the ComputerSystem. This is used to reference the system in the API.  | [readonly] 
**PowerState** | Pointer to [**PowerState**](PowerState.md) |  | [optional] 

## Methods

### NewComputerSystem

`func NewComputerSystem(id string, name string, ) *ComputerSystem`

NewComputerSystem instantiates a new ComputerSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerSystemWithDefaults

`func NewComputerSystemWithDefaults() *ComputerSystem`

NewComputerSystemWithDefaults instantiates a new ComputerSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerSystem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerSystem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerSystem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ComputerSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerSystem) SetName(v string)`

SetName sets Name field to given value.


### GetPowerState

`func (o *ComputerSystem) GetPowerState() PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ComputerSystem) GetPowerStateOk() (*PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ComputerSystem) SetPowerState(v PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ComputerSystem) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


