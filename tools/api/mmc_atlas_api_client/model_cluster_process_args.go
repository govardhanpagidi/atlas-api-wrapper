// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterProcessArgs struct for ClusterProcessArgs
type ClusterProcessArgs struct {
	DefaultReadConcern               *string  `json:"defaultReadConcern,omitempty"`
	DefaultWriteConcern              *string  `json:"defaultWriteConcern,omitempty"`
	FailIndexKeyTooLong              *bool    `json:"failIndexKeyTooLong,omitempty"`
	JavascriptEnabled                *bool    `json:"javascriptEnabled,omitempty"`
	MinimumEnabledTLSProtocol        *string  `json:"minimumEnabledTLSProtocol,omitempty"`
	NoTableScan                      *bool    `json:"noTableScan,omitempty"`
	OplogMinRetentionHours           *float32 `json:"oplogMinRetentionHours,omitempty"`
	OplogSizeMB                      *int     `json:"oplogSizeMB,omitempty"`
	SampleRefreshIntervalBIConnector *int     `json:"sampleRefreshIntervalBIConnector,omitempty"`
	SampleSizeBIConnector            *int     `json:"sampleSizeBIConnector,omitempty"`
}

// NewClusterProcessArgs instantiates a new ClusterProcessArgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterProcessArgs() *ClusterProcessArgs {
	this := ClusterProcessArgs{}
	return &this
}

// NewClusterProcessArgsWithDefaults instantiates a new ClusterProcessArgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterProcessArgsWithDefaults() *ClusterProcessArgs {
	this := ClusterProcessArgs{}
	return &this
}

// GetDefaultReadConcern returns the DefaultReadConcern field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetDefaultReadConcern() string {
	if o == nil || IsNil(o.DefaultReadConcern) {
		var ret string
		return ret
	}
	return *o.DefaultReadConcern
}

// GetDefaultReadConcernOk returns a tuple with the DefaultReadConcern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetDefaultReadConcernOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultReadConcern) {
		return nil, false
	}

	return o.DefaultReadConcern, true
}

// HasDefaultReadConcern returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasDefaultReadConcern() bool {
	if o != nil && !IsNil(o.DefaultReadConcern) {
		return true
	}

	return false
}

// SetDefaultReadConcern gets a reference to the given string and assigns it to the DefaultReadConcern field.
func (o *ClusterProcessArgs) SetDefaultReadConcern(v string) {
	o.DefaultReadConcern = &v
}

// GetDefaultWriteConcern returns the DefaultWriteConcern field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetDefaultWriteConcern() string {
	if o == nil || IsNil(o.DefaultWriteConcern) {
		var ret string
		return ret
	}
	return *o.DefaultWriteConcern
}

// GetDefaultWriteConcernOk returns a tuple with the DefaultWriteConcern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetDefaultWriteConcernOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultWriteConcern) {
		return nil, false
	}

	return o.DefaultWriteConcern, true
}

// HasDefaultWriteConcern returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasDefaultWriteConcern() bool {
	if o != nil && !IsNil(o.DefaultWriteConcern) {
		return true
	}

	return false
}

// SetDefaultWriteConcern gets a reference to the given string and assigns it to the DefaultWriteConcern field.
func (o *ClusterProcessArgs) SetDefaultWriteConcern(v string) {
	o.DefaultWriteConcern = &v
}

// GetFailIndexKeyTooLong returns the FailIndexKeyTooLong field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetFailIndexKeyTooLong() bool {
	if o == nil || IsNil(o.FailIndexKeyTooLong) {
		var ret bool
		return ret
	}
	return *o.FailIndexKeyTooLong
}

// GetFailIndexKeyTooLongOk returns a tuple with the FailIndexKeyTooLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetFailIndexKeyTooLongOk() (*bool, bool) {
	if o == nil || IsNil(o.FailIndexKeyTooLong) {
		return nil, false
	}

	return o.FailIndexKeyTooLong, true
}

// HasFailIndexKeyTooLong returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasFailIndexKeyTooLong() bool {
	if o != nil && !IsNil(o.FailIndexKeyTooLong) {
		return true
	}

	return false
}

// SetFailIndexKeyTooLong gets a reference to the given bool and assigns it to the FailIndexKeyTooLong field.
func (o *ClusterProcessArgs) SetFailIndexKeyTooLong(v bool) {
	o.FailIndexKeyTooLong = &v
}

