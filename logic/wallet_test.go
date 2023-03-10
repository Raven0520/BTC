package logic

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/pkg/errors"
	"github.com/raven0520/btc/bip39"
	"github.com/raven0520/btc/proto"
	"github.com/raven0520/btc/util"
	"github.com/stretchr/testify/assert"
)

func TestNewSegwit(t *testing.T) {
	post := &proto.NewSegwitPost{ChainID: 2}
	// Error Wordlist
	wordlist := gomonkey.ApplyFunc(bip39.SetWordListLang, func(lang int) error {
		return errors.Errorf("expected lang [%d - %d], got: %d", bip39.LangChineseSimplified, bip39.LangSpanish, 9)
	})
	response, err := NewSegwit(post)
	assert.Nil(t, response)
	assert.EqualError(t, err, "expected lang [0 - 7], got: 9")
	wordlist.Reset()
	// Success Case
	monkey := gomonkey.ApplyFunc(bip39.NewMnemonic, func(entropy []byte) (string, error) {
		return "exit fruit duty weekend romance upper human before nuclear rabbit slim frame", nil
	})
	defer monkey.Reset()
	external := &proto.SegwitResponse{
		Message: util.RequestOK,
		Data: &proto.Segwit{
			Address: "2N7iuhPBAdD4Zyhj7ZjACY7uJWmaZgBu5ZX",
			Private: "cW3yMs74DrwsFyyWpnubVJXPo7ptrFT7hi6rE92gHDGtbSrXBbpc",
			Public:  "03ff5fa11a73a5b0147fdd8c837ca00665f568de083ee0c8f2d0518bcfb1970e2e",
		},
	}
	response, err = NewSegwit(post)
	assert.NoError(t, err)
	assert.EqualValues(t, external, response)
}

func TestGenerateSegwit(t *testing.T) {
	// path "m/49'/1'/0'/0/0"
	post := &proto.MnemonicPost{
		ChainID:  2,
		Mnemonic: "exit fruit duty weekend romance upper human before nuclear rabbit slim frame",
		Pass:     "",
		Account:  0,
		External: true,
		Address:  0,
	}
	t.Run("external", func(t *testing.T) {
		external := &proto.SegwitResponse{
			Message: util.RequestOK,
			Data: &proto.Segwit{
				Address: "2N7iuhPBAdD4Zyhj7ZjACY7uJWmaZgBu5ZX",
				Private: "cW3yMs74DrwsFyyWpnubVJXPo7ptrFT7hi6rE92gHDGtbSrXBbpc",
				Public:  "03ff5fa11a73a5b0147fdd8c837ca00665f568de083ee0c8f2d0518bcfb1970e2e",
			},
		}
		response, err := GenerateSegwit(post)
		assert.NoError(t, err)
		assert.EqualValues(t, external, response)
	})
	// path "m/49'/1'/0'/1/0"
	t.Run("change", func(t *testing.T) {
		post.External = false
		change := &proto.SegwitResponse{
			Message: util.RequestOK,
			Data: &proto.Segwit{
				Address: "2N34mTbwU6PwyhtGFQy8iML9fq9C3qgCVDE",
				Private: "cPppReeEVsy9V6TPyXDUjERFMupHqeEcCF1EQJFNrbkjQ62vues9",
				Public:  "039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda",
			},
		}
		response, err := GenerateSegwit(post)
		assert.NoError(t, err)
		assert.EqualValues(t, change, response)
	})
}

func TestGenerateSegwitFromSeed(t *testing.T) {
	type args struct {
		post *proto.SeedPost
	}
	tests := []struct {
		name         string
		args         args
		wantResponse *proto.SegwitResponse
		wantErr      bool
	}{
		{
			name: "External",
			args: args{
				post: &proto.SeedPost{
					ChainID:  2,
					Seed:     "a1b70d538307815c03ac6ba0668564257871e4ec68ab1412d09d5f4d34a9ee73ea770681684cb60973c17761347b52d2336e539310154dd75ec0f0e9a7ab4b3f",
					Account:  0,
					External: true,
					Address:  0,
				},
			},
			wantResponse: &proto.SegwitResponse{
				Message: util.RequestOK,
				Data: &proto.Segwit{
					Address: "2N7iuhPBAdD4Zyhj7ZjACY7uJWmaZgBu5ZX",
					Private: "cW3yMs74DrwsFyyWpnubVJXPo7ptrFT7hi6rE92gHDGtbSrXBbpc",
					Public:  "03ff5fa11a73a5b0147fdd8c837ca00665f568de083ee0c8f2d0518bcfb1970e2e",
				},
			},
			wantErr: false,
		},
		{
			name: "Change",
			args: args{
				post: &proto.SeedPost{
					ChainID:  2,
					Seed:     "a1b70d538307815c03ac6ba0668564257871e4ec68ab1412d09d5f4d34a9ee73ea770681684cb60973c17761347b52d2336e539310154dd75ec0f0e9a7ab4b3f",
					Account:  0,
					External: false,
					Address:  0,
				},
			},
			wantResponse: &proto.SegwitResponse{
				Message: util.RequestOK,
				Data: &proto.Segwit{
					Address: "2N34mTbwU6PwyhtGFQy8iML9fq9C3qgCVDE",
					Private: "cPppReeEVsy9V6TPyXDUjERFMupHqeEcCF1EQJFNrbkjQ62vues9",
					Public:  "039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := GenerateSegwitFromSeed(tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSegwitFromSeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GenerateSegwitFromSeed() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestNewMultiSigAddress(t *testing.T) {
	type args struct {
		post *proto.MultiSigPost
	}
	tests := []struct {
		name        string
		args        args
		wantAddress string
		wantScript  string
		wantErr     bool
	}{
		{
			name:        "Get Chain Params Error",
			args:        args{post: &proto.MultiSigPost{ChainID: 4, Required: 0, PublicKeys: []string{}}},
			wantAddress: "",
			wantScript:  "",
			wantErr:     true,
		},
		{
			name:        "Need more public key",
			args:        args{post: &proto.MultiSigPost{ChainID: 2, Required: 2, PublicKeys: []string{"039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda"}}},
			wantAddress: "",
			wantScript:  "",
			wantErr:     true,
		},
		{
			name:        "Success",
			args:        args{post: &proto.MultiSigPost{ChainID: 2, Required: 2, PublicKeys: []string{"039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda", "03ff5fa11a73a5b0147fdd8c837ca00665f568de083ee0c8f2d0518bcfb1970e2e"}}},
			wantAddress: "2N5tLUpDaPC5bh1dXGHrLvhXKCtVyF2ppXV",
			wantScript:  "5221039b51299768241c89ae9958eeabcb27f11ababbfa33c240f2495ef11b7ce0acda2103ff5fa11a73a5b0147fdd8c837ca00665f568de083ee0c8f2d0518bcfb1970e2e52ae",
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddress, gotScript, err := NewMultiSigAddress(tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMultiSigAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAddress != tt.wantAddress {
				t.Errorf("NewMultiSigAddress() gotAddress = %v, want %v", gotAddress, tt.wantAddress)
			}
			if gotScript != tt.wantScript {
				t.Errorf("NewMultiSigAddress() gotScript = %v, want %v", gotScript, tt.wantScript)
			}
		})
	}
}
