package app

import "github.com/raven0520/btc/util"

// BaseConf 全局变量
var BaseConf *BaseConfig

// BaseConfig Base Config Type
type BaseConfig struct {
	Base   Base         `mapstructure:"base"`
	Path   Path         `mapstructure:"path"`
	Consul ConsulConfig `mapstructure:"consul"`
	Grpc   GrpcConfig   `mapstructure:"grpc"`
	Http   HttpConfig   `mapstructure:"http"`
}

// BaseConfig Base Configture
type Base struct {
	Env          string `mapstructure:"env"`
	Net          string `mapstructure:"net"`
	DebugMode    string `mapstructure:"debug_mode"`
	TimeLocation string `mapstructure:"time_location"`
}

// Path File path
type Path struct {
	Pid string `mapstructure:"pid"`
}

// ConsulConfig Consul Configture
type ConsulConfig struct {
	Env  string `mapstructure:"env"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// GrpcConfig GRPC Configture
type GrpcConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// HttpConfig HTTP Configture
type HttpConfig struct {
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
	ReadTimeout    int    `mapstructure:"read_timeout"`
	WriteTimeout   int    `mapstructure:"write_timeout"`
	MaxHeaderBytes int    `mapstructure:"max_header_bytes"`
}

// InitBaseConfig Initialize the basic settings
func InitBaseConfig(path string) error {
	BaseConf = &BaseConfig{}
	if err := util.ParseConfig(path, BaseConf); err != nil {
		return err
	}
	BaseConf.Base.Env = GetEnv()
	BaseConf.Base.DebugMode = GetDebugMode()
	BaseConf.Base.TimeLocation = GetTimeLocation()
	return nil
}

// GetEnv Get Env
func GetEnv() string {
	Env := "Dev" // 默认环境
	if BaseConf.Base.Env != "" {
		return BaseConf.Base.Env
	}
	return Env
}

// GetDebugMode Debug Mode
func GetDebugMode() string {
	Mode := "debug"
	if BaseConf.Base.DebugMode != "" {
		return BaseConf.Base.DebugMode
	}
	return Mode
}

// GetTimeLocation TimeLocation
func GetTimeLocation() string {
	TimeLocation := "Asia/Shanghai"
	if BaseConf.Base.TimeLocation == "" {
		return BaseConf.Base.TimeLocation
	}
	return TimeLocation
}
