package main

import (
	"context"
	"flag"
	"math/big"

	"depinrc-20/abi/powerc20factory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/sirupsen/logrus"
)

var (
	name                    string
	symbol                  string
	totalSupply             uint64
	decimals                uint64
	difficulty              uint64
	mintLimitPerAddress     uint64
	limitPerMint            uint64
	privateKey              string
	factoryContractAddress  string
	verifierContractAddress string
	chainEndpoint           = "https://babel-api.testnet.iotex.io"
	logger                  = logrus.New()
)

func init() {
	flag.StringVar(&name, "name", "DePinRC20-test", "token name")
	flag.StringVar(&symbol, "symbol", "DePinRC20-test", "abbreviation of the token")
	flag.Uint64Var(&totalSupply, "totalSupply", 1000000000, "total supply of the token")
	flag.Uint64Var(&decimals, "decimals", 18, "the number of decimals for the token")
	flag.Uint64Var(&difficulty, "difficulty", 50, "mining difficulty")
	flag.Uint64Var(&mintLimitPerAddress, "mintLimitPerAddress", 200000000, "mint limit per address")
	flag.Uint64Var(&limitPerMint, "limitPerMint", 1000, "limit per mint")
	flag.StringVar(&privateKey, "privateKey", "", "Private key for the IoTeX account")
	flag.StringVar(&factoryContractAddress, "factoryContractAddress", "0x2a96010335A9A35ca29CA47D35ed8730cC839b39", "Address of the factory contract")
	flag.StringVar(&verifierContractAddress, "verifierContractAddress", "0xb9c809BA641EFdAcd6ca34515F87affd1d8Cc1a5", "Address of the verifier contract")

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	flag.Parse()
	writer := uilive.New()

	writer.Start()
	defer writer.Stop()

	logger.Info(color.GreenString("Establishing connection with Ethereum client..."))
	client, err := ethclient.Dial(chainEndpoint)
	if err != nil {
		logger.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	logger.Info(color.GreenString("Successfully connected to Ethereum client."))
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		logger.Fatalf("Error in parsing private key: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logger.Fatalf("Failed to get chainID: %v", err)
	}
	logger.Infof(color.GreenString("Successfully connected to Ethereum network with Chain ID: %v"), chainID)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		logger.Fatalf("Failed to create transactor: %v", err)
	}

	contractAddr := common.HexToAddress(factoryContractAddress)
	contract, err := powerc20factory.NewPowerc20factory(contractAddr, client)
	if err != nil {
		logger.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	logger.Info(color.GreenString("DePinRC20Factory token contract successfully instantiated."))

	verifyContractAddr := common.HexToAddress("verifierContractAddress")

	tx, err := contract.CreatePoWERC20(auth, name, symbol, new(big.Int).SetUint64(totalSupply), uint8(decimals),
		new(big.Int).SetUint64(difficulty), new(big.Int).SetUint64(mintLimitPerAddress), new(big.Int).SetUint64(limitPerMint), verifyContractAddr)
	if err != nil {
		logger.Fatalf("Failed to submit factory transaction: %v", err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatalf("Failed to mine the transaction: %v", err)
	}
	logger.Infof(color.GreenString("Factory transaction successfully confirmed, Transaction Hash: %s"), color.CyanString(receipt.TxHash.Hex()))

	all, err := contract.GetAllContracts(nil)
	if err != nil {
		logger.Fatalf("Failed to get all contracts: %v", err)
	}
	logger.Infof(color.GreenString("New DePinRC20 contract: %v"), color.CyanString("%v", all[len(all)-1]))
}
