// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupRestoreSynchronousCreationOptions struct for CloudBackupRestoreSynchronousCreationOptions
type CloudBackupRestoreSynchronousCreationOptions struct {
	CallbackDelaySeconds   *int  `json:"callbackDelaySeconds,omitempty"`
	ReturnSuccessIfTimeOut *bool `json:"returnSuccessIfTimeOut,omitempty"`
	TimeOutInSeconds       *int  `json:"timeOutInSeconds,omitempty"`
}

// NewCloudBackupRestoreSynchronousCreationOptions instantiates a new CloudBackupRestoreSynchronousCreationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupRestoreSynchronousCreationOptions() *CloudBackupRestoreSynchronousCreationOptions {
	this := CloudBackupRestoreSynchronousCreationOptions{}
	return &this
}

// NewCloudBackupRestoreSynchronousCreationOptionsWithDefaults instantiates a new CloudBackupRestoreSynchronousCreationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupRestoreSynchronousCreationOptionsWithDefaults() *CloudBackupRestoreSynchronousCreationOptions {
	this := CloudBackupRestoreSynchronousCreationOptions{}
	return &this
}

// GetCallbackDelaySeconds returns the CallbackDelaySeconds field value if set, zero value otherwise
func (o *CloudBackupRestoreSynchronousCreationOptions) GetCallbackDelaySeconds() int {
	if o == nil || IsNil(o.CallbackDelaySeconds) {
		var ret int
		return ret
	}
	return *o.CallbackDelaySeconds
}

// GetCallbackDelaySecondsOk returns a tuple with the CallbackDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreSynchronousCreationOptions) GetCallbackDelaySecondsOk() (*int, bool) {
	if o == nil || IsNil(o.CallbackDelaySeconds) {
		return nil, false
	}

	return o.CallbackDelaySeconds, true
}

// HasCallbackDelaySeconds returns a boolean if a field has been set.
func (o *CloudBackupRestoreSynchronousCreationOptions) HasCallbackDelaySeconds() bool {
	if o != nil && !IsNil(o.CallbackDelaySeconds) {
		return true
	}

	return false
}

// SetCallbackDelaySeconds gets a reference to the given int and assigns it to the CallbackDelaySeconds field.
func (o *CloudBackupRestoreSynchronousCreationOptions) SetCallbackDelaySeconds(v int) {
	o.CallbackDelaySeconds = &v
}

// GetReturnSuccessIfTimeOut returns the ReturnSuccessIfTimeOut field value if set, zero value otherwise
func (o *CloudBackupRestoreSynchronousCreationOptions) GetReturnSuccessIfTimeOut() bool {
	if o == nil || IsNil(o.ReturnSuccessIfTimeOut) {
		var ret bool
		return ret
	}
	return *o.ReturnSuccessIfTimeOut
}

// GetReturnSuccessIfTimeOutOk returns a tuple with the ReturnSuccessIfTimeOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreSynchronousCreationOptions) GetReturnSuccessIfTimeOutOk() (*bool, bool) {
	if o == nil || IsNil(o.ReturnSuccessIfTimeOut) {
		return nil, false
	}

	return o.ReturnSuccessIfTimeOut, true
}

// HasReturnSuccessIfTimeOut returns a boolean if a field has been set.
func (o *CloudBackupRestoreSynchronousCreationOptions) HasReturnSuccessIfTimeOut() bool {
	if o != nil && !IsNil(o.ReturnSuccessIfTimeOut) {
		return true
	}

	return false
}

// SetReturnSuccessIfTimeOut gets a reference to the given bool and assigns it to the ReturnSuccessIfTimeOut field.
func (o *CloudBackupRestoreSynchronousCreationOptions) SetReturnSuccessIfTimeOut(v bool) {
	o.ReturnSuccessIfTimeOut = &v
}

// GetTimeOutInSeconds returns the TimeOutInSeconds field value if set, zero value otherwise
func (o *CloudBackupRestoreSynchronousCreationOptions) GetTimeOutInSeconds() int {
	if o == nil || IsNil(o.TimeOutInSeconds) {
		var ret int
		return ret
	}
	return *o.TimeOutInSeconds
}

// GetTimeOutInSecondsOk returns a tuple with the TimeOutInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreSynchronousCreationOptions) GetTimeOutInSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.TimeOutInSeconds) {
		return nil, false
	}

	return o.TimeOutInSeconds, true
}

// HasTimeOutInSeconds returns a boolean if a field has been set.
func (o *CloudBackupRestoreSynchronousCreationOptions) HasTimeOutInSeconds() bool {
	if o != nil && !IsNil(o.TimeOutInSeconds) {
		return true
	}

	return false
}

// SetTimeOutInSeconds gets a reference to the given int and assigns it to the TimeOutInSeconds field.
func (o *CloudBackupRestoreSynchronousCreationOptions) SetTimeOutInSeconds(v int) {
	o.TimeOutInSeconds = &v
}

func (o CloudBackupRestoreSynchronousCreationOptions) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupRestoreSynchronousCreationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallbackDelaySeconds) {
		toSerialize["callbackDelaySeconds"] = o.CallbackDelaySeconds
	}
	if !IsNil(o.ReturnSuccessIfTimeOut) {
		toSerialize["returnSuccessIfTimeOut"] = o.ReturnSuccessIfTimeOut
	}
	if !IsNil(o.TimeOutInSeconds) {
		toSerialize["timeOutInSeconds"] = o.TimeOutInSeconds
	}
	return toSerialize, nil
}
