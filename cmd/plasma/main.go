package main

import (
	"context"
	"fmt"
	"os"
	"github.com/pkg/errors"
	"github.com/yuzushioh/go-plasmamvp/client"
	"github.com/yuzushioh/go-plasmamvp/rootchain"
	"github.com/yuzushioh/go-plasmamvp/rootchain/contracts"
	"github.com/yuzushioh/go-plasmamvp/childchain/plasma"
	"github.com/ethereum/go-ethereum/common"
	"github.com/yuzushioh/go-plasmamvp/lib"
	"crypto/x509"
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


		amt := os.Args[2]
		value, err := lib.StrToBigInt(amt)
		if err != nil {
			return err
		}

		tx, err := client.Deposit(context.Background(), value)
		if err != nil {
			return err
		}

		fmt.Println(tx)
	case "sendtx":
		if len(os.Args) != 3 {
			return errors.New("usage: plasma sendtx <blknum1> <txindex1> <oindex1> <blknum2> <txindex2> <oindex2> <newowner1> <amount1> <newowner2> <amount2> <fee> <key1> [<key2>]")
		}

		bn1:= os.Args[2]
		blockNum1, err := lib.StrToBigInt(bn1)
		if err != nil {
			return err
		}

		ti1:= os.Args[3]
		txIndex1, err := lib.StrToBigInt(ti1)
		if err != nil {
			return err
		}

		oi1:= os.Args[4]
		outputIndex1, err := lib.StrToBigInt(oi1)
		if err != nil {
			return err
		}

		bn2:= os.Args[5]
		blockNum2, err := lib.StrToBigInt(bn2)
		if err != nil {
			return err
		}

		ti2:= os.Args[6]
		txIndex2, err := lib.StrToBigInt(ti2)
		if err != nil {
			return err
		}

		oi2:= os.Args[7]
		outputIndex2, err := lib.StrToBigInt(oi2)
		if err != nil {
			return err
		}

		no1 := os.Args[8]
		newOwner1 := common.HexToAddress(no1)

		am1:= os.Args[9]
		amount1, err := lib.StrToBigInt(am1)
		if err != nil {
			return err
		}

		no2 := os.Args[10]
		newOwner2 := common.HexToAddress(no2)

		am2:= os.Args[11]
		amount2, err := lib.StrToBigInt(am2)
		if err != nil {
			return err
		}

		f := os.Args[12]
		fee, err := lib.StrToBigInt(f)
		if err != nil {
			return err
		}

		tx := plasma.NewTransaction(
			blockNum1, txIndex1, outputIndex1,
			blockNum2, txIndex2, outputIndex2,
			amount1, amount2,
			fee,
			&newOwner1, &newOwner2,
		)

		k1 := os.Args[13]
		if k1 != "" {
			key1 ,err := x509.ParseECPrivateKey([]byte(k1))
			if err != nil {
				return err
			}

			tx.Sign1(key1)
		}

		k2 := os.Args[14]
		if k2 != "" {
			key2 ,err := x509.ParseECPrivateKey([]byte(k2))
			if err != nil {
				return err
			}

			tx.Sign2(key2)
		}

		fmt.Println(tx)

		childChain := plasma.NewChildChain()
		err = childChain.ApplyTransaction(tx)
		if err != nil {
			return err
		}

	default:
		return errors.Errorf("does not support %q", os.Args[1])
	}

	return nil
}
