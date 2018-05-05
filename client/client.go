package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "github.com/yuzushioh/go-plasmamvp/rootchain/contracts"
)

type Client struct {
	RootChain *contract.RootChain
}

func New(url, addr string) (*Client, error) {
	conn, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	rc, err := contract.NewRootChain(common.HexToAddress(addr), conn)
	if err != nil {
		return nil, err
	}

	return &Client{RootChain: rc}, nil
}

func (c *Client) Deposit(ctx context.Context, from common.Address, value *big.Int) (*types.Transaction, error) {
	opts := &bind.TransactOpts{
		From:    from,
		Value:   value,
		Context: ctx,
	}

	return c.RootChain.Deposit(opts)
}
