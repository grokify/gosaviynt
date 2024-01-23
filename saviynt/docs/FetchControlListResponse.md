# FetchControlListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | Pointer to **string** |  | [optional] 
**Controls** | Pointer to [**[]FetchControlListControl**](FetchControlListControl.md) |  | [optional] 
**DisplayCount** | Pointer to **int32** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewFetchControlListResponse

`func NewFetchControlListResponse() *FetchControlListResponse`

NewFetchControlListResponse instantiates a new FetchControlListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchControlListResponseWithDefaults

`func NewFetchControlListResponseWithDefaults() *FetchControlListResponse`

NewFetchControlListResponseWithDefaults instantiates a new FetchControlListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *FetchControlListResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *FetchControlListResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *FetchControlListResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *FetchControlListResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetControls

`func (o *FetchControlListResponse) GetControls() []FetchControlListControl`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *FetchControlListResponse) GetControlsOk() (*[]FetchControlListControl, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *FetchControlListResponse) SetControls(v []FetchControlListControl)`

SetControls sets Controls field to given value.

### HasControls

`func (o *FetchControlListResponse) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetDisplayCount

`func (o *FetchControlListResponse) GetDisplayCount() int32`

GetDisplayCount returns the DisplayCount field if non-nil, zero value otherwise.

### GetDisplayCountOk

`func (o *FetchControlListResponse) GetDisplayCountOk() (*int32, bool)`

GetDisplayCountOk returns a tuple with the DisplayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCount

`func (o *FetchControlListResponse) SetDisplayCount(v int32)`

SetDisplayCount sets DisplayCount field to given value.

### HasDisplayCount

`func (o *FetchControlListResponse) HasDisplayCount() bool`

HasDisplayCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *FetchControlListResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *FetchControlListResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *FetchControlListResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *FetchControlListResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetTotalCount

`func (o *FetchControlListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *FetchControlListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *FetchControlListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *FetchControlListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


