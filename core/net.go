package core

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ChainMainNet
// ChainTestNet3
// ChainRegtest
const (
	ChainMainNet = iota
	ChainTestNet3
	ChainRegtest
	// FlagBTCUseSegWitFormat BTC Segwit Format
	FlagUseSegWitFormat = "btc_use_segwit_fmt"
)

// ChainFlag2ChainParams get chainParams from const
func ChainFlag2ChainParams(chainID int) (*chaincfg.Params, error) {
	switch chainID {
	case ChainMainNet:
		return &chaincfg.MainNetParams, nil
	case ChainTestNet3:
		return &chaincfg.TestNet3Params, nil
	case ChainRegtest:
		return &chaincfg.RegressionNetParams, nil
	default:
		return nil, fmt.Errorf("Expected chain options: %d > MainNet , %d > TestNet3, %d > Regtest, got ：%d", ChainMainNet, ChainTestNet3, ChainRegtest, chainID)
	}
}
