package bip44

import "testing"

func TestGetCoinType(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name         string
		args         args
		wantCoinType uint32
		wantErr      bool
	}{
		{
			name:         "BTC",
			args:         args{symbol: "BTC"},
			wantCoinType: 0x00,
			wantErr:      false,
		},
		{
			name:         "ETH",
			args:         args{symbol: "ETH"},
			wantCoinType: 0x3c,
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCoinType, err := GetCoinType(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCoinType != tt.wantCoinType {
				t.Errorf("GetCoinType() = %v, want %v", gotCoinType, tt.wantCoinType)
			}
		})
	}
}
