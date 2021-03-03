package user

import (
	"datasite/errs"
	"encoding/json"
)

//User declares type for json destructuring
type User struct {
	ID                          string   `json:"id,omitempty"`
	City                        string   `json:"city,omitempty"`
	Company                     string   `json:"company,omitempty"`
	Country                     string   `json:"country,omitempty"`
	FirstName                   string   `json:"firstName,omitempty"`
	LastName                    string   `json:"lastName,omitempty"`
	OrganizationType            string   `json:"organizationType,omitempty"`
	Phone                       string   `json:"phone,omitempty"`
	State                       string   `json:"state,omitempty"`
	ZipCode                     string   `json:"zipCode,omitempty"`
	DisclaimerAccepted          bool     `json:"disclaimerAccepted,omitempty"`
	LanguageCode                string   `json:"languageCode,omitempty"`
	Email                       string   `json:"emailAddress,omitempty"`
	RegistrationID              string   `json:"registrationId,omitempty"`
	RegistrationIDGeneratedTime string   `json:"registrationIdGeneratedTime,omitempty"`
	ProjectIDs                  []string `json:"projectIds"`
}

//ToJSON returns formatted json string of User struct
func (user *User) ToJSON() ([]byte, error) {
	return json.MarshalIndent(user, "", "  ")
}

//Users collection
type Users struct {
	UserList []User
}

//ToJSON returns formatted json string of User struct
func (users *Users) ToJSON() ([]byte, error) {
	return json.MarshalIndent(users.UserList, "", "  ")
}

//UnmarshalUsers process user data points
func (users *Users) UnmarshalUsers(resp []byte) {
	err := json.Unmarshal(resp, &users.UserList)
	errs.HandleError(err)
}
