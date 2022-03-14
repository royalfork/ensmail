package ens

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/royalfork/soltest"
)

// Test allows easy ENS tests.  Chain is a simulated geth blockchain.
// Accts is a slice of accounts with pre-allocated ethereum
// balance. RegistryAddr and Registry refer to a deployed ENS registry
// on Chain, where root is owned by Accts[0].  The ".eth" TLD subnode
// is registered and also owned by Accts[0].  ResolverAddr and
// Resolver refer to a deployed PublicResolver on Chain.
type Test struct {
	Chain        soltest.TestChain
	Accts        []soltest.TestAccount
	RegistryAddr common.Address
	Registry     *ENSRegistry
	ResolverAddr common.Address
	Resolver     *PublicResolver
}

func NewTest() (t Test, err error) {
	t.Chain, t.Accts = soltest.New()

	t.RegistryAddr, _, t.Registry, err = DeployENSRegistry(t.Accts[0].Auth, t.Chain)
	if err != nil {
		return t, err
	}
	t.Chain.Commit()

	// Create eth tld
	ethLabel, err := LabelHash("eth")
	if err != nil {
		return t, err
	}

	if !t.Chain.Succeed(t.Registry.SetSubnodeOwner(t.Accts[0].Auth, [32]byte{}, ethLabel, t.Accts[0].Addr)) {
		return t, errors.New("unable to create eth tld")
	}

	t.ResolverAddr, _, t.Resolver, err = DeployPublicResolver(t.Accts[0].Auth, t.Chain, t.RegistryAddr)
	if err != nil {
		return t, err
	}
	t.Chain.Commit()

	return t, nil
}

// Nodehash of "eth".
var ethHash = func() [32]byte {
	nh, err := NameHash("eth")
	if err != nil {
		panic(err)
	}
	return nh
}()

// Register registers label under the .eth TLD (creating "label.eth"), setting owner to owner.
func (t Test) Register(owner common.Address, label string) ([32]byte, error) {
	lh, err := LabelHash(label)
	if err != nil {
		return [32]byte{}, err
	}

	if !t.Chain.Succeed(t.Registry.SetSubnodeOwner(t.Accts[0].Auth, ethHash, lh, owner)) {
		return [32]byte{}, errors.New("unable to register label")
	}
	return NameHash(label + ".eth")
}
