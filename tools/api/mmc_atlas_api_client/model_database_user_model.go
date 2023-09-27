// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// DatabaseUserModel struct for DatabaseUserModel
type DatabaseUserModel struct {
	AwsIAMType      *string                       `json:"awsIAMType,omitempty"`
	DatabaseName    *string                       `json:"databaseName,omitempty"`
	DeleteAfterDate *string                       `json:"deleteAfterDate,omitempty"`
	Labels          []DatabaseUserLabelDefinition `json:"labels,omitempty"`
	LdapAuthType    *string                       `json:"ldapAuthType,omitempty"`
	Password        *string                       `json:"password,omitempty"`
	Profile         *string                       `json:"profile,omitempty"`
	ProjectId       *string                       `json:"projectId,omitempty"`
	Roles           []DatabaseUserRoleDefinition  `json:"roles,omitempty"`
	Scopes          []DatabaseUserScopeDefinition `json:"scopes,omitempty"`
	Username        *string                       `json:"username,omitempty"`
	X509Type        *string                       `json:"x509Type,omitempty"`
}

// NewDatabaseUserModel instantiates a new DatabaseUserModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseUserModel() *DatabaseUserModel {
	this := DatabaseUserModel{}
	return &this
}

// NewDatabaseUserModelWithDefaults instantiates a new DatabaseUserModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseUserModelWithDefaults() *DatabaseUserModel {
	this := DatabaseUserModel{}
	return &this
}

// GetAwsIAMType returns the AwsIAMType field value if set, zero value otherwise
func (o *DatabaseUserModel) GetAwsIAMType() string {
	if o == nil || IsNil(o.AwsIAMType) {
		var ret string
		return ret
	}
	return *o.AwsIAMType
}

// GetAwsIAMTypeOk returns a tuple with the AwsIAMType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetAwsIAMTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AwsIAMType) {
		return nil, false
	}

	return o.AwsIAMType, true
}

// HasAwsIAMType returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasAwsIAMType() bool {
	if o != nil && !IsNil(o.AwsIAMType) {
		return true
	}

	return false
}

// SetAwsIAMType gets a reference to the given string and assigns it to the AwsIAMType field.
func (o *DatabaseUserModel) SetAwsIAMType(v string) {
	o.AwsIAMType = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise
func (o *DatabaseUserModel) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}

	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *DatabaseUserModel) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetDeleteAfterDate returns the DeleteAfterDate field value if set, zero value otherwise
func (o *DatabaseUserModel) GetDeleteAfterDate() string {
	if o == nil || IsNil(o.DeleteAfterDate) {
		var ret string
		return ret
	}
	return *o.DeleteAfterDate
}

// GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetDeleteAfterDateOk() (*string, bool) {
	if o == nil || IsNil(o.DeleteAfterDate) {
		return nil, false
	}

	return o.DeleteAfterDate, true
}

// HasDeleteAfterDate returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasDeleteAfterDate() bool {
	if o != nil && !IsNil(o.DeleteAfterDate) {
		return true
	}

	return false
}

// SetDeleteAfterDate gets a reference to the given string and assigns it to the DeleteAfterDate field.
func (o *DatabaseUserModel) SetDeleteAfterDate(v string) {
	o.DeleteAfterDate = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise
func (o *DatabaseUserModel) GetLabels() []DatabaseUserLabelDefinition {
	if o == nil || IsNil(o.Labels) {
		var ret []DatabaseUserLabelDefinition
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetLabelsOk() ([]DatabaseUserLabelDefinition, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}

	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []DatabaseUserLabelDefinition and assigns it to the Labels field.
func (o *DatabaseUserModel) SetLabels(v []DatabaseUserLabelDefinition) {
	o.Labels = v
}

// GetLdapAuthType returns the LdapAuthType field value if set, zero value otherwise
func (o *DatabaseUserModel) GetLdapAuthType() string {
	if o == nil || IsNil(o.LdapAuthType) {
		var ret string
		return ret
	}
	return *o.LdapAuthType
}

// GetLdapAuthTypeOk returns a tuple with the LdapAuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetLdapAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LdapAuthType) {
		return nil, false
	}

	return o.LdapAuthType, true
}

// HasLdapAuthType returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasLdapAuthType() bool {
	if o != nil && !IsNil(o.LdapAuthType) {
		return true
	}

	return false
}

// SetLdapAuthType gets a reference to the given string and assigns it to the LdapAuthType field.
func (o *DatabaseUserModel) SetLdapAuthType(v string) {
	o.LdapAuthType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise
func (o *DatabaseUserModel) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}

	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DatabaseUserModel) SetPassword(v string) {
	o.Password = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise
func (o *DatabaseUserModel) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}

	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *DatabaseUserModel) SetProfile(v string) {
	o.Profile = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *DatabaseUserModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DatabaseUserModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *DatabaseUserModel) GetRoles() []DatabaseUserRoleDefinition {
	if o == nil || IsNil(o.Roles) {
		var ret []DatabaseUserRoleDefinition
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetRolesOk() ([]DatabaseUserRoleDefinition, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []DatabaseUserRoleDefinition and assigns it to the Roles field.
func (o *DatabaseUserModel) SetRoles(v []DatabaseUserRoleDefinition) {
	o.Roles = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise
func (o *DatabaseUserModel) GetScopes() []DatabaseUserScopeDefinition {
	if o == nil || IsNil(o.Scopes) {
		var ret []DatabaseUserScopeDefinition
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetScopesOk() ([]DatabaseUserScopeDefinition, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}

	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []DatabaseUserScopeDefinition and assigns it to the Scopes field.
func (o *DatabaseUserModel) SetScopes(v []DatabaseUserScopeDefinition) {
	o.Scopes = v
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *DatabaseUserModel) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DatabaseUserModel) SetUsername(v string) {
	o.Username = &v
}

// GetX509Type returns the X509Type field value if set, zero value otherwise
func (o *DatabaseUserModel) GetX509Type() string {
	if o == nil || IsNil(o.X509Type) {
		var ret string
		return ret
	}
	return *o.X509Type
}

// GetX509TypeOk returns a tuple with the X509Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserModel) GetX509TypeOk() (*string, bool) {
	if o == nil || IsNil(o.X509Type) {
		return nil, false
	}

	return o.X509Type, true
}

// HasX509Type returns a boolean if a field has been set.
func (o *DatabaseUserModel) HasX509Type() bool {
	if o != nil && !IsNil(o.X509Type) {
		return true
	}

	return false
}

// SetX509Type gets a reference to the given string and assigns it to the X509Type field.
func (o *DatabaseUserModel) SetX509Type(v string) {
	o.X509Type = &v
}

func (o DatabaseUserModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabaseUserModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsIAMType) {
		toSerialize["awsIAMType"] = o.AwsIAMType
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !IsNil(o.DeleteAfterDate) {
		toSerialize["deleteAfterDate"] = o.DeleteAfterDate
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.LdapAuthType) {
		toSerialize["ldapAuthType"] = o.LdapAuthType
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.X509Type) {
		toSerialize["x509Type"] = o.X509Type
	}
	return toSerialize, nil
}
