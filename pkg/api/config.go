package api

type Config struct {
	// API Server Settings
	URL     string `yaml:"url"`
	Version string `yaml:"version"`
	Timeout int    `yaml:"timeout"`
	// API Authentication Settings
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}
