package view

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ankitpokhrel/jira-cli/pkg/jira"
	"github.com/ankitpokhrel/jira-cli/pkg/tui"
)

func TestIssueData(t *testing.T) {
	issue := IssueList{
		Total:   2,
		Project: "TEST",
		Server:  "https://test.local",
		Data: []jira.Issue{
			{
				Key: "TEST-1",
				Fields: jira.IssueFields{
					Summary: "This is a test",
					Resolution: struct {
						Name string `json:"name"`
					}{Name: "Fixed"},
					IssueType: struct {
						Name string `json:"name"`
					}{Name: "Bug"},
					Assignee: struct {
						Name string `json:"displayName"`
					}{Name: "Person A"},
					Priority: struct {
						Name string `json:"name"`
					}{Name: "High"},
					Reporter: struct {
						Name string `json:"displayName"`
					}{Name: "Person Z"},
					Status: struct {
						Name string `json:"name"`
					}{Name: "Done"},
					Created: "2020-12-13T14:05:20.974+0100",
					Updated: "2020-12-13T14:07:20.974+0100",
				},
			},
			{
				Key: "TEST-2",
				Fields: jira.IssueFields{
					Summary: "This is another test",
					IssueType: struct {
						Name string `json:"name"`
					}{Name: "Story"},
					Priority: struct {
						Name string `json:"name"`
					}{Name: "Normal"},
					Reporter: struct {
						Name string `json:"displayName"`
					}{Name: "Person A"},
					Status: struct {
						Name string `json:"name"`
					}{Name: "Open"},
					Created: "2020-12-13T14:05:20.974+0100",
					Updated: "2020-12-13T14:07:20.974+0100",
				},
			},
		},
	}

	expected := tui.TableData{
		[]string{
			"TYPE", "KEY", "SUMMARY", "ASSIGNEE", "REPORTER", "PRIORITY", "STATUS", "RESOLUTION",
			"CREATED", "UPDATED",
		},
		[]string{
			"Bug", "TEST-1", "This is a test", "Person A", "Person Z", "High", "Done", "Fixed",
			"2020-12-13 14:05:20", "2020-12-13 14:07:20",
		},
		[]string{
			"Story", "TEST-2", "This is another test", "", "Person A", "Normal", "Open", "",
			"2020-12-13 14:05:20", "2020-12-13 14:07:20",
		},
	}

	assert.Equal(t, expected, issue.data())
}