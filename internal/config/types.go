package config

// ProjectConfig represents a project entry in the config file.
type ProjectConfig struct {
	Name    string `yaml:"name"`
	Session string `yaml:"session"`
	Path    string `yaml:"path"`
	// Layout can be a string (reference) or inline LayoutConfig
	Layout interface{}       `yaml:"layout,omitempty"`
	Vars   map[string]string `yaml:"vars,omitempty"`
}

// ToolConfig defines an external tool command.
type ToolConfig struct {
	Cmd        string `yaml:"cmd"`
	WindowName string `yaml:"window_name"`
}

// ToolsConfig groups tool definitions.
type ToolsConfig struct {
	CursorAgent ToolConfig `yaml:"cursor_agent,omitempty"`
	CodexNew    ToolConfig `yaml:"codex_new,omitempty"`
	CodexResume ToolConfig `yaml:"codex_resume,omitempty"`
}

// TmuxSection holds tmux-specific config.
type TmuxSection struct {
	Config string `yaml:"config,omitempty"`
}
