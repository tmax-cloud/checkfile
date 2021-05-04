package init

import (
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
)

// New returns an init command
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init ",
		Short: "init ",
		RunE: func(cmd *cobra.Command, args []string) error {
			return checksum.InitSumsDB(args)
		},
	}
	return cmd
}
