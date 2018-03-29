package main

import (
	"github.com/google/go-github/github"
	"testing"
)

func TestCanTurnOffIssues(t *testing.T) {
	var tests = []struct {
		use_issues bool
		value      *github.Repository
		want       bool
	}{
		{false, &github.Repository{HasIssues: github.Bool(true), OpenIssuesCount: github.Int(17)}, false},
		{false, &github.Repository{HasIssues: github.Bool(true), OpenIssuesCount: github.Int(0)}, true},
		{false, &github.Repository{HasIssues: github.Bool(false), OpenIssuesCount: github.Int(17)}, false},
	}
	for _, test := range tests {
		if canTurnOffIssues(test.use_issues, test.value) != test.want {
			t.Errorf("canTurnOffIssues on %v expected %v", test.value, test.want)
		}
		if canTurnOffIssues(true, test.value) == true {
			t.Errorf("canTurnOffIssues with issues allowed should be false.")
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
