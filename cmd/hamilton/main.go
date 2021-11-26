package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	fmt.Println("Hamlton Blockchain initialized")
}

func main() {
	var hamiltonCmd = &cobra.Command{
		Use:   "hamilton",
		Short: "The Hamilton Blockchain CLI",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	err := hamiltonCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
