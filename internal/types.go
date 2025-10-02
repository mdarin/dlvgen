package internal

type DebugConfig struct {
	Version string         `json:"version"`
	Configs []LaunchConfig `json:"configurations"`
}

type LaunchConfig struct {
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	Request     string            `json:"request"`
	Program     string            `json:"program,omitempty"`
	Args        []string          `json:"args,omitempty"`
	Env         map[string]string `json:"env,omitempty"`
	Mode        string            `json:"mode,omitempty"`
	RemotePath  string            `json:"remotePath,omitempty"`
	Host        string            `json:"host,omitempty"`
	Port        int               `json:"port,omitempty"`
	ShowLog     bool              `json:"showLog,omitempty"`
	LogOutput   string            `json:"logOutput,omitempty"`
	Console     string            `json:"console,omitempty"`
	StopOnEntry bool              `json:"stopOnEntry,omitempty"`
	BuildFlags  string            `json:"buildFlags,omitempty"`
	Cwd         string            `json:"cwd,omitempty"`
}

type ConfigOptions struct {
	OutputFile    string
	OutputFormat  string
	ConfigType    string
	ProgramPath   string
	Args          string
	EnvVars       string
	RemoteHost    string
	RemotePort    int
	RemotePath    string
	ShowLog       bool
	StopOnEntry   bool
	BuildFlags    string
	WorkingDir    string
	ConsoleType   string
	Verbose       bool
	ShowExamples  bool
	ListTemplates bool
}
