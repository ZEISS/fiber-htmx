package cmd

// Config represents the configuration of the CLI.
type Config struct {
	Convert Convert
}

// Convert represents the configuration of the conversion command.
type Convert struct {
	File   string
	Output string
}

// DefaultConfig returns the default configuration.
func DefaultConfig() Config {
	return Config{
		Convert: Convert{},
	}
}
