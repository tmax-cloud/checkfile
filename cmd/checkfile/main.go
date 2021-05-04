package main

import (
	initcmd "github.com/cqbqdd11519/checkfile/cmd/checkfile/init"
	verifycmd "github.com/cqbqdd11519/checkfile/cmd/checkfile/verify"
	"github.com/spf13/cobra"
	"log"
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
