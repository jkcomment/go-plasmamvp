package main

import (
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Deployer deploys smart contract with abi and bin
type Deployer struct {
	con  *ethclient.Client
	auth *bind.TransactOpts
}

// NewDeployer creates Deployer
func NewDeployer() *Deployer {
	return &Deployer{}
}

// Dial creates ethclient with specified URL
func (d *Deployer) Dial(rawURL string) error {
	con, err := ethclient.Dial(rawURL)
	if err != nil {
		return err
	}
	d.con = con
	return nil
}

// SetPrvKey creates keyed-transactor with specified private key.
func (d *Deployer) SetPrvKey(keyHex string) {
	keyBytes := common.FromHex(keyHex)
	key := crypto.ToECDSAUnsafe(keyBytes)
	d.auth = bind.NewKeyedTransactor(key)
}

// Deploy deploys smart contract with given abi and bin
func (d *Deployer) Deploy(contractAbi, contractBin string) (*common.Address, error) {
	parsed, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	var dataInputs = make([]interface{}, 0)

	for _, i := range parsed.Constructor.Inputs {
		t := reflect.Zero(i.Type.Type)
		v := t.Interface()
		if t.Kind() == reflect.Ptr && t.IsNil() {
			elem := reflect.TypeOf(v).Elem()
			v2 := reflect.New(elem)
			v = v2.Interface()
		}
		dataInputs = append(dataInputs, v)
	}

	address, _, _, err := bind.DeployContract(d.auth, parsed, common.FromHex(contractBin), d.con, dataInputs...)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
