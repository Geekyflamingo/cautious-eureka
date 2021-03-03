package userprojects

import (
	"datasite/errs"
	"datasite/project"
	"datasite/user"
	"fmt"
	"io/ioutil"
	"net/http"
)

const registeredUsersEndpoint string = "https://5c3ce12c29429300143fe570.mockapi.io/api/registeredusers"
const unregisteredusersEndpoint string = "https://5c3ce12c29429300143fe570.mockapi.io/api/unregisteredusers"
const projectsEndpoint string = "https://5c3ce12c29429300143fe570.mockapi.io/api/projectmemberships"

//GetUserProjects blah
func GetUserProjects(w http.ResponseWriter, r *http.Request) {
	registeredUsers := user.Users{}
	registeredUsersResp := getData(registeredUsersEndpoint)
	registeredUsers.UnmarshalUsers(registeredUsersResp)

	unregisteredUsers := user.Users{}
	unregisteredUsersResp := getData(unregisteredusersEndpoint)
	unregisteredUsers.UnmarshalUsers(unregisteredUsersResp)

	users := user.Users{}
	users.UserList = append(registeredUsers.UserList, unregisteredUsers.UserList...)

	var projects []project.Project = project.UnmarshalProjects(getData(projectsEndpoint))
	usersprojects := processUserProjects(users, projects)

	w.WriteHeader(http.StatusOK)
	resp, err := usersprojects.ToJSON()
	errs.HandleError(err)
	fmt.Fprintf(w, string(resp))
}

func processUserProjects(users user.Users, projects []project.Project) user.Users {
	for i, user := range users.UserList {
		for _, project := range projects {
			if user.ID == project.UserID {
				(&users).UserList[i].ProjectIDs = append(user.ProjectIDs, project.ID)
			}
		}
		if (&users).UserList[i].ProjectIDs == nil {
			(&users).UserList[i].ProjectIDs = []string{}
		}
	}
	return users
}

func getData(endpoint string) []byte {
	resp, err := http.Get(endpoint)
	errs.HandleError(err)

	body, err := ioutil.ReadAll(resp.Body)
	errs.HandleError(err)

	return body
}
