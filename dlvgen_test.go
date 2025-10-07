package main

/*
func TestParseArgs(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"", nil},
		{"arg1", []string{"arg1"}},
		{"arg1,arg2,arg3", []string{"arg1", "arg2", "arg3"}},
		{"--port=8080,--verbose", []string{"--port=8080", "--verbose"}},
	}

	for _, test := range tests {
		result := parseArgs(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("parseArgs(%q) = %v, expected %v", test.input, result, test.expected)
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("parseArgs(%q) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestParseEnvVars(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]string
	}{
		{"", nil},
		{"KEY=value", map[string]string{"KEY": "value"}},
		{"KEY1=value1,KEY2=value2", map[string]string{"KEY1": "value1", "KEY2": "value2"}},
		{"DEBUG=true,LOG_LEVEL=debug", map[string]string{"DEBUG": "true", "LOG_LEVEL": "debug"}},
	}

	for _, test := range tests {
		result := parseEnvVars(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("parseEnvVars(%q) = %v, expected %v", test.input, result, test.expected)
		}
		for k, v := range test.expected {
			if result[k] != v {
				t.Errorf("parseEnvVars(%q)[%q] = %q, expected %q", test.input, k, result[k], v)
			}
		}
	}
}

func TestCreateLocalConfig(t *testing.T) {
	opts := ConfigOptions{
		ProgramPath: "./main.go",
		Args:        "arg1,arg2",
		EnvVars:     "DEBUG=true",
		WorkingDir:  "/test",
	}

	config := createLocalConfig("./main.go", opts)

	if config.Type != "go" {
		t.Errorf("Expected type 'go', got %s", config.Type)
	}
	if config.Request != "launch" {
		t.Errorf("Expected request 'launch', got %s", config.Request)
	}
	if config.Program != "./main.go" {
		t.Errorf("Expected program './main.go', got %s", config.Program)
	}
	if len(config.Args) != 2 {
		t.Errorf("Expected 2 args, got %d", len(config.Args))
	}
	if config.Env["DEBUG"] != "true" {
		t.Errorf("Expected DEBUG=true, got %s", config.Env["DEBUG"])
	}
	if config.Cwd != "/test" {
		t.Errorf("Expected cwd '/test', got %s", config.Cwd)
	}
}

func TestCreateRemoteConfig(t *testing.T) {
	opts := ConfigOptions{
		RemoteHost: "192.168.1.100",
		RemotePort: 2345,
		RemotePath: "/remote/path",
		WorkingDir: "/test",
	}

	config := createRemoteConfig("", opts)

	if config.Type != "go" {
		t.Errorf("Expected type 'go', got %s", config.Type)
	}
	if config.Request != "attach" {
		t.Errorf("Expected request 'attach', got %s", config.Request)
	}
	if config.Mode != "remote" {
		t.Errorf("Expected mode 'remote', got %s", config.Mode)
	}
	if config.Host != "192.168.1.100" {
		t.Errorf("Expected host '192.168.1.100', got %s", config.Host)
	}
	if config.Port != 2345 {
		t.Errorf("Expected port 2345, got %d", config.Port)
	}
	if config.RemotePath != "/remote/path" {
		t.Errorf("Expected remote path '/remote/path', got %s", config.RemotePath)
	}
}

func TestConfigMarshaling(t *testing.T) {
	config := DebugConfig{
		Version: "0.2.0",
		Configs: []LaunchConfig{
			{
				Name:    "Test Config",
				Type:    "go",
				Request: "launch",
				Program: "./main.go",
			},
		},
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	var decoded DebugConfig
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	if decoded.Version != config.Version {
		t.Errorf("Version mismatch: expected %s, got %s", config.Version, decoded.Version)
	}
	if len(decoded.Configs) != len(config.Configs) {
		t.Errorf("Configs length mismatch: expected %d, got %d", len(config.Configs), len(decoded.Configs))
	}
}

func TestFindMainProgram(t *testing.T) {
	// Create a temporary directory structure for testing
	tmpDir := t.TempDir()

	// Create a main.go file
	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`

	err := os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(mainContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test main.go: %v", err)
	}

	// Change to temp directory and test
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	os.Chdir(tmpDir)

	found := findMainProgram()
	if !strings.Contains(found, "main.go") {
		t.Errorf("Expected to find main.go, got %s", found)
	}
}

func TestConfigGeneration(t *testing.T) {
	opts := ConfigOptions{
		ConfigType:   "local",
		ProgramPath:  "./main.go",
		OutputFormat: "json",
	}

	config := generateConfig(opts)

	if config.Version != "0.2.0" {
		t.Errorf("Expected version '0.2.0', got %s", config.Version)
	}
	if len(config.Configs) != 1 {
		t.Errorf("Expected 1 config, got %d", len(config.Configs))
	}
	if config.Configs[0].Type != "go" {
		t.Errorf("Expected config type 'go', got %s", config.Configs[0].Type)
	}
}
*/
