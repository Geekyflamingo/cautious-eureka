package project

import (
	"datasite/errs"
	"encoding/json"
)

//Project declares type for json destructuring
type Project struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	UserID    string `json:"userId"`
}

//UnmarshalProjects process projects data points
func UnmarshalProjects(resp []byte) []Project {
	//Convert the body to Project type
	var projects []Project
	err := json.Unmarshal(resp, &projects)
	errs.HandleError(err)
	return projects
}
