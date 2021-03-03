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
func TestProcessData(t *testing.T) {

	t.Run("returns a slice of Users with projectIds", func(t *testing.T) {
		user1 := user.User{"1", "Jaydashire", "Goyette - Renner", "South Africa", "firstName 1", "lastName 1", "organizationType 1", "524.276.1570 x487", "SD", "68048", false, "en", "last1@mail.com", "", "", []string{}}
		user2 := user.User{"2", "Adolfofort", "Fisher - Bartoletti", "China", "firstName 2", "lastName 2", "organizationType 2", "(308) 197-9774 x339", "CO", "78569", false, "en", "last2@mail.com", "", "", []string{}}
		users := user.Users{}
		users.UserList = append([]user.User{}, user1, user2)
		projects := []project.Project{{"37", "37", "1"}, {"38", "38", "1"}}
		want := processUserProjects(users, projects)
		if want.UserList[0].ProjectIDs == nil {
			jsonstring, _ := want.UserList[0].ToJSON()
			t.Errorf(string(jsonstring))
		}
	})
}
