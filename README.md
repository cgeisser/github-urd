# github-urd [![Build Status](https://travis-ci.org/cgeisser/github-urd.svg?branch=master)](https://travis-ci.org/cgeisser/github-urd)

`urd` stands for Use Reasonable Defaults.

The goal of this tool is to enforce a standard GitHub configuration across many repositories.
The settings and checks are based on experiences working with organizations large and small to
standardize and streamline their development process.

Maybe you decided to use Jira for project management. Great! Turn off GitHub issues and projects.
Maybe you decided everyone should hook TravisCI into their PR process. Great! Audit that too.

# Run
Generate a [personal GitHub token.](https://github.com/settings/tokens)

```
GITHUB_AUTH_TOKEN=yourtoken go run github-urd.go --org=myorg
```

This program does *not* change any settings, it produces an audit log so you can
make changes manually.

# Audit GitHub Settings

This program checks for:

- Branch protection on the main branch with
  - admins included in the restrictions
  - code review approval
  - branch staleness check
- GitHub issues turned off via `--use_issues=false`. If you ask to turn off issues, projects will also be checked.
- GitHub wiki turned off via `--use_wiki=false`
- Hook installed for anything that looks like source code via `--code_hook_string="name_of_thing"`
