package core

import (
	"encoding/hex"
	"strings"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/pkg/errors"
	"github.com/raven0520/btc/bip44"
)

const symbol = "BTC"

// BTC key derivation service
type BTC struct {
	CoinInfo
	useSegWit bool //是否使用隔离见证地址
}

// New Factory of BTC key derivation service
//
// The order of publicKeys is important.
// using segWit will replace m/44'... ==> m/49'
func New(bip44Path string, isSegWit bool, seed []byte, chainID int) (c *BTC, err error) {
	c = new(BTC)

	c.Symbol = symbol
	c.useSegWit = isSegWit
	if isSegWit {
		bip44Path = strings.Replace(bip44Path, "m/44'", "m/49'", 1)
	}
	c.DerivationPath, err = bip44.GetCoinDerivationPath(bip44Path, c.Symbol)
	if err != nil {
		return nil, errors.Wrap(err, "bip44.GetCoinDerivationPath err")
	}
	c.ChainCfg, err = ChainFlag2ChainParams(chainID)
	if err != nil {
		return nil, err
	}
	c.MasterKey, err = hdkeychain.NewMaster(seed, c.ChainCfg)
	if err != nil {
		err = errors.Wrap(err, "hdkeychain.NewMaster")
		return
	}
	return c, nil
}

// deriveChildKey derives the child key of the derivation path.
func (c *BTC) deriveChildKey() (childKey *hdkeychain.ExtendedKey, err error) {
	childKey = c.MasterKey
	childKey.ChildIndex()
	for _, childOpt := range c.DerivationPath {
		childKey, err = childKey.Derive(childOpt)
		if err != nil {
			err = errors.Wrapf(err, "childKey.Child for %x", childOpt)
			return
		}
	}
	return
}

// derivePrivateKey derives the private key of the derivation path.
func (c *BTC) derivePrivateKey() (prikey *btcec.PrivateKey, err error) {
	childKey, err := c.deriveChildKey()
	if err != nil {
		err = errors.Wrap(err, "c.deriveChildKey")
		return
	}
	prikey, err = childKey.ECPrivKey()
	if err != nil {
		err = errors.Wrap(err, "childKey.ECPrivKey")
		return
	}

	return
}

// GetDerivationPath Return Derivation Path
func (c *BTC) GetDerivationPath() accounts.DerivationPath {
	return c.DerivationPath
}

// DerivePrivateKey derives the private key of the derivation path, encoded in string with WIF format
func (c *BTC) DerivePrivateKey() (privateKey string, err error) {
	prikey, err := c.derivePrivateKey()
	if err != nil {
		return
	}
	wif, err := btcutil.NewWIF(prikey, c.ChainCfg, true)
	if err != nil {
		return
	}
	privateKey = wif.String()
	return
}

// DerivePublicKey derives the public key of the derivation path.
func (c *BTC) DerivePublicKey() (publicKey string, err error) {
	prikey, err := c.derivePrivateKey()
	if err != nil {
		return
	}
	publicKey = hex.EncodeToString(prikey.PubKey().SerializeCompressed())
	return
}

// DeriveAddress derives the account address of the derivation path.
func (c *BTC) DeriveAddress() (address string, err error) {
	var addrP2PKH *btcutil.AddressPubKeyHash
	childKey, err := c.deriveChildKey()
	if err != nil {
		err = errors.Wrap(err, "c.deriveChildKey")
		return
	}
	if !c.useSegWit {
		addrP2PKH, err = childKey.Address(c.ChainCfg)
		if err != nil {
			err = errors.Wrap(err, "childKey.Address")
			return
		}
		address = addrP2PKH.String()
	} else {
		pubk, er := childKey.ECPubKey()
		if er != nil {
			return "", errors.Wrap(er, "childKey.ECPubKey err")
		}
		address, er = ConvertPubk2segWitP2WSHAddress(pubk, c.ChainCfg)
		if er != nil {
			return "", errors.Wrap(er, "convert segWit address err")
		}
	}
	return
}
