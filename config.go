package main

// Configuration Struct
type Config struct {
	CommandRepoUrl  string // The input repository URL
	OutputRepoUrl   string // The output repository URL
	RefreshInterval int    // Delay in seconds
	IsIgnoringError bool   // Determines error handling behavior
	Timeout         int    // Timeout in seconds
}

// LoadConfig initializes and returns a default configuration
func LoadConfig() *Config {
	return &Config{

		CommandRepoUrl:  "https://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:x-oauth-basic@github.com/priyank49/Upload",
		OutputRepoUrl:   "https://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:x-oauth-basic@github.com/priyank49/download",
		RefreshInterval: 60,   // seconds
		Timeout:         0,    // seconds
		IsIgnoringError: true, // continue even if error
	}
}
