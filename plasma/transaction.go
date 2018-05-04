package plasma

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Transaction TODO: add description
type Transaction struct {
	BlockNum1    big.Int
	TxIndex1     big.Int
	OutputIndex1 big.Int
	BlockNum2    big.Int
	TxIndex2     big.Int
	OutputIndex2 big.Int
	NewOwner1    common.Address
	Amount1      big.Int
	NewOwner2    common.Address
	Amount2      big.Int
	Sign1        []byte
	Sign2        []byte
}

func (tx *Transaction) isSingleUTXO() bool {
	if tx.BlockNum2.Comp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}
}
