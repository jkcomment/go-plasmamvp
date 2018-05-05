package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/yuzushioh/go-plasmamvp/rootchain"
	contract "github.com/yuzushioh/go-plasmamvp/rootchain/contracts"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("Specify a subcommand")
	}

	subCmd := os.Args[1]
	switch subCmd {
	case "deploy":
		if len(os.Args) != 3 {
			return errors.New("Specify private key")
		}

		prvkey := os.Args[2]
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
		if  contAddr == "" {
			return errors.New("exec following command before plasma deposit: export CONTRACT_ADDRESS=$(plasma deploy [Private Key])")
		}
		fmt.Printf("CONTRACT_ADDRESS: %s\n", contAddr)

		if len(os.Args) != 4 {
			return errors.New("usage: plasma deposit [ADDRESS] [AMOUNT]")
		}

		addr:= os.Args[2]
		amt := os.Args[3]
		// TODO: use deposit
		fmt.Printf("address: %s, amount: %s\n",addr,amt)
	default:
		return errors.Errorf("does not support %q", os.Args[1])
	}

	return nil
}
