package plasma

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	contract "github.com/yuzushioh/go-plasmamvp/contracts"
)

var (
	nullAddress = common.HexToAddress("")
	nullByte    = common.HexToHash("").Bytes()
)

// ChildChain TODO
type ChildChain struct {
	Blocks map[*big.Int]Block
}

// NewChildChain TODO
func NewChildChain(blocks map[*big.Int]Block) *ChildChain {
	return &ChildChain{Blocks: blocks}
}

// ApplyDeposit TODO
func (cc *ChildChain) ApplyDeposit(deposit *contract.RootChainDeposit) {
	tx := Transaction{
		BlockNum1:    deposit.DepositBlock,
		TxIndex1:     common.Big0,
		OutputIndex1: common.Big0,
		BlockNum2:    common.Big0,
		TxIndex2:     common.Big0,
		OutputIndex2: common.Big0,
		Amount1:      deposit.Amount,
		Amount2:      common.Big0,
		Fee:          common.Big0,
		NewOwner1:    &deposit.Depositor,
		NewOwner2:    &nullAddress,
	}

	cc.Blocks[deposit.DepositBlock] = Block{
		Transactions: []Transaction{
			tx,
		},
		sig: nullByte,
	}
}
