# ComputerSystemCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An id is a URI-reference for the object | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable name for the collection | [optional] [readonly] 
**Systems** | Pointer to [**[]ComputerSystem**](ComputerSystem.md) | Collection of ComputerSystem objects | [optional] 

## Methods

### NewComputerSystemCollection

`func NewComputerSystemCollection() *ComputerSystemCollection`

NewComputerSystemCollection instantiates a new ComputerSystemCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputerSystemCollectionWithDefaults

`func NewComputerSystemCollectionWithDefaults() *ComputerSystemCollection`

NewComputerSystemCollectionWithDefaults instantiates a new ComputerSystemCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputerSystemCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputerSystemCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputerSystemCollection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputerSystemCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ComputerSystemCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputerSystemCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputerSystemCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputerSystemCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSystems

`func (o *ComputerSystemCollection) GetSystems() []ComputerSystem`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *ComputerSystemCollection) GetSystemsOk() (*[]ComputerSystem, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *ComputerSystemCollection) SetSystems(v []ComputerSystem)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *ComputerSystemCollection) HasSystems() bool`

HasSystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


