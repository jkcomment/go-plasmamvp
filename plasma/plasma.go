package plasma

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yuzushioh/go-plasmamvp/contracts"
)

// Plasma TODO: add description
type Plasma struct {
	client *ethclient.Client

	rootchain contracts.RootChain
}
