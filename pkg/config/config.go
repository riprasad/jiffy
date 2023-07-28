package config

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type JiffyConfiguration struct {
	JiraToken string `mapstructure:"jira_token"`
	Jql       string `mapstructure:"jql"`
}

func InitConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/Jiffy")
	viper.AddConfigPath(".")

	viper.SetDefault("jira_token", "")
	viper.SetDefault("jql", "")

	_ = viper.BindEnv("jira_token", "JIRA_TOKEN")
	_ = viper.BindEnv("jql", "JQL")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetConfiguration() JiffyConfiguration {
	cfg := JiffyConfiguration{}
	err := viper.Unmarshal(&cfg)
	if err != nil {
		_ = errors.Errorf("unable to decode into struct, %v", err)
	}
	return cfg
}
