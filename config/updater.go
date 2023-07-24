package config

import (
	"github.com/tcnksm/go-latest"
)

func IsLatestVersion(version string) bool {
	githubTag := &latest.GithubTag{
		Owner:      "AllanCapistrano",
		Repository: "gbc",
	}

	res, _ := latest.Check(githubTag, version)

	return !res.Outdated
}
