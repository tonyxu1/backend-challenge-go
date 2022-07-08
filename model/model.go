package model

import (
	"math/big"
	"sync"
)

//Token represents a metadata of a token
type Token struct {
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Address     string   `json:"address"`
	Decimals    int      `json:"decimals"`
	TotalSupply *big.Int `json:"totalSupply"`
}

type TokenSlice struct {
	Tokens []Token `json:"tokens"`
}

var (
	//TokenInfo is a map of token info, this map needs to be thread safe
	TokenInfo = sync.Map{}
)
