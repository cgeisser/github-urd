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
    {&github.Repository{HasIssues: github.Bool(false), OpenIssuesCount: github.Int(17)}, false},
	}
	for  _, test := range tests {
		if canTurnOffIssues(test.value) != test.want {
			t.Errorf("canTurnOffIssues on %v expected %v", test.value, test.want)
		}
	}
}
