package service

import (
	"context"

	"github.com/raven0520/btc/format"
	"github.com/raven0520/btc/logic"
	"github.com/raven0520/btc/util"
	"github.com/raven0520/proto/btc"
)

// BtcService BTC Service
type BtcService struct {
	btc.UnimplementedBTCServiceServer
	// pb.UnimplementedBTCServiceServer
}

// NewSegwit Generate Segwit
func (bs *BtcService) NewSegwit(_ context.Context, post *btc.EmptyPost) (*btc.SegwitResponse, error) {
	return logic.NewSegwit()
}

// GenerateSegwit Generate Segwit From Post Mnemonic
func (bs *BtcService) SegwitFromMnemonic(_ context.Context, post *btc.MnemonicPost) (*btc.SegwitResponse, error) {
	return logic.GenerateSegwit(post)
}

// SegwitFromSeed Generate Segwit From Post Seed
func (bs *BtcService) SegwitFromSeed(_ context.Context, post *btc.SeedPost) (*btc.SegwitResponse, error) {
	return logic.GenerateSegwitFromSeed(post)
}

// MultiSig Generate Multiple signatures
func (bs *BtcService) MultiSig(_ context.Context, post *btc.MultiSigPost) (*btc.MultiSigResponse, error) {
	address, script, err := logic.NewMultiSigAddress(post)
	return &btc.MultiSigResponse{
		Message: format.Message(util.RequestOK, err),
		Data: &btc.MultiSig{
			Address: address,
			Script:  script,
		},
	}, err
}
