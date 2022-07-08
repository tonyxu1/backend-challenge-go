package api

import (
	"backend-challenge-go/model"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
)

// GetTokens responds with a list of tokens, this function is used by the /tokens endpoint
// it doesn't need to be exported because it is only used by the /tokens endpoint, just in case.
// Extend this function to add more tokens to the list, expected comma separated list of tokens info
// or contract addresses beginning with '0x'
// example: ?q=bit,eth, 0xdbf1344a0ff21bc098eb9ad4eef7de0f9722c02b
// if empty tokens returns, please make sure the contract address is in addresses.toml file.
func GetTokens(c echo.Context) error {
	tokenParam := c.QueryParam("q")
	if tokenParam == "" {
		return c.String(500, "No token info provided, please provide a token info by using the query param 'q'")
	}

	resp := createTokenResp(model.TokenInfo, tokenParam)
	return c.JSON(200, resp)
}

// Build response for GetTokens.
func createTokenResp(tokenMap sync.Map, q string) model.TokenSlice {
	var ts model.TokenSlice
	for _, qeryParam := range strings.Split(q, ",") {
		qeryParam = strings.TrimSpace(strings.ToLower(qeryParam))
		if strings.HasPrefix(qeryParam, "0x") {
			tokenMap.Range(func(key, value interface{}) bool {
				v := value.(*model.Token)

				if strings.ToLower(v.Address) == qeryParam {

					ts.Tokens = append(ts.Tokens, *v)
					return false
				}
				return true
			})
		} else {
			tokenMap.Range(func(key, value interface{}) bool {
				v := value.(*model.Token)
				name := strings.ToLower(v.Name)
				symbol := strings.ToLower(v.Symbol)

				if strings.Contains(name, qeryParam) || strings.Contains(symbol, qeryParam) {

					ts.Tokens = append(ts.Tokens, *v)
				}
				return true
			})
		}

	}
	return ts
}
