package init

import (
	"fmt"
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
	"os"
)

// New returns an init command
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init ",
		Short: "init ",
		Run: func(cmd *cobra.Command, args []string) {
			if err := checksum.InitSumsDB(args); err != nil {
				fmt.Println("[checkfile] " + err.Error())
				os.Exit(1)
			}
		},
	}
	return cmd
}