// GetJavascriptEnabled returns the JavascriptEnabled field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetJavascriptEnabled() bool {
	if o == nil || IsNil(o.JavascriptEnabled) {
		var ret bool
		return ret
	}
	return *o.JavascriptEnabled
}

// GetJavascriptEnabledOk returns a tuple with the JavascriptEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetJavascriptEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.JavascriptEnabled) {
		return nil, false
	}

	return o.JavascriptEnabled, true
}

// HasJavascriptEnabled returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasJavascriptEnabled() bool {
	if o != nil && !IsNil(o.JavascriptEnabled) {
		return true
	}

	return false
}

// SetJavascriptEnabled gets a reference to the given bool and assigns it to the JavascriptEnabled field.
func (o *ClusterProcessArgs) SetJavascriptEnabled(v bool) {
	o.JavascriptEnabled = &v
}

// GetMinimumEnabledTLSProtocol returns the MinimumEnabledTLSProtocol field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetMinimumEnabledTLSProtocol() string {
	if o == nil || IsNil(o.MinimumEnabledTLSProtocol) {
		var ret string
		return ret
	}
	return *o.MinimumEnabledTLSProtocol
}

// GetMinimumEnabledTLSProtocolOk returns a tuple with the MinimumEnabledTLSProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetMinimumEnabledTLSProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumEnabledTLSProtocol) {
		return nil, false
	}

	return o.MinimumEnabledTLSProtocol, true
}

// HasMinimumEnabledTLSProtocol returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasMinimumEnabledTLSProtocol() bool {
	if o != nil && !IsNil(o.MinimumEnabledTLSProtocol) {
		return true
	}

	return false
}

// SetMinimumEnabledTLSProtocol gets a reference to the given string and assigns it to the MinimumEnabledTLSProtocol field.
func (o *ClusterProcessArgs) SetMinimumEnabledTLSProtocol(v string) {
	o.MinimumEnabledTLSProtocol = &v
}

// GetNoTableScan returns the NoTableScan field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetNoTableScan() bool {
	if o == nil || IsNil(o.NoTableScan) {
		var ret bool
		return ret
	}
	return *o.NoTableScan
}

// GetNoTableScanOk returns a tuple with the NoTableScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetNoTableScanOk() (*bool, bool) {
	if o == nil || IsNil(o.NoTableScan) {
		return nil, false
	}

	return o.NoTableScan, true
}

// HasNoTableScan returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasNoTableScan() bool {
	if o != nil && !IsNil(o.NoTableScan) {
		return true
	}

	return false
}

// SetNoTableScan gets a reference to the given bool and assigns it to the NoTableScan field.
func (o *ClusterProcessArgs) SetNoTableScan(v bool) {
	o.NoTableScan = &v
}

// GetOplogMinRetentionHours returns the OplogMinRetentionHours field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetOplogMinRetentionHours() float32 {
	if o == nil || IsNil(o.OplogMinRetentionHours) {
		var ret float32
		return ret
	}
	return *o.OplogMinRetentionHours
}

// GetOplogMinRetentionHoursOk returns a tuple with the OplogMinRetentionHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetOplogMinRetentionHoursOk() (*float32, bool) {
	if o == nil || IsNil(o.OplogMinRetentionHours) {
		return nil, false
	}

	return o.OplogMinRetentionHours, true
}

// HasOplogMinRetentionHours returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasOplogMinRetentionHours() bool {
	if o != nil && !IsNil(o.OplogMinRetentionHours) {
		return true
	}

	return false
}

// SetOplogMinRetentionHours gets a reference to the given float32 and assigns it to the OplogMinRetentionHours field.
func (o *ClusterProcessArgs) SetOplogMinRetentionHours(v float32) {
	o.OplogMinRetentionHours = &v
}

// GetOplogSizeMB returns the OplogSizeMB field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetOplogSizeMB() int {
	if o == nil || IsNil(o.OplogSizeMB) {
		var ret int
		return ret
	}
	return *o.OplogSizeMB
}

// GetOplogSizeMBOk returns a tuple with the OplogSizeMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetOplogSizeMBOk() (*int, bool) {
	if o == nil || IsNil(o.OplogSizeMB) {
		return nil, false
	}

	return o.OplogSizeMB, true
}

