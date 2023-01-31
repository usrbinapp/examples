package cmd

import "github.com/usrbinapp/usrbin-go"

func NewSDK(currentVersion string) (*usrbin.SDK, error) {
	return usrbin.New(
		currentVersion,
		usrbin.UsingGitHubUpdateChecker("github.com/usrbinapp/examples"),
		usrbin.UsingHomebrewFormula("usrbinapp/examples/cobra"),
	)
}
