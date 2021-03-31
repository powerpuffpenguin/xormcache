package cmd

import (
	"fmt"

	"github.com/powerpuffpenguin/xormcache/version"

	"github.com/spf13/cobra"
)

const (
	// App .
	App = `xormcache`
)

var v bool
var rootCmd = &cobra.Command{
	Use:   App,
	Short: `generate tools create full`,
	Run: func(cmd *cobra.Command, args []string) {
		if v {
			fmt.Println(version.Platform)
			fmt.Println(version.Tag)
			fmt.Println(version.Commit)
			fmt.Println(version.Date)
		} else {
			fmt.Println(App)
			fmt.Println(version.Platform)
			fmt.Println(version.Tag)
			fmt.Println(version.Commit)
			fmt.Println(version.Date)
			fmt.Printf(`Use "%v --help" for more information about this program.
`, App)
		}
	},
}

func init() {
	flags := rootCmd.Flags()
	flags.BoolVarP(&v,
		`version`,
		`v`,
		false,
		`show version`,
	)
}

// Execute run command
func Execute() error {
	return rootCmd.Execute()
}
