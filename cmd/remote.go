/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dlvgen/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var remoteOpts struct {
	Path string `json:"remotePath,omitempty"`
	Host string `json:"host,omitempty"`
	Port int
}

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remote called")

		options := config.ConfigOptions{
			ConfigType:   cmd.Use,
			Args:         opts.Args,
			EnvVars:      opts.EnvVars,
			ShowLog:      opts.ShowLog,
			StopOnEntry:  opts.StopOnEntry,
			BuildFlags:   opts.BuildFlags,
			WorkingDir:   opts.WorkingDir,
			ConsoleType:  opts.ConsoleType,
			OutputFormat: opts.Format,
			OutputFile:   opts.OutputFile,
			RemotePath:   remoteOpts.Path,
			RemoteHost:   remoteOpts.Host,
			RemotePort:   remoteOpts.Port,
		}

		launchCfg := config.GenerateConfig(options)
		config.OutputConfig(launchCfg, options)
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	remoteCmd.Flags().StringVarP(&remoteOpts.Path, "path", "", "", "Remote path")
	remoteCmd.Flags().StringVarP(&remoteOpts.Host, "host", "", "", "Host example: host=192.168.1.100 ")
	remoteCmd.Flags().IntVarP(&remoteOpts.Port, "port", "", 0, "Port example: port=2345")
}
