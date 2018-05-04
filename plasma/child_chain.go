package plasma

import contract "github.com/yuzushioh/go-plasmamvp/contracts"

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
