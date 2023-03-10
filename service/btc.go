package service

import (
	"context"

	"github.com/raven0520/btc/format"
	"github.com/raven0520/btc/logic"
	"github.com/raven0520/btc/proto"
	"github.com/raven0520/btc/util"
)

// BtcService BTC Service
type BtcService struct {
	proto.UnimplementedBTCServiceServer
	// pb.UnimplementedBTCServiceServer
}

// NewSegwit Generate Segwit
func (bs *BtcService) NewSegwit(_ context.Context, post *proto.NewSegwitPost) (*proto.SegwitResponse, error) {
	return logic.NewSegwit(post)
}

// GenerateSegwit Generate Segwit From Post Mnemonic
func (bs *BtcService) SegwitFromMnemonic(_ context.Context, post *proto.MnemonicPost) (*proto.SegwitResponse, error) {
	return logic.GenerateSegwit(post)
}

// SegwitFromSeed Generate Segwit From Post Seed
func (bs *BtcService) SegwitFromSeed(_ context.Context, post *proto.SeedPost) (*proto.SegwitResponse, error) {
	return logic.GenerateSegwitFromSeed(post)
}

// MultiSig Generate Multiple signatures
func (bs *BtcService) MultiSig(_ context.Context, post *proto.MultiSigPost) (*proto.MultiSigResponse, error) {
	address, script, err := logic.NewMultiSigAddress(post)
	return &proto.MultiSigResponse{
		Message: format.Message(util.RequestOK, err),
		Data: &proto.MultiSig{
			Address: address,
			Script:  script,
		},
	}, err
}
