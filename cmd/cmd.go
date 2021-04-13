package cmd

import (
	"Dscan/lib"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	target     string
	targetFile string
	ruleFile   string
	accurate   bool
	fuzz       bool
	dictFile   string

	rootCmd = &cobra.Command{
		Use:   "Dscan",
		Short: "A scanner for web directory",
		Long:  `Dscan is a tool for scanning web directory, which supports accurate directory scanning and custom rule sets.`,
	}

	scan = &cobra.Command{
		Use:   "scan [--accurate|--fuzz] [other flags]",
		Short: "Use accurate model or fuzz model to scan",
		Long:  `You can choose to scan in either accurate mode, which uses the scan rule set, or fuzz mode, which uses the directory keyword dictionary`,
		Run:   lib.Scan,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&targetFile, "file", "f", "", "target file to scan")
	rootCmd.PersistentFlags().StringVar(&ruleFile, "rule", "rules/rules.yaml", "rule file")
	rootCmd.PersistentFlags().Bool("accurate", false, "Whether to use accurate mode")
	rootCmd.AddCommand(scan)
}

// Execute cmd from here
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
