package ensmail

import (
	"context"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/royalfork/ensmail/pkg/ens"
	"golang.org/x/net/idna"
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

	node, err := nameHash(name + TLDSuffix)
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

// Implementation of
// https://docs.ens.domains/ens-improvement-proposals/ensip-1-ens#namehash-algorithm
func nameHash(name string) ([32]byte, error) {
	var node common.Hash

	// Because strings.Split("", ".") returns slice of len 1, must
	// return 0x0 before any hashing occurs.
	if name == "" {
		return node, nil
	}

	labels := strings.Split(name, ".")
	for i := len(labels) - 1; i >= 0; i-- {
		labelHash, err := labelHash(labels[i])
		if err != nil {
			return node, err
		}

		node = crypto.Keccak256Hash(node[:], labelHash[:])
	}

	return node, nil
}

// From https://docs.ens.domains/ens-improvement-proposals/ensip-1-ens#name-syntax:
// Each label must be a valid normalised label as described in UTS46
// with the options transitional=false and useSTD3AsciiRules=true
var ensProfile = idna.New(idna.Transitional(false), idna.StrictDomainName(true), idna.MapForLookup())

func labelHash(label string) ([32]byte, error) {
	// By definition, labelhash of "" is 0x0
	if label == "" {
		return [32]byte{}, nil
	}

	if strings.Contains(label, ".") {
		return [32]byte{}, errors.New("label contains period")
	}

	normalizedLabel, err := ensProfile.ToUnicode(label)
	if err != nil {
		return [32]byte{}, err
	}

	return crypto.Keccak256Hash([]byte(normalizedLabel)), nil
}
