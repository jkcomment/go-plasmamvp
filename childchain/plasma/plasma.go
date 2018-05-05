package plasma

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	contract "github.com/yuzushioh/go-plasmamvp/contracts"
)

// Plasma defines fields necessary for plasma.
type Plasma struct {
	con          *ethclient.Client
	contractAddr common.Address
	childChain   *ChildChain
}

// NewPlasma creates a new instance of Plasma.
func NewPlasma(contractAddr common.Address) *Plasma {
	return &Plasma{
		contractAddr: contractAddr,
		childChain:   NewChildChain(),
	}
}

// Dial connects a client to the given URL and
// pass the connection into Plasma
func (p *Plasma) Dial(rawURL string) error {
	con, err := ethclient.Dial(rawURL)
	if err != nil {
		return err
	}

	p.con = con
	return nil
}

// WatchDeposit watches root chain for the deposit event.
func (p *Plasma) WatchDeposit(ctx context.Context) error {
	filterer, err := contract.NewRootChainFilterer(p.contractAddr, p.con)
	if err != nil {
		return err
	}

	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}

	sink := make(chan *contract.RootChainDeposit)
	sub, err := filterer.WatchDeposit(opts, sink)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case deposit := <-sink:
				p.childChain.ApplyDeposit(deposit)

			case err := <-sub.Err():
				log.Warn("Plasma: Deposit subscription error ", err)
				sub.Unsubscribe()
			}
		}
	}()

	return nil
}