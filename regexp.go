package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"testing"
)

// RemoveEmptyTokens of a splitted string. Tokens such as "" are removed.
func RemoveEmptyTokens(ss []string) []string {
	res := []string{}
	for _, s := range ss {
		if s != "" {
			res = append(res, s)
		}
	}

	return res
}

// ArrayAsRegexAnyMatchExpression converting array to regexp string for matching any of elements
func ArrayAsRegexAnyMatchExpression(todos []string) string {
	if len(todos) == 0 {
		panic("Empty list of todo strings")
	}
	return fmt.Sprintf("(?:%s)", strings.Join(todos, "|"))
}

// GetCapturingGroups gets the submatches
func GetCapturingGroups(pattern *regexp.Regexp, input string) map[string]string {
	groups := map[string]string{}
	matches := pattern.FindStringSubmatch(input)
	for i, name := range pattern.SubexpNames() {
		groups[name] = matches[i]
	}

	return groups
}

func Test_IssueTracker_IssueURLFor(t *testing.T) {

	var tests = []struct {
		input  string
		taskID string
		want   string
	}{
		{"https://github.com/username1/todocheck", "1", "https://api.github.com/repos/username1/todocheck/issues/1"},
		{"github.com/username/todocheck/", "8", "https://api.github.com/repos/username/todocheck/issues/8"},
		{"github.com/uSER-1989/todocheck/", "8", "https://api.github.com/repos/user-1989/todocheck/issues/8"},
		{"github.com/user2020/hyphen-1/", "8", "https://api.github.com/repos/user2020/hyphen-1/issues/8"},
		{"github.com/user2020/hyphen-1_underscore/", "8", "https://api.github.com/repos/user2020/hyphen-1_underscore/issues/8"},
		{"GITHUB.com/u1u/hyphen-1_underscore_x-2/", "8", "https://api.github.com/repos/u1u/hyphen-1_underscore_x-2/issues/8"},
	}

	for _, tt := range tests {

		var it IssueTracker

		testname := fmt.Sprintf("%q", tt.input)
		t.Run(testname, func(t *testing.T) {
			it.Origin = tt.input
			res := it.IssueURLFor(tt.taskID)
			if res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}

}

func Test_IssueTracker_IssueURLFor(t *testing.T) {

	var tests = []struct {
		input  string
		taskID string
		want   string
	}{
		{"gitlab.com/Username/project", "1", "https://gitlab.com/api/v4/projects/username%2Fproject/issues/1"},
		{"gitlab.com/user-5/project8-9_1", "1111", "https://gitlab.com/api/v4/projects/user-5%2Fproject8-9_1/issues/1111"},
		{"https://gitLAB.com/user-1/project", "2020", "https://gitlab.com/api/v4/projects/user-1%2Fproject/issues/2020"},
		{"https://gitlab.com/u/project9-1_2-3-4", "2020", "https://gitlab.com/api/v4/projects/u%2Fproject9-1_2-3-4/issues/2020"},
		{"gitlab.myorg.com/Username/project", "1", "https://gitlab.myorg.com/api/v4/projects/username%2Fproject/issues/1"},
		{"gitlab.myorg.com/user-5/project8-9_1", "1111", "https://gitlab.myorg.com/api/v4/projects/user-5%2Fproject8-9_1/issues/1111"},
		{"https://gitLAB.myorg.com/user-1/project", "2020", "https://gitlab.myorg.com/api/v4/projects/user-1%2Fproject/issues/2020"},
		{"https://gitlab.myORG.com/u/project9-1_2-3-4", "2020", "https://gitlab.myorg.com/api/v4/projects/u%2Fproject9-1_2-3-4/issues/2020"},
		{"myorg.com/Username/project", "20201225", "https://myorg.com/api/v4/projects/username%2Fproject/issues/20201225"},
		{"myorg.co.uk/Username/project", "20201226", "https://myorg.co.uk/api/v4/projects/username%2Fproject/issues/20201226"},
	}

	for _, tt := range tests {

		var it IssueTracker

		testname := fmt.Sprintf("%q", tt.input)
		t.Run(testname, func(t *testing.T) {
			it.Origin = tt.input
			res := it.IssueURLFor(tt.taskID)
			if res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}

}

func Test_IssueTracker_IssueURLFor(t *testing.T) {
	var tests = []struct {
		input  string
		taskID string
		want   string
	}{
		{"https://pivotaltracker.com/n/projects/1", "2020", "https://www.pivotaltracker.com/services/v5/projects/1/stories/2020"},
		{"pivotaltracker.com/n/projects/1-2_3", "2020", "https://www.pivotaltracker.com/services/v5/projects/1-2_3/stories/2020"},
		{"pIVOTALtracker.com/projects/1-2_3", "2020", "https://www.pivotaltracker.com/services/v5/projects/1-2_3/stories/2020"},
		{"httPS://pivotaltracker.com/projects/1-2_3", "2020", "https://www.pivotaltracker.com/services/v5/projects/1-2_3/stories/2020"},
	}

	for _, tt := range tests {

		var it IssueTracker

		testname := fmt.Sprintf("%q", tt.input)
		t.Run(testname, func(t *testing.T) {
			it.Origin = tt.input
			res := it.IssueURLFor(tt.taskID)
			if res != tt.want {
				t.Errorf("got %s, want %s", res, tt.want)
			}
		})
	}

}

//github
func (it *IssueTracker) issueAPIOrigin() string {

	pattern := regexp.MustCompile(`^(?P<scheme>^(https?:))//github.com/(?P<owner>[a-zA-Z0-9\-]+)/(?P<repo>[a-zA-Z0-9-_.]+)/?$`)

	origin := strings.ToLower(it.Origin)

	if !strings.HasPrefix(origin, "http") {
		origin = "https://" + origin
	}

	groups := GetCapturingGroups(pattern, origin)

	return fmt.Sprintf("%s//api.github.com/repos/%s/%s/issues/", groups["scheme"], groups["owner"], groups["repo"])
}

//gitlab
func (it *IssueTracker) issueAPIOrigin() string {

	pattern := regexp.MustCompile(`^(?P<scheme>^(https?:))//(?P<domain>[a-zA-Z0-9\-.]+)/(?P<owner>[a-zA-Z0-9\-]+)/(?P<repo>[a-zA-Z0-9-_.]+)/?$`)

	origin := strings.ToLower(it.Origin)

	if !strings.HasPrefix(origin, "http:") && !strings.HasPrefix(origin, "https:") {
		origin = "https://" + origin
	}

	groups := GetCapturingGroups(pattern, origin)

	urlEncodedProject := url.QueryEscape(fmt.Sprintf("%s/%s", groups["owner"], groups["repo"]))
	return fmt.Sprintf("%s//%s/api/v4/projects/%s/issues/", groups["scheme"], groups["domain"], urlEncodedProject)
}

//pivotaltracket
func (it *IssueTracker) issueAPIOrigin() string {

	origin := strings.ToLower(it.Origin)

	pattern := regexp.MustCompile(`^(?P<scheme>^(https?:))//pivotaltracker.com/(n/)?projects/(?P<project>[a-zA-Z0-9-_.]+)`)

	if !strings.HasPrefix(origin, "http:") && !strings.HasPrefix(origin, "https:") {
		origin = "https://" + origin
	}

	groups := GetCapturingGroups(pattern, origin)

	return fmt.Sprintf("%s//www.pivotaltracker.com/services/v5/projects/%s/stories/", groups["scheme"], groups["project"])
}
