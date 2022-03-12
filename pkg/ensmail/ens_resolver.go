package ensmail

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/idna"
)

type ENSResolver struct {
}

func NewENSResolver() (ENSResolver, error) {
	return ENSResolver{}, nil
}

// Resolve receives a 2nd level .eth label, and returns the email address associated
func (r ENSResolver) Email(name string) (string, error) {
	// normalize name
	// Namehash name
	// get resolver from registry
	// ask resolver for text/email
	// ensure email is valid
	return "", nil
}

// Implementation of
// https://docs.ens.domains/ens-improvement-proposals/ensip-1-ens#namehash-algorithm
func nameHash(name string) ([32]byte, error) {
	var node common.Hash

	// Because strings.Split("", ".") returns slice of len 1, special
	// case is needed to return before any hashing occurs.
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
