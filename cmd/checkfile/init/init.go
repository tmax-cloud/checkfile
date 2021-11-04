package init

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmax-cloud/checkfile/pkg/checksum"
)

// New returns an init command
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init FILES...",
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
