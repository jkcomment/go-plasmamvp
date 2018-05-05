package main

import (
	"fmt"
	"os"

	"github.com/yuzushioh/go-plasmamvp/rootchain"
	contract "github.com/yuzushioh/go-plasmamvp/rootchain/contracts"
)

func main() {
	deployer := rootchain.NewDeploy()

	if err := deployer.Dial("http://localhost:8545"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "specify private key")
		return
	}

	prvkey := os.Args[1]

	deployer.SetPrvKey(prvkey)

	addr, err := deployer.Deploy(contract.RootChainABI, contract.RootChainBin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	fmt.Println(addr.String())
}
