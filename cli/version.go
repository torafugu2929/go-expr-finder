package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	goexprfinder "github.com/torafugu2929/go-expr-finder"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-expr-finder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(goexprfinder.Version())
	},
}
