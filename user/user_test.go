package user

import (
	"testing"

	"github.com/go-test/deep"
)

var unregdata = []byte(`[{"id":"214","emailAddress":"email14@somewhere.com","languageCode":"en","registrationId":"jwsMJNOk3oM3hVM5bGcF14","registrationIdGeneratedTime":"1561650268514"},{"id":"215","emailAddress":"email15@somewhere.com","languageCode":"en","registrationId":"jwsMJNOk3oM3hVM5bGcF15","registrationIdGeneratedTime":"1561650268515"}]`)

var regdata = []byte(`[{"id":"1","city":"Jaydashire","company":"Goyette - Renner","country":"South Africa","firstName":"firstName 1","lastName":"lastName 1","organizationType":"organizationType 1","phone":"524.276.1570 x487","state":"SD","zipCode":"68048","disclaimerAccepted":false,"languageCode":"en","emailAddress":"last1@mail.com"},{"id":"2","city":"Adolfofort","company":"Fisher - Bartoletti","country":"China","firstName":"firstName 2","lastName":"lastName 2","organizationType":"organizationType 2","phone":"(308) 197-9774 x339","state":"CO","zipCode":"78569","disclaimerAccepted":false,"languageCode":"en","emailAddress":"last2@mail.com"}]`)

func TestUnmarshalUsers(t *testing.T) {

	assertCorrectUsersStruct := func(t testing.TB, got, want Users) {
		deep.NilSlicesAreEmpty = true
		t.Helper()
		match := false
		for _, gu := range got.UserList {
			for _, wu := range want.UserList {
				if diff := deep.Equal(wu, gu); diff == nil {
					match = true
				}
			}
			if !match {
				t.Errorf("got %v want %v", got, want)
			}
		}

	}

	t.Run("Registered Users", func(t *testing.T) {
		user1 := User{"1", "Jaydashire", "Goyette - Renner", "South Africa", "firstName 1", "lastName 1", "organizationType 1", "524.276.1570 x487", "SD", "68048", false, "en", "last1@mail.com", "", "", []string{}}
		user2 := User{"2", "Adolfofort", "Fisher - Bartoletti", "China", "firstName 2", "lastName 2", "organizationType 2", "(308) 197-9774 x339", "CO", "78569", false, "en", "last2@mail.com", "", "", []string{}}
		users := Users{}
		want := Users{}

		users.UnmarshalUsers(regdata)
		want.UserList = append(Users{}.UserList, user1, user2)

		assertCorrectUsersStruct(t, users, want)

	})

	t.Run("Unregistered Users", func(t *testing.T) {
		user1 := User{"214", "", "", "", "", "", "", "", "", "", false, "en", "email14@somewhere.com", "jwsMJNOk3oM3hVM5bGcF14", "1561650268514", []string{}}
		user2 := User{"215", "", "", "", "", "", "", "", "", "", false, "en", "email15@somewhere.com", "jwsMJNOk3oM3hVM5bGcF15", "1561650268515", []string{}}
		users := Users{}
		want := Users{}

		users.UnmarshalUsers(unregdata)
		want.UserList = append(Users{}.UserList, user1, user2)

		assertCorrectUsersStruct(t, users, want)
	})
}
