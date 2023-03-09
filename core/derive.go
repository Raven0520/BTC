package core

import "github.com/raven0520/btc/bip44"

// NewBip44Deriver btc bip44 实现
func NewBip44Deriver(bip44Path string, isSegWit bool, seed []byte, chainID int) (bip44.Deriver, error) {
	coin, err := New(bip44Path, isSegWit, seed, chainID)
	return coin, err
}
