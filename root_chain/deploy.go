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

	if len(os.Args) > 1 {
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
