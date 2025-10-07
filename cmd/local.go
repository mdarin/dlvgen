/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dlvgen/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("local called")

		options := config.ConfigOptions{
			Args:         opts.Args,
			EnvVars:      opts.EnvVars,
			ShowLog:      opts.ShowLog,
			StopOnEntry:  opts.StopOnEntry,
			BuildFlags:   opts.BuildFlags,
			WorkingDir:   opts.WorkingDir,
			ConsoleType:  opts.ConsoleType,
			OutputFormat: opts.Format,
			OutputFile:   opts.OutputFile,
		}

		launchCfg := config.GenerateConfig(options)
		config.OutputConfig(launchCfg, options)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(localCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// localCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// localCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
