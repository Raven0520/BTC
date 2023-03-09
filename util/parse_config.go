package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var ConfigPath string // Config path
var Env string        // Environment name

// GetConfigPath 获取配置文件路径
func GetConfigPath(fileName string) string {
	return ConfigPath + "/" + fileName + ".toml"
}

// ParseConfigPath Parse config path
func ParseConfigPath(configPath string) error {
	path := strings.Split(configPath, "/")
	prefix := strings.Join(path[:len(path)-1], "/")
	ConfigPath = prefix
	Env = path[len(path)-2]
	return nil
}

// ParseConfig Parse Configture
func ParseConfig(path string, config interface{}) error {
	file, err := os.Open(path) // Open config file
	if err != nil {
		return fmt.Errorf("open config file %v failed: %v ", path, err)
	}
	data, err := io.ReadAll(file) // Read config file
	if err != nil {
		return fmt.Errorf("read config %v failed: %v ", path, err)
	}
	v := viper.New() // Read config file with Viper
	v.SetConfigType("toml")
	if err := v.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return fmt.Errorf("viper read config faild, config: %v, err: %v ", string(data), err)
	}
	if err := v.Unmarshal(config); err != nil {
		return fmt.Errorf("viper Parse config faild, config: %v, err: %v ", string(data), err)
	}
	return nil
}
