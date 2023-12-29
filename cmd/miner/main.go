package main

import (
	"context"
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"sync"
	"time"

	"depinrc-20/abi/powerc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/gosuri/uilive"
	"github.com/machinefi/go-iden3-crypto/poseidon"
	"github.com/sirupsen/logrus"
)

var (
	infuraURL       = "https://babel-api.testnet.iotex.io"
	privateKey      string
	contractAddress string
	workerCount     int
	logger          = logrus.New()
)

func init() {
	flag.StringVar(&privateKey, "privateKey", "", "Private key for the IoTex account")
	flag.StringVar(&contractAddress, "contractAddress", "", "Address of the IoTex DePinRC20 contract")
	flag.IntVar(&workerCount, "workerCount", 10, "Number of concurrent mining workers")

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableHTMLEscape: true,
	})
}

type result struct {
	contract   string
	sender     string
	nonce      string
	difficulty string
	diff       string
}

func mineWorker(ctx context.Context, wg *sync.WaitGroup, difficulty *big.Int, contractAddress, sender string,
	resultChan chan<- *result, errorChan chan<- error, hashCountChan chan<- int) {
	defer wg.Done()

	var nonce *big.Int
	var err error

	for {
		select {
		case <-ctx.Done():
			return
		default:
			nonce, err = rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 256))
			if err != nil {
				errorChan <- fmt.Errorf("failed to generate random nonce: %v", err)
				return
			}

			noncePadded := common.LeftPadBytes(nonce.Bytes(), 32)
			nonceStr := "0x" + fmt.Sprintf("%x", noncePadded)
			diff, difficulty, err := getDifficultyAndDiff(difficulty, contractAddress, nonceStr, sender)

			if err == nil {
				resultChan <- &result{
					contract:   contractAddress,
					sender:     sender,
					nonce:      nonceStr,
					difficulty: difficulty,
					diff:       diff,
				}
				return
			}
			hashCountChan <- 1
		}
	}
}

func getDifficultyAndDiff(difficulty *big.Int, contract, nonce, address string) (string, string, error) {
	contractBI := new(big.Int)
	contractBI.SetString(contract[2:], 16)
	nonceBI := new(big.Int)
	nonceBI.SetString(nonce[2:], 16)
	addressBI := new(big.Int)
	addressBI.SetString(address[2:], 16)

	poseidonHash, err := poseidon.HashWithWidth([]*big.Int{contractBI, addressBI, nonceBI}, 5)
	if err != nil {
		return "", "", err
	}

	bigMax := new(big.Int).Lsh(big.NewInt(1), 256)
	bigMax = bigMax.Sub(bigMax, big.NewInt(1))
	difficultyBI := bigMax.Rsh(bigMax, uint(difficulty.Uint64()))

	if poseidonHash.Cmp(difficultyBI) > 0 {
		return "", "", errors.New(
			fmt.Sprintf("hash %s is greater than difficulty %s", poseidonHash.String(), difficultyBI.String()))
	}

	diffBI := new(big.Int)
	diffBI = diffBI.Sub(difficultyBI, poseidonHash)
	return diffBI.String(), difficultyBI.String(), nil
}

func main() {
	banner := `
// _______   _______ .______    __  .__   __. .______        ______       ___     ___      .___  ___.  __  .__   __.  _______ .______      
// |       \ |   ____||   _  \  |  | |  \ |  | |   _  \      /      |     |__ \   / _ \     |   \/   | |  | |  \ |  | |   ____||   _  \     
// |  .--.  ||  |__   |  |_)  | |  | |   \|  | |  |_)  |    |  ,----' ______ ) | | | | |    |  \  /  | |  | |   \|  | |  |__   |  |_)  |    
// |  |  |  ||   __|  |   ___/  |  | |  .    | |      /     |  |     |______/ /  | | | |    |  |\/|  | |  | |  .    | |   __|  |      /
// |  '--'  ||  |____ |  |      |  | |  |\   | |  |\  \----.|   ----.      / /_  | |_| |    |  |  |  | |  | |  |\   | |  |____ |  |\  \----.
// |_______/ |_______|| _|      |__| |__| \__| | _|  ._____| \______|     |____|  \___/     |__|  |__| |__| |__| \__| |_______|| _|  ._____|
	`
	fmt.Println(banner)
	flag.Parse()
	writer := uilive.New()

	writer.Start()
	defer writer.Stop()

	logger.Info(color.GreenString("Establishing connection with Ethereum client..."))
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		logger.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	logger.Info(color.GreenString("Successfully connected to Ethereum client."))
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		logger.Fatalf("Error in parsing private key: %v", err)
	}
	//fmt.Println(crypto.PubkeyToAddress(privateKeyECDSA.PublicKey))

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logger.Fatalf("Failed to get chainID: %v", err)
	}
	logger.Infof(color.GreenString("Successfully connected to Ethereum network with Chain ID: %v"), chainID)

	//auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	//if err != nil {
	//	logger.Fatalf("Failed to create transactor: %v", err)
	//}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := powerc20.NewPowerc20(contractAddr, client)
	if err != nil {
		logger.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	logger.Info(color.GreenString("DePinRC20 token contract successfully instantiated."))

	contractName, err := contract.Name(nil)
	if err != nil {
		logger.Fatalf("Failed to get contract name: %v", err)
	}
	logger.Infof(color.GreenString("Contract Name: %s"), color.RedString(contractName))

	difficulty, err := contract.Difficulty(nil)
	if err != nil {
		logger.Fatalf("Failed to get difficulty: %v", err)
	}
	logger.Infof(color.GreenString("Current mining difficulty level: %d"), difficulty)

	difficultyUint := uint(difficulty.Uint64())
	target := new(big.Int).Lsh(big.NewInt(1), 256-difficultyUint)
	logger.Infof(color.GreenString("Target number is: %d"), target)

	resultChan := make(chan *result)
	errorChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sender := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey).String()

	logger.Info(color.YellowString("Mining workers started..."))

	hashCountChan := make(chan int)
	totalHashCount := 0
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				timestamp := time.Now().Format("2006-01-02 15:04:05")
				hashesPerSecond := float64(totalHashCount) / 1000.0
				fmt.Fprintf(writer, "%s[%s] %s\n", color.BlueString("Mining"), timestamp, color.GreenString("Total hashes per second: %8.2f K/s", hashesPerSecond))
				totalHashCount = 0
			case count := <-hashCountChan:
				totalHashCount += count
			}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 1; /*workerCount*/ i++ {
		wg.Add(1)
		go mineWorker(ctx, &wg, difficulty, contractAddr.String(), sender, resultChan, errorChan, hashCountChan)
	}

	select {
	case nonce := <-resultChan:
		ticker.Stop()
		cancel()
		wg.Wait()
		logger.Infof(color.GreenString("Successfully discovered a valid nonce: %s"), nonce.nonce)

		cmd := fmt.Sprintf(`ioctl ws message send --project-id 20000 --project-version "0.1" --data "{\"depinRC20Address\": \"%s\",\"sender\": \"%s\",\"nonce\": \"%s\",\"difficulty\": \"%s\",\"diff\": \"%s\"}"`,
			nonce.contract, nonce.sender, nonce.nonce, nonce.difficulty, nonce.diff)

		logger.Infof(color.GreenString("Use this cmd to submit nonce: %s"), color.CyanString(cmd))

	case err := <-errorChan:
		cancel()
		wg.Wait()
		logger.Fatalf("Mining operation failed due to an error: %v", err)
	}
	logger.Info(color.GreenString("Mining process successfully completed"))
}
