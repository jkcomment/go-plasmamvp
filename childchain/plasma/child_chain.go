package plasma

import "github.com/yuzushioh/go-plasmamvp/rootchain/contracts"

// ChildChain TODO
type ChildChain struct {
}

// NewChildChain TODO
func NewChildChain() *ChildChain {
	return &ChildChain{}
}

// ApplyDeposit TODO
func (cc *ChildChain) ApplyDeposit(deposit *contract.RootChainDeposit) {

}

func (cc *ChildChain) ApplyTransaction(tx *Transaction) error{

}
