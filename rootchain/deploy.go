package rootchain

import (
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Deploy deploys smart contract with abi and bin
type Deploy struct {
	con  *ethclient.Client
	auth *bind.TransactOpts
}

// NewDeploy creates Deployer
func NewDeploy(prvkeyHex string) *Deploy {
	return &Deploy{
		auth: setPrvKey(prvkeyHex),
	}
}

// Dial creates ethclient with specified URL
func (d *Deploy) Dial(rawURL string) error {
	con, err := ethclient.Dial(rawURL)
	if err != nil {
		return err
	}
	d.con = con
	return nil
}

// Deploy deploys smart contract with given abi and bin
func (d *Deploy) Deploy(contractAbi, contractBin string) (*common.Address, error) {
	parsed, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	dataInputs := make([]interface{}, 0, len(parsed.Constructor.Inputs))

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

// SetPrvKey creates keyed-transactor with specified private key.
func setPrvKey(prvkeyHex string) *bind.TransactOpts {
	keyBytes := common.FromHex(prvkeyHex)
	key := crypto.ToECDSAUnsafe(keyBytes)
	return bind.NewKeyedTransactor(key)
}
