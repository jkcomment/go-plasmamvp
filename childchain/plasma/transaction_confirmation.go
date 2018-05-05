package plasma

import "math/big"

// TransactionConfirmation TODO
type TransactionConfirmation struct {
	BlockNum big.Int
	TxIndex  big.Int
	TxHash   big.Int
	Sig1     []byte
	Sig2     []byte
}
