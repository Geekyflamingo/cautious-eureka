package user

import "testing"

var unregdata = []byte(`[{"id":"214","emailAddress":"email14@somewhere.com","languageCode":"en","registrationId":"jwsMJNOk3oM3hVM5bGcF14","registrationIdGeneratedTime":"1561650268514"},{"id":"215","emailAddress":"email15@somewhere.com","languageCode":"en","registrationId":"jwsMJNOk3oM3hVM5bGcF15","registrationIdGeneratedTime":"1561650268515"}]`)

var regdata = []byte(`[{"id":"1","city":"Jaydashire","company":"Goyette - Renner","country":"South Africa","firstName":"firstName 1","lastName":"lastName 1","organizationType":"organizationType 1","phone":"524.276.1570 x487","state":"SD","zipCode":"68048","disclaimerAccepted":false,"languageCode":"en","emailAddress":"last1@mail.com"},{"id":"2","city":"Adolfofort","company":"Fisher - Bartoletti","country":"China","firstName":"firstName 2","lastName":"lastName 2","organizationType":"organizationType 2","phone":"(308) 197-9774 x339","state":"CO","zipCode":"78569","disclaimerAccepted":false,"languageCode":"en","emailAddress":"last2@mail.com"}]`)

func TestUnmarshalUsers(t *testing.T) {

	assertCorrectUsersStruct := func(t testing.TB, got, want []User) {
		t.Helper()
		match := false
		for _, gu := range got {
			for _, wu := range want {
				if gu.ID == wu.ID &&
					gu.City == wu.City &&
					gu.Company == wu.Company &&
					gu.Country == wu.Country &&
					gu.FirstName == wu.FirstName &&
					gu.OrganizationType == wu.OrganizationType &&
					gu.Phone == wu.Phone &&
					gu.State == wu.State &&
					gu.ZipCode == wu.ZipCode &&
					gu.DisclaimerAccepted == wu.DisclaimerAccepted &&
					gu.LanguageCode == wu.LanguageCode &&
					gu.Email == wu.Email &&
					gu.RegistrationID == wu.RegistrationID &&
					gu.RegistrationIDGeneratedTime == wu.RegistrationIDGeneratedTime {
					match = true
				}
			}
			if !match {
				t.Errorf("got %v want %v", got, want)
			}
		}

	}

	t.Run("Registered Users", func(t *testing.T) {

		users := []User{}

		got := UnmarshalUsers(regdata)
		want := append(users, User{"1", "Jaydashire", "Goyette - Renner", "South Africa", "firstName 1", "lastName 1", "organizationType 1", "524.276.1570 x487", "SD", "68048", false, "en", "last1@mail.com", "", "", []string{}}, User{"2", "Adolfofort", "Fisher - Bartoletti", "China", "firstName 2", "lastName 2", "organizationType 2", "(308) 197-9774 x339", "CO", "78569", false, "en", "last2@mail.com", "", "", []string{}})

		assertCorrectUsersStruct(t, got, want)

	})

	t.Run("Unregistered Users", func(t *testing.T) {

		users := []User{}

		got := UnmarshalUsers(unregdata)
		want := append(users, User{"214", "", "", "", "", "", "", "", "", "", false, "en", "email14@somewhere.com", "jwsMJNOk3oM3hVM5bGcF14", "1561650268514", []string{}}, User{"215", "", "", "", "", "", "", "", "", "", false, "en", "email15@somewhere.com", "jwsMJNOk3oM3hVM5bGcF15", "1561650268515", []string{}})

		assertCorrectUsersStruct(t, got, want)
	})
}
