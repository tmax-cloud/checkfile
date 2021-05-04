package verify

import (
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
)

// New returns a verify command
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify ",
		Short: "verify ",
		RunE: func(cmd *cobra.Command, args []string) error {
			return checksum.VerifySums(args)
		},
	}
	return cmd
}
