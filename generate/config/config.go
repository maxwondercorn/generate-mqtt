// config is a central location for configuration options. It also contains
// config file parsing logic.

// credit to nanopack
// https://github.com/nanopack/portal/blob/d165bff1df343d390432e078a69f715667ad004a/config/config.go

package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

// Defaults if not defined in scanner.toml
var (
	BrokerName = "testing123"
	BrokerHost = "127.0.0.1"
	BrokerPort = 1883
	BrokerUser = ""
	BrokerPwd  = ""
	ConfigFile *string
	Delay  = 1000
	RootTag    = ""
)

func init() {
	// set defaults if missing in TOML file or commandline
	viper.SetDefault("broker.name", BrokerName)
	viper.SetDefault("broker.host", BrokerHost)
	viper.SetDefault("broker.port", BrokerPort)
	viper.SetDefault("broker.user", BrokerUser)
	viper.SetDefault("broker.password", BrokerPwd)

	viper.SetDefault("mqtt.root_tag", RootTag)
	viper.SetDefault("scanner.delay", Delay)
}

func LoadConfigFile() error {

	ConfigFile = flag.String("config", "./generator.toml", "The path to the onfiguration file")
	flag.Parse()

	viper.SetConfigFile(*ConfigFile)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}

	// Set values - config file will override commandline
	BrokerName = viper.GetString("broker.name")
	BrokerHost = viper.GetString("broker.host")
	BrokerPort = viper.GetInt("broker.port")
	BrokerUser = viper.GetString("broker.user")
	BrokerPwd = viper.GetString("broker.password")

	RootTag = viper.GetString("mqtt.root_tag")

	Delay = viper.GetInt("scanner.delay")

	return nil
}
