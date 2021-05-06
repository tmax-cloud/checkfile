package verify

import (
	"fmt"
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
	"os"
)

// New returns a verify command
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify ",
		Short: "verify ",
		Run: func(cmd *cobra.Command, args []string) {
			if err := checksum.VerifySums(); err != nil {
				fmt.Println("[checkfile] " + err.Error())
				os.Exit(1)
			}
		},
	}
	return cmd
}
