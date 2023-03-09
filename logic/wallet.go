package logic

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/pkg/errors"
	"github.com/raven0520/btc/app"
	"github.com/raven0520/btc/bip39"
	"github.com/raven0520/btc/bip44"
	"github.com/raven0520/btc/core"
	"github.com/raven0520/btc/format"
	"github.com/raven0520/proto/btc"
)

// NewSegwit Generate Segwit From Random Mnemonic
func NewSegwit() (response *btc.SegwitResponse, err error) {
	var (
		entropy, seed []byte
		mnemonic      string
		deriver       bip44.Deriver
	)
	chainID := ChainNet[app.BaseConf.Base.Net]
	err = bip39.SetWordListLang(bip39.LangEnglish)
	if err != nil {
		return
	}
	if entropy, err = bip39.NewEntropy(128); err != nil {
		return
	}
	if mnemonic, err = bip39.NewMnemonic(entropy); err != nil {
		return
	}
	if seed, err = bip39.NewSeedWithErrorChecking(mnemonic, ""); err != nil {
		return
	}
	if deriver, err = core.NewBip44Deriver(bip44.FullPathFormat, true, seed, chainID); err != nil {
		return
	}
	return format.SegwitResponse(deriver)
}

// GenerateSegwit Generate Segwit From Post Mnemonic
func GenerateSegwit(post *btc.MnemonicPost) (response *btc.SegwitResponse, err error) {
	var (
		external = 0
		seed     []byte
		pathf    string
		deriver  bip44.Deriver
	)
	if !post.External {
		external = 1
	}
	chainID := ChainNet[app.BaseConf.Base.Net]
	if seed, err = bip39.NewSeedWithErrorChecking(post.Mnemonic, post.Pass); err != nil {
		return
	}
	pathf = fmt.Sprintf("m/44'/%s'/%d'/%d/%d", "%d", post.Account, external, post.Address)
	if deriver, err = core.NewBip44Deriver(pathf, true, seed, chainID); err != nil {
		return
	}
	return format.SegwitResponse(deriver)
}

// GenerateSegwit Generate Segwit From Seed & Path
func GenerateSegwitFromSeed(post *btc.SeedPost) (response *btc.SegwitResponse, err error) {
	var (
		external = 0
		pathf    string
		deriver  bip44.Deriver
	)
	chainID := ChainNet[app.BaseConf.Base.Net]
	seed, err := hex.DecodeString(post.Seed)
	if err != nil {
		return
	}
	if !post.External {
		external = 1
	}
	pathf = fmt.Sprintf("m/44'/%s'/%d'/%d/%d", "%d", post.Account, external, post.Address)
	if deriver, err = core.NewBip44Deriver(pathf, true, seed, chainID); err != nil {
		return
	}
	return format.SegwitResponse(deriver)
}

// Params:
//	chainID: 0 Mainet 1 TestNet3 2 Regtest
//	cmd.NRequired num required to sign
//	cmd.Keys Hex-encoded public key
// Verify：len(cmd.Keys) >= cmd.NRequired

//	Return:
//	Address, RedeemScript

// NewMultiSigAddress Generate MultiSigAddress
func NewMultiSigAddress(post *btc.MultiSigPost) (address, script string, err error) {
	chainParams, err := core.ChainFlag2ChainParams(int(post.ChainID))
	if err != nil {
		return
	}
	if int(post.Required) > len(post.PublicKeys) {
		return "", "", errors.New("Need more public key")
	}
	rs, err := core.CreateMultiSig(btcjson.NewCreateMultisigCmd(int(post.Required), post.PublicKeys), chainParams)
	if err != nil {
		return
	}
	return rs.Address, rs.RedeemScript, nil
}
