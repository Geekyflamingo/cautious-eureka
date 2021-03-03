package userprojects

import (
	"datasite/project"
	"datasite/user"
	"testing"
)

func TestGetData(t *testing.T) {

	t.Run("returns registered user 1", func(t *testing.T) {

		got := getData("https://5c3ce12c29429300143fe570.mockapi.io/api/registeredusers/1")
		want := string([]byte(`{"id":"1","city":"Jaydashire","company":"Goyette - Renner","country":"South Africa","firstName":"firstName 1","lastName":"lastName 1","organizationType":"organizationType 1","phone":"524.276.1570 x487","state":"SD","zipCode":"68048","disclaimerAccepted":false,"languageCode":"en","emailAddress":"last1@mail.com"}`))

		if string(got) != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
func TestProcessUserProjects(t *testing.T) {

	t.Run("returns a slice of Users with projectIds", func(t *testing.T) {
		user1 := user.User{
			ID:                          "1",
			City:                        "Jaydashire",
			Company:                     "Goyette - Renner",
			Country:                     "South Africa",
			FirstName:                   "firstName 1",
			LastName:                    "lastName 1",
			OrganizationType:            "organizationType 1",
			Phone:                       "524.276.1570 x487",
			State:                       "SD",
			ZipCode:                     "68048",
			DisclaimerAccepted:          false,
			LanguageCode:                "en",
			Email:                       "last1@mail.com",
			RegistrationID:              "",
			RegistrationIDGeneratedTime: "",
			ProjectIDs:                  []string{},
		}
		user2 := user.User{
			ID:                          "2",
			City:                        "Adolfofort",
			Company:                     "Fisher - Bartoletti",
			Country:                     "China",
			FirstName:                   "firstName 2",
			LastName:                    "lastName 2",
			OrganizationType:            "organizationType 2",
			Phone:                       "(308) 197-9774 x339",
			State:                       "CO",
			ZipCode:                     "78569",
			DisclaimerAccepted:          false,
			LanguageCode:                "en",
			Email:                       "last2@mail.com",
			RegistrationID:              "",
			RegistrationIDGeneratedTime: "",
			ProjectIDs:                  []string{},
		}

		users := user.Users{}
		users.UserList = append([]user.User{}, user1, user2)
		projects := []project.Project{{ID: "37", ProjectID: "37", UserID: "1"},
			{ID: "38", ProjectID: "38", UserID: "1"}}
		want := processUserProjects(users, projects)
		if want.UserList[0].ProjectIDs == nil {
			jsonstring, _ := want.UserList[0].ToJSON()
			t.Errorf(string(jsonstring))
		}
	})
}
