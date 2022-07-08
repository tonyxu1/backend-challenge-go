package main

import (
	"backend-challenge-go/model"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

//TODO: Make the interval & number of chunks configurable
const (
	interval    = 20 * time.Minute
	totalChunks = 20
	apiKey      = "RSUpmPce42zYBEHER3XN0B_VOCCI-Esl"
	methodName  = "getContractMetadata"
)

var (
	// Addresses is a map of chunk of address from addresses.jsonl
	addresses   = make([][]string, 0)
	addressFile = "./data/addresses.jsonl"

	flagRPCURL = flag.String("rpc-url", "https://eth-mainnet.alchemyapi.io/nft/v2", "RPC URL")

	fullUrl = fmt.Sprintf("%s/%s/%s?contractAddress=", *flagRPCURL, apiKey, methodName)
)

type contract struct {
	Address          string `json:"address"`
	ContractMetadata struct {
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		TotalSupply string `json:"totalSupply"`
		TokenType   string `json:"tokenType"`
	} `json:"contractMetadata"`
}

//Start cron job to get token information from ethereum by alchemy api
func StartCronJob() {
	populateAddressData()
	populateTokenData()
	go func() {
		for range time.Tick(interval) {
			populateAddressData()
			populateTokenData()
		}
	}()
}

// populateTokenData pupulates token data to model.TokenInfo map
// split the file addresses.jsonl into chunks based on the logical cores of the machine
// and each chunk will be processed by a separate goroutine
func populateTokenData() {

	// Reset the map
	model.TokenInfo = sync.Map{}

	for i := 0; i < totalChunks; i++ {
		go func(i int) {
			for _, address := range addresses[i] {
				url := fullUrl + address
				resp, err := http.Get(url)
				if err != nil {
					zlog.Error("Error getting token data", zap.String("url", url), zap.Error(err))
					continue
				}
				defer resp.Body.Close()

				var c contract
				err = json.NewDecoder(resp.Body).Decode(&c)
				if err != nil {
					zlog.Error("Error decoding token data", zap.String("url", url), zap.Error(err))
					continue
				}
				supply, ok := new(big.Int).SetString(c.ContractMetadata.TotalSupply, 0)
				if !ok {
					zlog.Error("Error parsing token data", zap.String("total supply", c.ContractMetadata.TotalSupply))
					continue
				}
				model.TokenInfo.Store(address, &model.Token{
					Name:        c.ContractMetadata.Name,
					Symbol:      c.ContractMetadata.Symbol,
					Address:     c.Address,
					Decimals:    0, //NFT contract, decimal places are 0
					TotalSupply: supply,
				})
			}
		}(i)
	}
}

func populateAddressData() {

	addresses = make([][]string, 0)
	for i := 0; i < totalChunks; i++ {
		addresses = append(addresses, make([]string, 0))
	}

	//If environment variable "ADDRESS_FILE" is not set use default address file: ./data/addresses.jsonl
	tmpDir := os.Getenv("ADDRESS_FILE_DIR")
	if tmpDir != "" {
		addressFile = path.Join(tmpDir, "/addresses.jsonl")
	}

	f, err := os.Open(addressFile)
	if err != nil {
		zlog.Fatal("fatal error while trying to open addresses.toml", zap.Error(err))
	}
	defer f.Close()
	idx := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// Unmarshal to json array is still another option
		//remove the "address" word
		alltext := strings.Split(scanner.Text(), ":")

		//remove double quotes, and closing curly bracket
		a := strings.Replace(alltext[1], "\"", "", -1)
		a = strings.Trim(a, "}")
		addresses[idx] = append(addresses[idx], a)
		idx++
		if idx == totalChunks {
			idx = 0
		}
	}

	if err := scanner.Err(); err != nil {
		zlog.Fatal("fatal error while scan the addresses.toml", zap.Error(err))
	}
}
