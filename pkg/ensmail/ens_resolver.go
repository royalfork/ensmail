package ensmail

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/royalfork/ensmail/pkg/ens"
)

var (
	ErrNoResolver = errors.New("no resolver set")
	ErrNoEmail    = errors.New("no email set")
)

type ENSResolver struct {
	caller   bind.ContractCaller
	registry *ens.ENSCaller
}

func NewENSResolver(registryAddr common.Address, caller bind.ContractCaller) (*ENSResolver, error) {
	registry, err := ens.NewENSCaller(registryAddr, caller)
	if err != nil {
		return nil, err
	}

	return &ENSResolver{
		caller:   caller,
		registry: registry,
	}, nil
}

// Email returns the email text record for the given name.  Before
// querying the ENS registry, the ".eth" suffix is added to name.
func (r *ENSResolver) Email(ctx context.Context, name string) (string, error) {
	const (
		TLDSuffix = ".eth"
		// Defined by https://docs.ens.domains/ens-improvement-proposals/ensip-5-text-records
		TextEmailKey = "email"
	)

	node, err := ens.NameHash(name + TLDSuffix)
	if err != nil {
		return "", err
	}

	callOpts := &bind.CallOpts{Context: ctx}

	resolverAddr, err := r.registry.Resolver(callOpts, node)
	if err != nil {
		return "", err
	} else if resolverAddr == (common.Address{}) {
		return "", ErrNoResolver
	}

	resolver, err := ens.NewTextResolverCaller(resolverAddr, r.caller)
	if err != nil {
		return "", err
	}

	email, err := resolver.Text(callOpts, node, TextEmailKey)
	if err != nil {
		return "", err
	} else if email == "" {
		return "", ErrNoEmail
	}

	return email, nil
}
