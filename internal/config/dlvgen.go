package config

import (
	"dlvgen/internal/finder"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func GenerateConfig(opts ConfigOptions) DebugConfig {
	slog.Debug("Generating configuration", "type", opts.ConfigType)

	programPath := finder.FindMainProgram([]string{opts.ProgramPath})

	configs := []LaunchConfig{}

	switch opts.ConfigType {
	case "local":
		configs = append(configs, createLocalConfig(programPath, opts))

	case "remote":
		configs = append(configs, createRemoteConfig(programPath, opts))

	case "test":
		configs = append(configs, createTestConfig(opts))

	default:
		slog.Warn("Unknown config type, using local", "type", opts.ConfigType)
		configs = append(configs, createLocalConfig(programPath, opts))
	}

	return DebugConfig{
		Version: "0.2.0",
		Configs: configs,
	}
}

func createLocalConfig(programPath string, opts ConfigOptions) LaunchConfig {
	cwd := opts.WorkingDir
	if cwd == "" {
		cwd, _ = os.Getwd()
	}

	name := fmt.Sprintf("Launch debug [%s]", filepath.Base(cwd))

	config := LaunchConfig{
		Name:        name,
		Type:        "go",
		Request:     "launch",
		Program:     programPath,
		Args:        parseArgs(opts.Args),
		Env:         parseEnvVars(opts.EnvVars),
		Mode:        "debug",
		ShowLog:     opts.ShowLog,
		StopOnEntry: opts.StopOnEntry,
		BuildFlags:  opts.BuildFlags,
		Cwd:         cwd,
		Console:     opts.ConsoleType,
	}

	slog.Debug("Created local debug configuration", "program", programPath)
	return config
}

func createRemoteConfig(programPath string, opts ConfigOptions) LaunchConfig {
	cwd := opts.WorkingDir
	if cwd == "" {
		cwd, _ = os.Getwd()
	}

	name := fmt.Sprintf("Remote debug [%s]", filepath.Base(cwd))

	config := LaunchConfig{
		Name:       name,
		Type:       "go",
		Request:    "attach",
		Mode:       "remote",
		RemotePath: opts.RemotePath,
		Host:       opts.RemoteHost,
		Port:       opts.RemotePort,
		ShowLog:    opts.ShowLog,
		Cwd:        cwd,
		Console:    opts.ConsoleType,
	}

	slog.Debug("Created remote debug configuration", "host", opts.RemoteHost, "port", opts.RemotePort)
	return config
}

func createTestConfig(opts ConfigOptions) LaunchConfig {
	cwd := opts.WorkingDir
	if cwd == "" {
		cwd, _ = os.Getwd()
	}

	name := fmt.Sprintf("Test debug [%s]", filepath.Base(cwd))

	config := LaunchConfig{
		Name:    name,
		Type:    "go",
		Request: "test",
		Mode:    "test",
		Args:    parseArgs(opts.Args),
		Env:     parseEnvVars(opts.EnvVars),
		ShowLog: opts.ShowLog,
		Cwd:     cwd,
		Console: opts.ConsoleType,
	}

	slog.Debug("Created test debug configuration")
	return config
}

// func findMainProgram() string {
// 	slog.Debug("Searching for main.go files")

// 	searchPaths := []string{
// 		".",
// 		"cmd",
// 		"app",
// 		"src",
// 		"main",
// 	}

// 	var candidates []string

// 	for _, path := range searchPaths {
// 		err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
// 			if err != nil {
// 				return nil
// 			}

// 			if !info.IsDir() && strings.HasSuffix(filePath, ".go") {
// 				content, err := os.ReadFile(filePath)
// 				if err == nil {
// 					if strings.Contains(string(content), "package main") &&
// 						strings.Contains(string(content), "func main()") {
// 						candidates = append(candidates, filePath)
// 						slog.Debug("Found main program candidate", "file", filePath)
// 					}
// 				}
// 			}
// 			return nil
// 		})

// 		if err != nil {
// 			slog.Debug("Error walking path", "path", path, "error", err)
// 		}
// 	}

// 	if len(candidates) > 0 {
// 		// Prioritize by depth and common patterns
// 		bestCandidate := candidates[0]
// 		for _, candidate := range candidates {
// 			if strings.Contains(candidate, "cmd/") || strings.Contains(candidate, "main/") {
// 				bestCandidate = candidate
// 				break
// 			}
// 			if strings.Count(candidate, string(filepath.Separator)) < strings.Count(bestCandidate, string(filepath.Separator)) {
// 				bestCandidate = candidate
// 			}
// 		}
// 		slog.Info("Selected main program", "file", bestCandidate)
// 		return bestCandidate
// 	}

// 	slog.Warn("No main.go files found, using default")
// 	return "./main.go"
// }

func parseArgs(argsStr string) []string {
	if argsStr == "" {
		return nil
	}
	return strings.Split(argsStr, ",")
}

func parseEnvVars(envStr string) map[string]string {
	if envStr == "" {
		return nil
	}

	env := make(map[string]string)
	pairs := strings.Split(envStr, ",")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			env[kv[0]] = kv[1]
		}
	}
	return env
}

func OutputConfig(config DebugConfig, opts ConfigOptions) {
	var output []byte
	var err error

	if opts.OutputFormat == "pretty" {
		output, err = json.MarshalIndent(config, "", "  ")
	} else {
		output, err = json.Marshal(config)
	}

	if err != nil {
		slog.Error("Failed to marshal config", "error", err)
		os.Exit(1)
	}

	if opts.OutputFile != "" {
		err = os.WriteFile(opts.OutputFile, output, 0644)
		if err != nil {
			slog.Error("Failed to write output file", "file", opts.OutputFile, "error", err)
			os.Exit(1)
		}
		slog.Info("Configuration written to file", "file", opts.OutputFile)
	} else {
		fmt.Println(string(output))
	}
}

func showExamples() {
	fmt.Printf("%s\n\n", Blue("DlvGen Usage Examples"))

	fmt.Printf("%s\n", Yellow("1. Basic local debug configuration:"))
	fmt.Printf("  dlvgen\n\n")

	fmt.Printf("%s\n", Yellow("2. Local debug with custom program and arguments:"))
	fmt.Printf("  dlvgen -p ./cmd/myapp/main.go -a \"--port=8080,--verbose\"\n\n")

	fmt.Printf("%s\n", Yellow("3. Remote debug configuration:"))
	fmt.Printf("  dlvgen -t remote --host=192.168.1.100 --port=2345\n\n")

	fmt.Printf("%s\n", Yellow("4. Test debug configuration:"))
	fmt.Printf("  dlvgen -t test -a \"-v,-race\"\n\n")

	fmt.Printf("%s\n", Yellow("5. With environment variables:"))
	fmt.Printf("  dlvgen -e \"DEBUG=true,LOG_LEVEL=debug\"\n\n")

	fmt.Printf("%s\n", Yellow("6. Save to .vscode/launch.json:"))
	fmt.Printf("  dlvgen -o .vscode/launch.json --format=pretty\n\n")

	fmt.Printf("%s\n", Yellow("7. Verbose output with custom build flags:"))
	fmt.Printf("  dlvgen -v --build-flags=\"-tags=development\" --stop-on-entry\n\n")
}

func listTemplates() {
	fmt.Printf("%s\n\n", Blue("Available Configuration Templates"))

	templates := map[string]string{
		"local":  "Local development debugging",
		"remote": "Remote application debugging",
		"test":   "Go test debugging",
	}

	for name, desc := range templates {
		fmt.Printf("  %s: %s\n", Green(name), desc)
	}
	fmt.Println()
}
