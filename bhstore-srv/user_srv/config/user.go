package config

type Config struct {
	MysqlInfo MysqlConfig `json:"mysql,omitempty" mapstructure:"mysql" `
}

type MysqlConfig struct {
	Host     string `json:"host,omitempty" mapstructure:"host"`
	Port     string `json:"port,omitempty" mapstructure:"port"`
	Name     string `json:"name,omitempty" mapstructure:"name"`
	User     string `json:"user,omitempty" mapstructure:"user"`
	Password string `json:"password,omitempty" mapstructure:"password"`
}
