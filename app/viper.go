package app

import (
	"bytes"
	"os"
	"strings"

	"github.com/raven0520/btc/util"
	"github.com/spf13/viper"
)

// IgnoreFileList File names that do not need to be parsed
var IgnoreFileList = []string{
	".DS_Store",
}

// InitViperConfig 初始化配置文件
func InitViperConfig() error {
	file, err := os.Open(util.ConfigPath + "/")
	if err != nil {
		return err
	}
	fileList, err := file.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f := range fileList {
		// Ignore other file
		if util.InSliceString(f.Name(), IgnoreFileList) {
			continue
		}
		if !f.IsDir() {
			bts, err := os.ReadFile(util.ConfigPath + "/" + f.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			v.SetConfigType("toml")
			err = v.ReadConfig(bytes.NewBuffer(bts))
			if err != nil {
				return err
			}
			pathArr := strings.Split(f.Name(), ".")
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper)
			}
			ViperConfMap[pathArr[0]] = v
		}
	}
	return nil
}