// HasOplogSizeMB returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasOplogSizeMB() bool {
	if o != nil && !IsNil(o.OplogSizeMB) {
		return true
	}

	return false
}

// SetOplogSizeMB gets a reference to the given int and assigns it to the OplogSizeMB field.
func (o *ClusterProcessArgs) SetOplogSizeMB(v int) {
	o.OplogSizeMB = &v
}

// GetSampleRefreshIntervalBIConnector returns the SampleRefreshIntervalBIConnector field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetSampleRefreshIntervalBIConnector() int {
	if o == nil || IsNil(o.SampleRefreshIntervalBIConnector) {
		var ret int
		return ret
	}
	return *o.SampleRefreshIntervalBIConnector
}

// GetSampleRefreshIntervalBIConnectorOk returns a tuple with the SampleRefreshIntervalBIConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetSampleRefreshIntervalBIConnectorOk() (*int, bool) {
	if o == nil || IsNil(o.SampleRefreshIntervalBIConnector) {
		return nil, false
	}

	return o.SampleRefreshIntervalBIConnector, true
}

// HasSampleRefreshIntervalBIConnector returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasSampleRefreshIntervalBIConnector() bool {
	if o != nil && !IsNil(o.SampleRefreshIntervalBIConnector) {
		return true
	}

	return false
}

// SetSampleRefreshIntervalBIConnector gets a reference to the given int and assigns it to the SampleRefreshIntervalBIConnector field.
func (o *ClusterProcessArgs) SetSampleRefreshIntervalBIConnector(v int) {
	o.SampleRefreshIntervalBIConnector = &v
}

// GetSampleSizeBIConnector returns the SampleSizeBIConnector field value if set, zero value otherwise
func (o *ClusterProcessArgs) GetSampleSizeBIConnector() int {
	if o == nil || IsNil(o.SampleSizeBIConnector) {
		var ret int
		return ret
	}
	return *o.SampleSizeBIConnector
}

// GetSampleSizeBIConnectorOk returns a tuple with the SampleSizeBIConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterProcessArgs) GetSampleSizeBIConnectorOk() (*int, bool) {
	if o == nil || IsNil(o.SampleSizeBIConnector) {
		return nil, false
	}

	return o.SampleSizeBIConnector, true
}

// HasSampleSizeBIConnector returns a boolean if a field has been set.
func (o *ClusterProcessArgs) HasSampleSizeBIConnector() bool {
	if o != nil && !IsNil(o.SampleSizeBIConnector) {
		return true
	}

	return false
}

// SetSampleSizeBIConnector gets a reference to the given int and assigns it to the SampleSizeBIConnector field.
func (o *ClusterProcessArgs) SetSampleSizeBIConnector(v int) {
	o.SampleSizeBIConnector = &v
}

func (o ClusterProcessArgs) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterProcessArgs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultReadConcern) {
		toSerialize["defaultReadConcern"] = o.DefaultReadConcern
	}
	if !IsNil(o.DefaultWriteConcern) {
		toSerialize["defaultWriteConcern"] = o.DefaultWriteConcern
	}
	if !IsNil(o.FailIndexKeyTooLong) {
		toSerialize["failIndexKeyTooLong"] = o.FailIndexKeyTooLong
	}
	if !IsNil(o.JavascriptEnabled) {
		toSerialize["javascriptEnabled"] = o.JavascriptEnabled
	}
	if !IsNil(o.MinimumEnabledTLSProtocol) {
		toSerialize["minimumEnabledTLSProtocol"] = o.MinimumEnabledTLSProtocol
	}
	if !IsNil(o.NoTableScan) {
		toSerialize["noTableScan"] = o.NoTableScan
	}
	if !IsNil(o.OplogMinRetentionHours) {
		toSerialize["oplogMinRetentionHours"] = o.OplogMinRetentionHours
	}
	if !IsNil(o.OplogSizeMB) {
		toSerialize["oplogSizeMB"] = o.OplogSizeMB
	}
	if !IsNil(o.SampleRefreshIntervalBIConnector) {
		toSerialize["sampleRefreshIntervalBIConnector"] = o.SampleRefreshIntervalBIConnector
	}
	if !IsNil(o.SampleSizeBIConnector) {
		toSerialize["sampleSizeBIConnector"] = o.SampleSizeBIConnector
	}
	return toSerialize, nil
}
