package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/postables/Postables-Payment-Channel/src/airdrop/bindings"
)

func main() {
	if len(os.Args) > 2 || len(os.Args) < 2 {
		fmt.Println("improper invocation")
		fmt.Println("./airdrop <mode>")
		fmt.Println("mode: .....")
		os.Exit(1)
	}
	keyPath := os.Getenv("KEY_PATH")
	keyPass := os.Getenv("KEY_PASS")
	ipcPath := os.Getenv("IPC_PATH")
	if keyPath == "" {
		log.Fatal("KEY_PATH env var not set")
	}
	if keyPass == "" {
		log.Fatal("KEY_PASS env var not set")
	}
	if ipcPath == "" {
		log.Fatal("IPC_PATH env var not set")
	}
	keyBytes, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(ipcPath)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(string(keyBytes)), keyPass)
	if err != nil {
		log.Fatal(err)
	}

	runMode := os.Args[1]
	switch runMode {
	case "deploy":
		_, tx, _, err := bindings.DeployAirdrop(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		_, err = bind.WaitDeployed(context.Background(), client, tx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("airdrop contract deployment succeeded")
	default:
		log.Fatal("invalid run mode")
	}

}
