package config

type ServerConfig struct {
	Name      string      `mapstructure:"name" json:"name"`
	Host      string      `mapstructure:"host" json:"host"`
	Port      int         `mapstructure:"port" json:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
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

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
