package core

import (
	"log"
	"testing"

	"github.com/raven0520/btc/bip39"
	"github.com/raven0520/btc/bip44"
	"github.com/stretchr/testify/assert"
)

var (
	testMnemonic = "exit fruit duty weekend romance upper human before nuclear rabbit slim frame"
	btc          = &BTC{}
)

func init() {
	var err error
	seed, err := bip39.NewSeedWithErrorChecking(testMnemonic, "")
	if err != nil {
		log.Fatal(err)
	}
	btc, err = New(bip44.FullPathFormat, true, seed, ChainRegtest)
	if err != nil {
		log.Fatal(err)
	}
}

func TestBTC_DerivePrivateKey(t *testing.T) {
	pk, err := btc.DerivePrivateKey()
	assert.NoError(t, err)
	assert.Equal(t, "cW3yMs74DrwsFyyWpnubVJXPo7ptrFT7hi6rE92gHDGtbSrXBbpc", pk)
}

func TestBTC_DerivePublicKey(t *testing.T) {
	pub, err := btc.DerivePublicKey()
	assert.NoError(t, err)
	assert.Equal(t, "03ff5fa11a73a5b0147fdd8c837ca00665f568de083ee0c8f2d0518bcfb1970e2e", pub)
}

func TestBTC_DeriveAddress(t *testing.T) {
	address, err := btc.DeriveAddress()
	assert.NoError(t, err)
	assert.Equal(t, "2N7iuhPBAdD4Zyhj7ZjACY7uJWmaZgBu5ZX", address)
}
