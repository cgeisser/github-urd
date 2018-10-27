package main

import (
	"github.com/google/go-github/github"
	"testing"
)

func TestCanTurnOffIssues(t *testing.T) {
	var tests = []struct {
		value *github.Repository
		want  bool
	}{
		{&github.Repository{HasIssues: github.Bool(true), OpenIssuesCount: github.Int(17)}, false},
		{&github.Repository{HasIssues: github.Bool(true), OpenIssuesCount: github.Int(0)}, true},
		{&github.Repository{HasIssues: github.Bool(false), HasProjects: github.Bool(true), OpenIssuesCount: github.Int(0)}, true},
		{&github.Repository{HasIssues: github.Bool(true), HasProjects: github.Bool(true), OpenIssuesCount: github.Int(0)}, true},
		{&github.Repository{HasIssues: github.Bool(false), OpenIssuesCount: github.Int(17)}, false},
	}
	for _, test := range tests {
		if canTurnOffIssues(false, test.value) != test.want {
			t.Errorf("canTurnOffIssues on %v expected %v", test.value, test.want)
		}
		if canTurnOffIssues(true, test.value) == true {
			t.Errorf("canTurnOffIssues with issues allowed should be false.")
		}
	}
}

func TestCanTurnOffWiki(t *testing.T) {
	var tests = []struct {
		value *github.Repository
		want  bool
	}{
		{&github.Repository{HasWiki: github.Bool(false)}, false},
		{&github.Repository{HasWiki: github.Bool(true)}, true},
	}
	for _, test := range tests {
		if canTurnOffWiki(false, test.value) != test.want {
			t.Errorf("canTurnOffWiki on %v expected %v", test.value, test.want)
		}
		if canTurnOffWiki(true, test.value) == true {
			t.Errorf("canTurnOffWiki with wiki allowed should be false.")
		}
	}
}

func TestHasRequiredHook(t *testing.T) {
	var tests = []struct {
		hook_string string
		hooks       []*github.Hook
		want        bool
	}{
		{"", []*github.Hook{}, true},
		{"weasel", []*github.Hook{}, false},
		{"weasel", []*github.Hook{&github.Hook{Active: github.Bool(false), URL: github.String("weasel")}}, false},
		{"weasel", []*github.Hook{&github.Hook{Active: github.Bool(true), URL: github.String("weasel")}}, true},
		{"weasel", []*github.Hook{&github.Hook{Active: github.Bool(false), URL: github.String("notit")}}, false},
		{"weasel", []*github.Hook{&github.Hook{Active: github.Bool(true), URL: github.String("notit")}}, false},
	}
	for _, test := range tests {
		if hasRequiredHook(test.hook_string, test.hooks) != test.want {
			t.Errorf("hasRequiredHook on '%s' %v expected %v", test.hook_string, test.hooks, test.want)
		}
	}
}
