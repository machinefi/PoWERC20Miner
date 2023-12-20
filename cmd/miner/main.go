package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"Powerc20Worker/abi/powerc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/gosuri/uilive"
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
	flag.StringVar(&privateKey, "privateKey", "", "Private key for the Ethereum account")
	flag.StringVar(&contractAddress, "contractAddress", "0x53870E89D4E93442ed7FACD68f7C0221B27e5958", "Address of the Ethereum contract")
	flag.IntVar(&workerCount, "workerCount", 10, "Number of concurrent mining workers")

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

type result struct {
	challenge  string
	address    string
	nonce      string
	difficulty string
	diff       string
}

func mineWorker(ctx context.Context, wg *sync.WaitGroup, diffStr, challengeStr, address string, fromAddress common.Address, client *ethclient.Client, resultChan chan<- *result, errorChan chan<- error, challenge *big.Int, target *big.Int, hashCountChan chan<- int) {
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
			diff, difficulty, err := getDifficultyAndDiff(diffStr, challengeStr, nonceStr, address)
			// diff := ""
			// difficulty := ""
			// err = errors.New("")
			// challengePadded := common.LeftPadBytes(challenge.Bytes(), 32)
			// addressBytes := fromAddress.Bytes()
			// data := append(challengePadded, append(addressBytes, noncePadded...)...)
			// hash := crypto.Keccak256Hash(data)
			// if hash.Big().Cmp(target) == -1 {
			// 	resultChan <- nonce
			// 	return
			// }
			if err == nil {
				resultChan <- &result{
					challenge:  challengeStr,
					address:    address,
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

func getDifficultyAndDiff(sourceDifficulty, challenge, nonce, address string) (diff, difficulty string, err error) {
	bin := "./gen-poseidon-pse.darwin."
	bin += runtime.GOARCH
	cmd := exec.Command(bin, sourceDifficulty, challenge, address, nonce)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", err
	}
	outputs := strings.Split(string(output), ",")
	if len(outputs) != 2 {
		logger.Fatalf("Unexpected poseidon result")
	}
	diff = strings.TrimPrefix(strings.TrimSpace(outputs[0]), "diff:")
	difficulty = strings.TrimPrefix(strings.TrimSpace(outputs[1]), "difficulty:")
	return
}

func main() {
	banner := `
//  ____    __        _______ ____   ____ ____   ___    __  __ _                 
// |  _ \ __\ \      / / ____|  _ \ / ___|___ \ / _ \  |  \/  (_)_ __   ___ _ __ 
// | |_) / _ \ \ /\ / /|  _| | |_) | |     __) | | | | | |\/| | | '_ \ / _ \ '__|
// |  __/ (_) \ V  V / | |___|  _ <| |___ / __/| |_| | | |  | | | | | |  __/ |   
// |_|   \___/ \_/\_/  |_____|_| \_\\____|_____|\___/  |_|  |_|_|_| |_|\___|_|   
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainID)
	if err != nil {
		logger.Fatalf("Failed to create transactor: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	contract, err := powerc20.NewPowerc20(contractAddr, client)
	if err != nil {
		logger.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	logger.Info(color.GreenString("PoWERC20 token contract successfully instantiated."))

	contractName, err := contract.Name(nil)
	if err != nil {
		logger.Fatalf("Failed to get contract name: %v", err)
	}
	logger.Infof(color.GreenString("Contract Name: %s"), color.RedString(contractName))

	challenge, err := contract.Challenge(nil)
	if err != nil {
		logger.Fatalf("Failed to get challenge: %v", err)
	}
	logger.Infof(color.GreenString("Current mining challenge number: %d"), challenge)

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

	diffStr := strconv.Itoa(int(difficulty.Uint64()))
	challengeStr := "0x" + fmt.Sprintf("%x", challenge)
	address := crypto.PubkeyToAddress(privateKeyECDSA.PublicKey).String()

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
		go mineWorker(ctx, &wg, diffStr, challengeStr, address, auth.From, client, resultChan, errorChan, challenge, target, hashCountChan)
	}

	select {
	case nonce := <-resultChan:
		ticker.Stop()
		cancel()
		wg.Wait()
		logger.Infof(color.GreenString("Successfully discovered a valid nonce: %s"), nonce.nonce)

		cmd := fmt.Sprintf(`ioctl ws message send --project-id 20000 --project-version "0.1" --data "{\"challenge\": \"%s\",\"address\": \"%s\",\"nonce\": \"%s\",\"difficulty\": \"%s\",\"diff\": \"%s\"}"`, nonce.challenge, nonce.address, nonce.nonce, nonce.difficulty, nonce.diff)

		logger.Infof(color.GreenString("Use this cmd to submit nonce: %s"), color.CyanString(cmd))

	case err := <-errorChan:
		cancel()
		wg.Wait()
		logger.Fatalf("Mining operation failed due to an error: %v", err)
	}
	logger.Info(color.GreenString("Mining process successfully completed"))
}
