package main

import (
	"fmt"
	"os"

	"github.com/yuzushioh/go-plasmamvp/rootchain"
	contract "github.com/yuzushioh/go-plasmamvp/rootchain/contracts"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "specify private key")
		os.Exit(1)
	}

	prvkey := os.Args[1]
	deploy := rootchain.NewDeploy(prvkey)

	if err := deploy.Dial("http://localhost:8545"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	addr, err := deploy.Deploy(contract.RootChainABI, contract.RootChainBin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(addr.String())
}
