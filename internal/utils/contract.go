package utils

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethRPC struct {
	client *ethclient.Client
}

func (r *ethRPC) Call(contractABI, address, funcName string) ([]interface{}, error) {
	contractAbi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(address)
	data, err := contractAbi.Pack(funcName)
	if err != nil {
		return nil, err
	}
	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}
	res, err := r.client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling contract: %v", err)
	}
	callRes, err := contractAbi.Unpack(funcName, res)
	if err != nil {
		return nil, fmt.Errorf("error unpacking result: %v", err)

	}
	return callRes, nil
}

func CastNumberWithDecimals(number interface{}, decimals int32) float64 {
	num := new(big.Float).SetInt(number.(*big.Int))
	dcmls := new(big.Float).SetFloat64(math.Pow(10, float64(decimals)))
	castedNumber, _ := new(big.Float).Quo(num, dcmls).Float64()
	return castedNumber
}

func (r *ethRPC) Close() {
	r.client.Close()
}
func NewRPC() *ethRPC {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/07ba725ae9104ef2b7a06d12f3ca38fe")
	if err != nil {
		log.Fatalf("Error connecting to rpc: %v", err)
	}
	return &ethRPC{
		client: client,
	}
}
