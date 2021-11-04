package main

import (
	"log"

	"github.com/spf13/cobra"
	initcmd "github.com/tmax-cloud/checkfile/cmd/checkfile/init"
	verifycmd "github.com/tmax-cloud/checkfile/cmd/checkfile/verify"
)

func main() {
	cmd := &cobra.Command{
		Use:   "checkfile [Command]",
		Short: "checkfile calculates/verifies checksum of file(s)",
	}

	cmd.AddCommand(initcmd.New())
	cmd.AddCommand(verifycmd.New())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
