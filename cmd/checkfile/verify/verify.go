package verify

import (
	"os"

	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
)

// New returns a verify command
func New() *cobra.Command {
	outputDests := []string{"/dev/stdout"}

	cmd := &cobra.Command{
		Use:   "verify ",
		Short: "verify ",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := checksum.VerifySums()
			if err != nil {
				if err2 := WriteStrings(err.Error(), "application/x-www-form-urlencoded", outputDests); err2 != nil {
					panic(err2)
				}
				panic(err)
			}

			if err := WriteResult(result, outputDests); err != nil {
				panic(err)
			}
			if result.IsTampered {
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringSliceVarP(&outputDests, "output", "o", []string{"/dev/stdout"}, "output file path or http endpoint")

	return cmd
}
