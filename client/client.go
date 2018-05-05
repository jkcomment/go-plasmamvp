package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yuzushioh/go-ethereum/crypto"
	contract "github.com/yuzushioh/go-plasmamvp/rootchain/contracts"
)

// Client todo
type Client struct {
	RootChain *contract.RootChain
	Auth      *bind.TransactOpts
}

// New TODO
func New(url, contAddr, prvKey string) (*Client, error) {
	conn, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	rc, err := contract.NewRootChain(common.HexToAddress(contAddr), conn)
	if err != nil {
		return nil, err
	}

	return &Client{
			RootChain: rc,
			Auth:      setPrvKey(prvKey),
		},
		nil
}

// Deposit TODO
func (c *Client) Deposit(ctx context.Context, value *big.Int) (*types.Transaction, error) {
	c.Auth.Context = ctx
	c.Auth.Value = value
	return c.RootChain.Deposit(c.Auth)
}

// SetPrvKey creates keyed-transactor with specified private key.
func setPrvKey(prvkeyHex string) *bind.TransactOpts {
	keyBytes := common.FromHex(prvkeyHex)
	key := crypto.ToECDSAUnsafe(keyBytes)
	return bind.NewKeyedTransactor(key)
}
