package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/yuzushioh/go-plasmamvp/childchain/plasma"
	"github.com/yuzushioh/go-plasmamvp/client"
	"github.com/yuzushioh/go-plasmamvp/lib"
	"github.com/yuzushioh/go-plasmamvp/rootchain"
	contract "github.com/yuzushioh/go-plasmamvp/rootchain/contracts"
)

func main() {
	contAddr := os.Getenv("CONTRACT_ADDRESS")
	plasma := plasma.NewPlasma(common.HexToAddress(contAddr))

	if err := plasma.Dial("ws://localhost:8545"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if err := plasma.WatchDeposit(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("Specify a subcommand")
	}

	prvkey := os.Getenv("PRIVATE_KEY")
	if prvkey == "" {
		return errors.New("exec following command before plasma deposit: export PRIVATE_KEY=[Private Key]")
	}
	fmt.Printf("PRIVATE_KEY: %s\n", prvkey)

	subCmd := os.Args[1]

	switch subCmd {
	case "deploy":
		if len(os.Args) != 2 {
			return errors.New("usage: plasma deploy")
		}

		deploy := rootchain.NewDeploy(prvkey)

		if err := deploy.Dial("http://localhost:8545"); err != nil {
			return err
		}

		addr, err := deploy.Deploy(contract.RootChainABI, contract.RootChainBin)
		if err != nil {
			return err
		}

		fmt.Println(addr.String())

	case "deposit":
		contAddr := os.Getenv("CONTRACT_ADDRESS")

		if contAddr == "" {
			return errors.New("exec following command before plasma deposit: export CONTRACT_ADDRESS=$(plasma deploy [Private Key])")
		}
		fmt.Printf("CONTRACT_ADDRESS: %s\n", contAddr)

		if len(os.Args) != 3 {
			return errors.New("usage: plasma deposit [AMOUNT]")
		}

		client, err := client.New("http://localhost:8545", contAddr, prvkey)
		if err != nil {
			return err
		}

		value, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}

		tx, err := client.Deposit(context.Background(), big.NewInt(int64(value)))
		if err != nil {
			return err
		}

		fmt.Println(lib.SprintTransaction(tx))

	default:
		return errors.Errorf("does not support %q", os.Args[1])
	}

	return nil
}
