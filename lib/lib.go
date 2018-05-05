package lib

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strconv"
)

// SetPrvKey creates keyed-transactor with specified private key.
func SetPrvKey(prvkeyHex string) *bind.TransactOpts {
	keyBytes := common.FromHex(prvkeyHex)
	key := crypto.ToECDSAUnsafe(keyBytes)
	return bind.NewKeyedTransactor(key)
}

func StrToBigInt(num string) (*big.Int,error) {
	value, err := strconv.Atoi(num)
	if err != nil {
		return nil,err
	}

	return big.NewInt(int64(value)),nil
}
