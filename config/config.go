package config

type (
	Config struct {
		Configs `mapstructure:"configs"`
	}

	Configs struct {
		Address     string      `mapstructure:"addr"`
		LogProperty LogProperty `mapstructure:"log_property"`
	}

	LogProperty struct {
		Level  string `mapstructure:"level"`
		Output string `mapstructure:"output"`
		Format string `mapstructure:"format"`
	}
)
