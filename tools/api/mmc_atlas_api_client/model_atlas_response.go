// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// AtlasResponse struct for AtlasResponse
type AtlasResponse struct {
	ErrorCode    *string                `json:"errorCode,omitempty"`
	ErrorMessage *string                `json:"errorMessage,omitempty"`
	Response     map[string]interface{} `json:"response,omitempty"`
	Status       *string                `json:"status,omitempty"`
}

// NewAtlasResponse instantiates a new AtlasResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtlasResponse() *AtlasResponse {
	this := AtlasResponse{}
	return &this
}

// NewAtlasResponseWithDefaults instantiates a new AtlasResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtlasResponseWithDefaults() *AtlasResponse {
	this := AtlasResponse{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise
func (o *AtlasResponse) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}

	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AtlasResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *AtlasResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *AtlasResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AtlasResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AtlasResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetResponse returns the Response field value if set, zero value otherwise
func (o *AtlasResponse) GetResponse() map[string]interface{} {
	if o == nil || IsNil(o.Response) {
		var ret map[string]interface{}
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasResponse) GetResponseOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Response) {
		return map[string]interface{}{}, false
	}

	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AtlasResponse) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given map[string]interface{} and assigns it to the Response field.
func (o *AtlasResponse) SetResponse(v map[string]interface{}) {
	o.Response = v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *AtlasResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AtlasResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AtlasResponse) SetStatus(v string) {
	o.Status = &v
}

func (o AtlasResponse) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AtlasResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}
