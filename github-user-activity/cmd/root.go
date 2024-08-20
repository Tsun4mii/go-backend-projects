package cmd

import (
	"fmt"
	activity "github_user_activity/packages"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github-activity",
		Short: "Github User Activity is a CLI tool for fetching user activity",
		Long: `Github User Activity is a CLI tool for fetching user activity. It allows you to fetch user activity by providing the username.

Example:
> github-activity Tsun4mii`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDisplayActivityCmd(args)
		},
	}

	return cmd
}

func RunDisplayActivityCmd(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a username")
	}

	username := args[0]
	activities, err := activity.FetchGithubActivities(username)
	if err != nil {
		return err
	}

	return activity.DisplayActivity(username, activities)
}
