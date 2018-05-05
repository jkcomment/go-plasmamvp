package main

import (
	"fmt"
	"os"

	contract "github.com/yuzushioh/go-plasmamvp/root_chain/contracts"
)

func main() {
	deployer := NewDeployer()

	if err := deployer.Dial("http://localhost:8545"); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	prvkey := os.Args[1]
	if len(prvkey) > 1 {
		fmt.Fprintf(os.Stdout, "private key is necessary for the deploy")
		return
	}

	deployer.SetPrvKey(prvkey)

	addr, err := deployer.Deploy(contract.RootChainABI, contract.RootChainBin)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	fmt.Println(addr.String())
}
