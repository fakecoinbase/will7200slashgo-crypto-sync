/*
 * CoinBase API
 *
 * Coinbase provides a simple and powerful REST API to integrate bitcoin, bitcoin cash, litecoin and ethereum payments into your business or application.  This API reference provides information on available endpoints and how to interact with it. To read more about the API, visit our API documentation.
 *
 * API version: 2019-11-15
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package coinbase

import (
	"encoding/json"
)

// User struct for User
type User struct {
	Id              *string `json:"id,omitempty"`
	Name            *string `json:"name,omitempty"`
	Username        *string `json:"username,omitempty"`
	ProfileLocation *string `json:"profile_location,omitempty"`
	ProfileBio      *string `json:"profile_bio,omitempty"`
	ProfileUrl      *string `json:"profile_url,omitempty"`
	AvatarUrl       *string `json:"avatar_url,omitempty"`
	Resource        *string `json:"resource,omitempty"`
	ResourcePath    *string `json:"resource_path,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetProfileLocation returns the ProfileLocation field value if set, zero value otherwise.
func (o *User) GetProfileLocation() string {
	if o == nil || o.ProfileLocation == nil {
		var ret string
		return ret
	}
	return *o.ProfileLocation
}

// GetProfileLocationOk returns a tuple with the ProfileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProfileLocationOk() (*string, bool) {
	if o == nil || o.ProfileLocation == nil {
		return nil, false
	}
	return o.ProfileLocation, true
}

// HasProfileLocation returns a boolean if a field has been set.
func (o *User) HasProfileLocation() bool {
	if o != nil && o.ProfileLocation != nil {
		return true
	}

	return false
}

// SetProfileLocation gets a reference to the given string and assigns it to the ProfileLocation field.
func (o *User) SetProfileLocation(v string) {
	o.ProfileLocation = &v
}

// GetProfileBio returns the ProfileBio field value if set, zero value otherwise.
func (o *User) GetProfileBio() string {
	if o == nil || o.ProfileBio == nil {
		var ret string
		return ret
	}
	return *o.ProfileBio
}

// GetProfileBioOk returns a tuple with the ProfileBio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProfileBioOk() (*string, bool) {
	if o == nil || o.ProfileBio == nil {
		return nil, false
	}
	return o.ProfileBio, true
}

// HasProfileBio returns a boolean if a field has been set.
func (o *User) HasProfileBio() bool {
	if o != nil && o.ProfileBio != nil {
		return true
	}

	return false
}

// SetProfileBio gets a reference to the given string and assigns it to the ProfileBio field.
func (o *User) SetProfileBio(v string) {
	o.ProfileBio = &v
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise.
func (o *User) GetProfileUrl() string {
	if o == nil || o.ProfileUrl == nil {
		var ret string
		return ret
	}
	return *o.ProfileUrl
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProfileUrlOk() (*string, bool) {
	if o == nil || o.ProfileUrl == nil {
		return nil, false
	}
	return o.ProfileUrl, true
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *User) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl != nil {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given string and assigns it to the ProfileUrl field.
func (o *User) SetProfileUrl(v string) {
	o.ProfileUrl = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *User) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *User) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *User) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *User) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *User) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *User) SetResource(v string) {
	o.Resource = &v
}

// GetResourcePath returns the ResourcePath field value if set, zero value otherwise.
func (o *User) GetResourcePath() string {
	if o == nil || o.ResourcePath == nil {
		var ret string
		return ret
	}
	return *o.ResourcePath
}

// GetResourcePathOk returns a tuple with the ResourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetResourcePathOk() (*string, bool) {
	if o == nil || o.ResourcePath == nil {
		return nil, false
	}
	return o.ResourcePath, true
}

// HasResourcePath returns a boolean if a field has been set.
func (o *User) HasResourcePath() bool {
	if o != nil && o.ResourcePath != nil {
		return true
	}

	return false
}

// SetResourcePath gets a reference to the given string and assigns it to the ResourcePath field.
func (o *User) SetResourcePath(v string) {
	o.ResourcePath = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.ProfileLocation != nil {
		toSerialize["profile_location"] = o.ProfileLocation
	}
	if o.ProfileBio != nil {
		toSerialize["profile_bio"] = o.ProfileBio
	}
	if o.ProfileUrl != nil {
		toSerialize["profile_url"] = o.ProfileUrl
	}
	if o.AvatarUrl != nil {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.ResourcePath != nil {
		toSerialize["resource_path"] = o.ResourcePath
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}