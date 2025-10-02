/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// opts хранит значения общих (persistent) флагов
var opts struct {
	OutputFile  string
	Format      string
	ProgramPath string
	Args        string
	EnvVars     string
	ShowLog     bool
	StopOnEntry bool
	BuildFlags  string
	WorkingDir  string
	ConsoleType string
	Verbose     bool
}

var blue = color.New(color.FgBlue).SprintFunc()

var rootCmd = &cobra.Command{
	Use:   "dlvgen",
	Short: fmt.Sprintf("%s - A VS Code Debug Configuration Generator for Go", blue("dlvgen")),
	Long: `dlvgen is a CLI tool that simplifies the creation of VS Code's launch.json
files for debugging Go applications.

It can generate configurations for local, remote, and test debugging scenarios,
automatically detecting the main entry point of your application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		setupLogging(opts.Verbose)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}

func init() {
	// Persistent flags are available to all subcommands
	rootCmd.PersistentFlags().StringVarP(&opts.OutputFile, "output", "o", "", "Output file (default: stdout)")
	rootCmd.PersistentFlags().StringVarP(&opts.Format, "format", "f", "pretty", "Output format: json|pretty")
	rootCmd.PersistentFlags().BoolVarP(&opts.Verbose, "verbose", "v", false, "Enable verbose logging")

	// Flags for local/remote/test configs
	rootCmd.PersistentFlags().StringVarP(&opts.ProgramPath, "program", "p", "", "Main program path (auto-detected if empty)")
	rootCmd.PersistentFlags().StringVarP(&opts.Args, "args", "a", "", "Program arguments (comma-separated)")
	rootCmd.PersistentFlags().StringVarP(&opts.EnvVars, "env", "e", "", "Environment variables (key=value,key2=value2)")
	rootCmd.PersistentFlags().StringVar(&opts.BuildFlags, "build-flags", "", "Build flags for compilation")
	rootCmd.PersistentFlags().StringVar(&opts.WorkingDir, "cwd", "", "Working directory (default: current dir)")
	rootCmd.PersistentFlags().StringVar(&opts.ConsoleType, "console", "integratedTerminal", "Console type: integratedTerminal|externalTerminal|internalConsole")
	rootCmd.PersistentFlags().BoolVar(&opts.ShowLog, "show-log", false, "Show delve's log output")
	rootCmd.PersistentFlags().BoolVar(&opts.StopOnEntry, "stop-on-entry", false, "Stop on entry point")

	// Add subcommands
	rootCmd.AddCommand(localCmd, remoteCmd, testCmd, examplesCmd, listCmd)
}

func setupLogging(verbose bool) {
	level := slog.LevelWarn
	if verbose {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})))
	slog.Debug("Verbose logging enabled")
}
