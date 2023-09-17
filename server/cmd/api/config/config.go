package config

type ServerConfig struct {
	Name                string    `mapstructure:"name" json:"name"`
	Host                string    `mapstructure:"host" json:"host"`
	Port                int       `mapstructure:"port" json:"port"`
	UserProviderService RPCConfig `mapstructure:"user_provider" json:"user_provider"`
}

type LocalConfig struct {
	Name         string       `mapstructure:"name" json:"name"`
	LogConfig    LogConfig    `mapstructure:"log" json:"log"`
	ConsulConfig ConsulConfig `mapstructure:"consul" json:"consul"`
}

type LogConfig struct {
	LogDir   string `mapstructure:"dir" json:"dir"`
	LogDebug bool   `mapstructure:"level" json:"level"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type RPCConfig struct {
	Name string `mapstructure:"name" json:"name"`
}
