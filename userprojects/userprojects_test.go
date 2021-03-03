package userprojects

import (
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
