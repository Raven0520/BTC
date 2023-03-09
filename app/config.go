package app

import (
	"strings"

	"github.com/spf13/viper"
)

var ViperConfMap map[string]*viper.Viper // Configture Map

// GetStringConfig get config value of string format
func GetStringConfig(key string) string {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return ""
	}
	v, ok := ViperConfMap[keys[0]]
	if !ok {
		return ""
	}
	confString := v.GetString(strings.Join(keys[1:], "."))
	return confString
}

// GetIntConfig get config value of int format
func GetIntConfig(key string) int {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return 0
	}
	v := ViperConfMap[keys[0]]
	conf := v.GetInt(strings.Join(keys[1:], "."))
	return conf
}
