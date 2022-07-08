package api

import (
	"backend-challenge-go/model"
	"backend-challenge-go/util"
	"reflect"
	"sync"
	"testing"
)

var tokenInfo = sync.Map{}

func loadTokenData() {
	tokenInfo.Store("1", &model.Token{
		Name:        "Ethereum",
		Symbol:      "ETH",
		Address:     "0xEeeeeeeeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		Decimals:    18,
		TotalSupply: util.BigIntFromInt(1000000000000000000),
	})
	tokenInfo.Store("2", &model.Token{
		Name:        "Bitcoin",
		Symbol:      "BTC",
		Address:     "0xBitcoin",
		Decimals:    8,
		TotalSupply: util.BigIntFromInt(21000000000),
	})
	tokenInfo.Store("3", &model.Token{
		Name:        "Ethereum Classic",
		Symbol:      "ETC",
		Address:     "0xEeeeeeeeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		Decimals:    18,
		TotalSupply: util.BigIntFromInt(1000000000000000000),
	})
	tokenInfo.Store("4", &model.Token{
		Name:        "Litecoin",
		Symbol:      "LTC",
		Address:     "0xLitecoin",
		Decimals:    8,
		TotalSupply: util.BigIntFromInt(21000000000),
	})
	tokenInfo.Store("5", &model.Token{
		Name:        "Bitcoin Cash",
		Symbol:      "BCH",
		Address:     "0xBitcoinCash",
		Decimals:    8,
		TotalSupply: util.BigIntFromInt(21000000000),
	})
}

func Test_createTokenResp(t *testing.T) {

	loadTokenData()

	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want model.TokenSlice
	}{
		// TODO: Add test cases.
		{
			name: "Should return BTC token",
			args: args{
				q: "btc",
			},
			want: model.TokenSlice{
				Tokens: []model.Token{
					{
						Name:        "Bitcoin",
						Symbol:      "BTC",
						Address:     "0xBitcoin",
						Decimals:    8,
						TotalSupply: util.BigIntFromInt(21000000000),
					},
				},
			},
		}, {
			name: "Should return ETH and ETC token",
			args: args{
				q: "eth",
			},
			want: model.TokenSlice{
				Tokens: []model.Token{
					{
						Name:        "Ethereum",
						Symbol:      "ETH",
						Address:     "0xEeeeeeeeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
						Decimals:    18,
						TotalSupply: util.BigIntFromInt(1000000000000000000),
					}, {
						Name:        "Ethereum Classic",
						Symbol:      "ETC",
						Address:     "0xEeeeeeeeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
						Decimals:    18,
						TotalSupply: util.BigIntFromInt(1000000000000000000),
					},
				},
			},
		}, {
			name: "Should return BTC, LTC, and BCH tokens",
			args: args{
				q: "coin",
			},
			want: model.TokenSlice{
				Tokens: []model.Token{
					{
						Name:        "Bitcoin",
						Symbol:      "BTC",
						Address:     "0xBitcoin",
						Decimals:    8,
						TotalSupply: util.BigIntFromInt(21000000000),
					}, {
						Name:        "Litecoin",
						Symbol:      "LTC",
						Address:     "0xLitecoin",
						Decimals:    8,
						TotalSupply: util.BigIntFromInt(21000000000),
					}, {
						Name:        "Bitcoin Cash",
						Symbol:      "BCH",
						Address:     "0xBitcoinCash",
						Decimals:    8,
						TotalSupply: util.BigIntFromInt(21000000000),
					},
				},
			},
		}, {
			name: "Should return empty slice",
			args: args{
				q: "abcdefg",
			},
			want: model.TokenSlice{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTokenResp(tokenInfo, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTokenResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
