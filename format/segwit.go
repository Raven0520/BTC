package format

import (
	"github.com/raven0520/btc/bip44"
	"github.com/raven0520/btc/util"
	"github.com/raven0520/proto/btc"
)

// SegwitResponse Return Segwit
func SegwitResponse(deriver bip44.Deriver) (response *btc.SegwitResponse, err error) {
	var address, prvk, pubk string
	if address, err = deriver.DeriveAddress(); err != nil {
		return
	}
	if prvk, err = deriver.DerivePrivateKey(); err != nil {
		return
	}
	if pubk, err = deriver.DerivePublicKey(); err != nil {
		return
	}
	response = &btc.SegwitResponse{
		Message: Message(util.RequestOK, err),
		Data: &btc.Segwit{
			Address: address,
			Private: prvk,
			Public:  pubk,
		},
	}
	return
}
