package main

import (
	"fmt"
	"os"

	"github.com/evanchen7/blockchain/fs"
	"github.com/spf13/cobra"
)

const (
	flagDataDir = "datadir"
	flagPort    = "port"
	flagIP      = "ip"
)

func main() {
	var hamiltonCmd = &cobra.Command{
		Use:   "hamilton",
		Short: "The Hamilton Blockchain CLI",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	hamiltonCmd.AddCommand(versionCmd)
	hamiltonCmd.AddCommand(balancesCmd())
	hamiltonCmd.AddCommand(runCmd())
	hamiltonCmd.AddCommand(migrateCmd())

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

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)

	return fs.ExpandPath(dataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
