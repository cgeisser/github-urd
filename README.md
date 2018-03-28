# github-urd [![Build Status](https://travis-ci.org/cgeisser/github-urd.svg?branch=master)](https://travis-ci.org/cgeisser/github-urd)

`urd` stands for Use Reasonable Defaults.

# Run
Generate a [personal GitHub token.](https://github.com/settings/tokens)

```
GITHUB_AUTH_TOKEN=yourtoken go run github-urd.go --org=myorg
```
# Audit GitHub Settings

This program checks for:

- Branch protection on the main branch with
  - admins included in the restrictions
  - code review approval
  - branch staleness check
- GitHub issues turned off via `--use_issues=false`
- CodeClimate hook installed for anything that looks like source code.
