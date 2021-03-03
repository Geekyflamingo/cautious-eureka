package user

import (
	"encoding/json"
	"log"
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

//UnmarshalUsers process user data points
func UnmarshalUsers(resp []byte) []User {
	//Convert the body to User type
	var users []User

	err := json.Unmarshal(resp, &users)
	if err != nil {
		log.Fatalln(err)
	}
	return users
}
