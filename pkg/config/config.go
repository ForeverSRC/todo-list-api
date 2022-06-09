package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func Init() {
	newConfig("./pkg/conf", "")
}

func newConfig(path string, mode string) {
	// Default
	Config = viper.New()
	Config.AddConfigPath(path)
	Config.Set("path", path)
	if mode != "" {
		Config.SetConfigName(mode)
		Config.Set("mode", mode)
		err := Config.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("failed to read config: %v", err))

		}
	} else {
		Config.SetConfigName("default")
		err := Config.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("failed to read config: %v", err))
		}
		Config.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

		// Merge according to RUN_MODE
		Config.SetEnvPrefix("run")
		Config.AutomaticEnv()
		mode = Config.GetString("mode")
		if len(mode) > 0 {
			Config.SetConfigName(mode)
			err = Config.MergeInConfig()
			if err != nil {
				panic(fmt.Errorf("failed to merge config: %v", err))
			}
		}
	}
}
