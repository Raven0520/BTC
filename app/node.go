package app

import (
	"fmt"

	"github.com/raven0520/btc/util"
)

// BTCNodeMap
var BTCNodeMap = map[string]*BTCNodeConfig{}

// BTCNodeList Valid BTCNodes
type BTCNodeList struct {
	Node map[string]*BTCNodeConfig `mapstructure:"node"`
}

// BTCNodeConfig BTCNode configture
type BTCNodeConfig struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Debug bool   `json:"debug"`
}

// InitBTCNodeConfig
func InitBTCNodeConfig(path string) error {
	BTCNodeList := &BTCNodeList{}
	err := util.ParseConfig(path, &BTCNodeList)
	if err != nil {
		return err
	}
	if len(BTCNodeList.Node) == 0 {
		fmt.Printf("[INFO] %s\n", " empty BTCNode config.")
	}
	for name, v := range BTCNodeList.Node {
		BTCNodeMap[name] = &BTCNodeConfig{
			Host:  v.Host,
			Port:  v.Port,
			User:  v.User,
			Pass:  v.Pass,
			Debug: v.Debug,
		}
	}
	return nil
}

// GetBTCNodeConfig get BTCNode configture
func GetBTCNodeConfig(name string) (BTCNode *BTCNodeConfig, ok bool) {
	BTCNode, ok = BTCNodeMap[name]
	return
}
