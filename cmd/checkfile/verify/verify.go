package verify

import (
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
	"os"
)

// New returns a verify command
func New() *cobra.Command {
	outputDest := "/dev/stdout"

	cmd := &cobra.Command{
		Use:   "verify ",
		Short: "verify ",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := checksum.VerifySums()
			if err != nil {
				if err2 := WriteString(err.Error(), "application/x-www-form-urlencoded", outputDest); err2 != nil {
					panic(err2)
				}
				panic(err)
			}

			if err := WriteResult(result, outputDest); err != nil {
				panic(err)
			}
			if result.IsTampered {
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&outputDest, "output", "o", "/dev/stdout", "output file path or http endpoint")

	return cmd
}
