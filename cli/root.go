package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-expr-finder",
	Short: "find expr in go code",
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
