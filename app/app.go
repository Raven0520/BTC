package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/raven0520/btc/util"
)

var TimeLocation *time.Location
var InitProgress = []string{
	"Start Loading Resources ------------------------------------------------",
	"[INFO]  Config Path : %s \n",
}

// Init Application
func Init(configPath string) error {
	return InitModule(configPath, []string{"base", "node"})
}

// InitModule Init modules
func InitModule(configPath string, modules []string) error {
	config := flag.String("config", configPath, "input config file like ./config/develop/") // Output config path format
	testing.Init()
	flag.Parse()
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Println("Start Loading Resources ------------------------------------------------")
	log.Printf("[INFO]  Config Path : %s \n", *config) // Output config path
	util.SetLocalIPs()
	// Parse Config Path
	if err := util.ParseConfigPath(*config); err != nil {
		return err
	}
	log.Printf("[INFO] %s\n", " Parse Config Path Done.")

	if err := InitViperConfig(); err != nil {
		return err
	}
	log.Printf("[INFO] %s\n", " Viper Config Done.")

	// Load Base Configture
	if util.InSliceString("base", modules) {
		if err := InitBaseConfig(util.GetConfigPath("base")); err != nil {
			fmt.Printf("[ERROR] InitBaseConfig: %s\n", err.Error())
		}
		log.Printf("[INFO] %s\n", " Base Config Done.")
	}
	// Set Time Zone
	location, err := time.LoadLocation(BaseConf.Base.TimeLocation)
	if err != nil {
		return err
	}
	// Load Node Configture
	if util.InSliceString("node", modules) {
		if err := InitBTCNodeConfig(util.GetConfigPath("node")); err != nil {
			fmt.Printf("[ERROR] InitNodeConfig: %s\n", err.Error())
		}
		log.Printf("[INFO] %s\n", " Node Config Done.")
	}
	TimeLocation = location
	log.Println("--------------------------------------------- Loading Resources Success ")
	return nil
}
