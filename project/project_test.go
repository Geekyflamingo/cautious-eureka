package project

import "testing"

func TestUnmarshalProjects(t *testing.T) {
	// some JSON data
	data := []byte(`
		[{"id":"37","projectId":"37","userId":"214"},
		{"id":"38","projectId":"38","userId":"215"}]
	`)

	proj := []Project{}

	got := UnmarshalProjects(data)
	want := append(proj, Project{"37", "37", "214"}, Project{"38", "38", "215"})

	match := false
	for _, gp := range got {
		for _, wp := range want {
			if gp == wp {
				match = true
			}
		}
		if !match {
			t.Errorf("got %q want %q", got, want)
		}
	}
}
