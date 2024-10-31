package config

type APIConfig struct {
	Host        string        `mapstructure:"host" json:"host,omitempty"`
	Port        string        `mapstructure:"port" json:"port,omitempty"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
}
type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host,omitempty"`
	Port string `mapstructure:"port" json:"port,omitempty"`
}
