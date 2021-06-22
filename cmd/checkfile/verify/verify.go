package verify

import (
	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
	"os"
)

// New returns a verify command
func New() *cobra.Command {
	outputFileWriter := "/dev/stdout"

	cmd := &cobra.Command{
		Use:   "verify ",
		Short: "verify ",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := checksum.VerifySums()
			if err != nil {
				if err2 := WriteStringToFile(err.Error(), outputFileWriter); err2 != nil {
					panic(err)
				}
				panic(err)
			}

			if err := WriteToFile(result, outputFileWriter); err != nil {
				panic(err)
			}
			if result.IsTampered {
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&outputFileWriter, "outputFile", "o", "/dev/stdout", "output file path")

	return cmd
}
