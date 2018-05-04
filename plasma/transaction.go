package plasma

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction TODO: add description
type Transaction struct {
	// Input 1
	BlockNum1    *big.Int
	TxIndex1     *big.Int
	OutputIndex1 *big.Int
	Sig1         []byte

	// Input 2
	BlockNum2    *big.Int
	TxIndex2     *big.Int
	OutputIndex2 *big.Int
	Sig2         []byte

	// Output 1
	NewOwner1 *common.Address
	Amount1   *big.Int

	// Output 2
	NewOwner2 *common.Address
	Amount2   *big.Int

	// Fee
	Fee *big.Int

	Confirmation1 bool
	Confirmation2 bool

	Spent1 bool
	Spent2 bool
}

func NewTransaction(blockNum1, txIndex1, outputIndex1, blockNum2, txIndex2, outputIndex2, amount1, amount2, fee *big.Int, newOwner1, newOwner2 *common.Address) *Transaction {
	return &Transaction{
		BlockNum1: blockNum1,
		TxIndex1: txIndex1,
		OutputIndex1: outputIndex1,
		Sig1: nil,
		BlockNum2: blockNum2,
		TxIndex2: txIndex2,
		OutputIndex2: outputIndex2,
		Sig2: nil,
		NewOwner1: newOwner1,
		Amount1: amount1,
		NewOwner2: newOwner2,
		Amount2: amount2,
		Fee: fee,
		Confirmation1: false,
		Confirmation2: false,
		Spent1: false,
		Spent2: false
	}
}

// Hash hashes transaction in sha3(rlp.encode(tx)) format.
func (tx *Transaction) Hash() (h common.Hash) {
	return rlpHash(tx)
}

// GetSender1 returns the address of sender 1 in tx
func (tx *Transaction) GetSender1() (common.Address, error) {
	return getSender(tx.Hash().Bytes(), tx.Sig1)
}

// GetSender2 returns the address of sender 2 in tx
func (tx *Transaction) GetSender2() (common.Address, error) {
	return getSender(tx.Hash.Bytes(), tx.Sig2)
}

// Sign1 signs tx and pass it to sig1 in tx
func (tx *Transaction) Sign1(prvKey ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash().Bytes(), prvKey)

	if err != nil {
		return err
	}

	tx.Sig1 = sig
	return nil
}

// Sign2 signs tx and pass it to sig2 in tx
func (tx *Transaction) Sign2(prv ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash().Bytes(), prv)

	if err != nil {
		return err
	}

	tx.Sig2 = sig
	return nil
}

// IsSingleUTXO returns if the tx conrains single UTXO.
func (tx *Transaction) IsSingleUTXO() bool {
	if tx.BlockNum2.Comp(big.NewInt(0)) == 0 {
		return true
	} else {
		return false
	}
}

// getSender returns the address of sender, in this case, address of public key
// of private key used to sign sig.
func getSender(hash []byte, sig []byte) (common.Address, error) {
	pubKeyBytes, err := crypto.Ecrecover(hash, sig)

	if err != nil {
		return common.Address{}, err
	}

	pubKey := crypto.ToECDSAPub(pubKeyBytes)
	return crypto.PubkeyToAddress(*pubKey), nil
}

// rlpHash returns the hash of rlp-encoded value of x.
func rlpHash(x interface{}) (h common.Hash) {
	w := sha3.NewKeccak256()
	rlp.Encode(w, x)
	w.Sum(h[:0])
	return h
}
