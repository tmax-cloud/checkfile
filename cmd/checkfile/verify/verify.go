package verify

import (
	"fmt"
	"os"

	"github.com/cqbqdd11519/checkfile/pkg/checksum"
	"github.com/spf13/cobra"
)

// New returns a verify command
func New() *cobra.Command {
	outputDests := []string{"/dev/stdout"}
	var outputTemplates []string

	cmd := &cobra.Command{
		Use:   "verify ",
		Short: "verify ",
		Run: func(cmd *cobra.Command, args []string) {
			// Set default templates
			if len(outputTemplates) == 0 {
				for range outputDests {
					outputTemplates = append(outputTemplates, "{{.}}")
				}
			}

			if len(outputDests) != len(outputTemplates) {
				panic(fmt.Errorf("length of output options and outputTemplate options should be identical"))
			}

			result, err := checksum.VerifySums()
			if err != nil {
				if err2 := WriteStrings(err.Error(), "application/x-www-form-urlencoded", outputDests, outputTemplates); err2 != nil {
					panic(err2)
				}
				panic(err)
			}

			if err := WriteResult(result, outputDests, outputTemplates); err != nil {
				panic(err)
			}
			if result.IsTampered {
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringSliceVarP(&outputDests, "output", "o", []string{"/dev/stdout"}, "output file path or http endpoint")
	cmd.PersistentFlags().StringSliceVarP(&outputTemplates, "outputTemplates", "t", []string{}, "output file path or http endpoint")

	return cmd
}
