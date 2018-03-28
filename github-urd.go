package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

var (
	org        = flag.String("org", "", "Organization in GitHub to audit.")
	use_issues = flag.Bool("use_issues", true, "Organization uses GitHub issues.")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *org == "" {
		log.Fatal("--org is required")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the selected organization
	opt := &github.RepositoryListByOrgOptions{}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, *org, opt)
		if err != nil {
			fmt.Printf("%v\n", err)
			break
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, repo := range allRepos {
		if repo.GetPermissions()["admin"] == false {
			continue
		}

		if canTurnOffIssues(*use_issues, repo) {
			fmt.Printf("Turn off github issues for: %s. Current open issues: %d\n",
				repo.GetFullName(),
				repo.GetOpenIssuesCount())
		}

		protection, _, err := client.Repositories.GetBranchProtection(ctx, repo.Owner.GetLogin(), repo.GetName(),
			repo.GetDefaultBranch())
		if err != nil {
			fmt.Printf("%s %v\n", repo.GetName(), err)
		} else {

			if protection == nil || protection.GetRequiredStatusChecks() == nil ||
				protection.GetEnforceAdmins() == nil ||
				protection.GetRequiredPullRequestReviews() == nil ||
				!(protection.RequiredStatusChecks.Strict &&
					protection.EnforceAdmins.Enabled) {
				fmt.Printf("Fix branch protection on: %s %s\n", *repo.FullName, *repo.DefaultBranch)
			}
		}
		hooks, _, err := client.Repositories.ListHooks(ctx, repo.Owner.GetLogin(), repo.GetName(),
			nil)

		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			if repo.GetLanguage() != "" && !repo.GetPrivate() && !hasCodeClimate(hooks) {
				fmt.Printf("Install codeclimate on: %s\n", repo.GetFullName())
			}
		}

	}
}

func canTurnOffIssues(use_issues bool, repo *github.Repository) bool {
	if !use_issues && repo.GetHasIssues() && repo.GetOpenIssuesCount() == 0 {
		return true
	}
	return false
}

func hasCodeClimate(hooks []*github.Hook) bool {
	for _, hook := range hooks {
		if strings.Contains(hook.String(), "codeclimate") && hook.GetActive() {
			return true
		}
	}
	return false
}
