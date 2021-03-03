package userprojects

import (
	"datasite/project"
	"datasite/user"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const registeredUsersEndpoint string = "https://5c3ce12c29429300143fe570.mockapi.io/api/registeredusers"
const unregisteredusersEndpoint string = "https://5c3ce12c29429300143fe570.mockapi.io/api/unregisteredusers"
const projectsEndpoint string = "https://5c3ce12c29429300143fe570.mockapi.io/api/projectmemberships"

//GetUserProjects blah
func GetUserProjects(w http.ResponseWriter, r *http.Request) {
	users := append(user.UnmarshalUsers(getData(registeredUsersEndpoint)),
		user.UnmarshalUsers(getData(unregisteredusersEndpoint))...)

	projects := project.UnmarshalProjects(getData(projectsEndpoint))

	for i, user := range users {
		for _, project := range projects {
			if user.ID == project.UserID {
				(&users[i]).ProjectIDs = append(user.ProjectIDs, project.ID)
			}
		}
		if (&users[i]).ProjectIDs == nil {
			(&users[i]).ProjectIDs = []string{}
		}
	}
	w.WriteHeader(http.StatusOK)
	resp, _ := json.MarshalIndent(users, "", "  ")
	fmt.Fprintf(w, string(resp))
}

func getData(endpoint string) []byte {
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
