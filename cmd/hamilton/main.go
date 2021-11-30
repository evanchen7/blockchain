package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	fmt.Println("Hamilton Blockchain initialized")
}

const flagDataDir = "datadir"

func main() {
	var hamiltonCmd = &cobra.Command{
		Use:   "hamilton",
		Short: "The Hamilton Blockchain CLI",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	hamiltonCmd.AddCommand(versionCmd)
	hamiltonCmd.AddCommand(balancesCmd())
	hamiltonCmd.AddCommand(txCmd())

	err := hamiltonCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data directory wher ethe DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
