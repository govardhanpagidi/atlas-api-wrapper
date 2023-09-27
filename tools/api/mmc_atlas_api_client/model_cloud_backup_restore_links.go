// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupRestoreLinks struct for CloudBackupRestoreLinks
type CloudBackupRestoreLinks struct {
	Href *string `json:"href,omitempty"`
	Rel  *string `json:"rel,omitempty"`
}

// NewCloudBackupRestoreLinks instantiates a new CloudBackupRestoreLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupRestoreLinks() *CloudBackupRestoreLinks {
	this := CloudBackupRestoreLinks{}
	return &this
}

// NewCloudBackupRestoreLinksWithDefaults instantiates a new CloudBackupRestoreLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupRestoreLinksWithDefaults() *CloudBackupRestoreLinks {
	this := CloudBackupRestoreLinks{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise
func (o *CloudBackupRestoreLinks) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreLinks) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}

	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *CloudBackupRestoreLinks) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *CloudBackupRestoreLinks) SetHref(v string) {
	o.Href = &v
}

// GetRel returns the Rel field value if set, zero value otherwise
func (o *CloudBackupRestoreLinks) GetRel() string {
	if o == nil || IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreLinks) GetRelOk() (*string, bool) {
	if o == nil || IsNil(o.Rel) {
		return nil, false
	}

	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *CloudBackupRestoreLinks) HasRel() bool {
	if o != nil && !IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *CloudBackupRestoreLinks) SetRel(v string) {
	o.Rel = &v
}

func (o CloudBackupRestoreLinks) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupRestoreLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	return toSerialize, nil
}
