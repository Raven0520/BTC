package format

import (
	"github.com/raven0520/btc/bip44"
	"github.com/raven0520/btc/proto"
	"github.com/raven0520/btc/util"
)

// SegwitResponse Return Segwit
func SegwitResponse(deriver bip44.Deriver) (response *proto.SegwitResponse, err error) {
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
	response = &proto.SegwitResponse{
		Message: Message(util.RequestOK, err),
		Data: &proto.Segwit{
			Address: address,
			Private: prvk,
			Public:  pubk,
		},
	}
	return
}
