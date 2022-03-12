package ensmail

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/royalfork/ensmail/pkg/ens"
	"github.com/royalfork/soltest"
)

func TestEmail(t *testing.T) {
	// Before subtests run, the following is setup:
	// - Registry is deployed, with accts[0] owning root
	// - .eth TLD registered, owned by accts[0]
	// - name.eth domain registered, owned by accts[1]
	// - public resolver deployed

	chain, accts := soltest.New()

	_, _, registry, err := ens.DeployENSRegistry(accts[0].Auth, chain)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	// Create eth tld
	ethLabel, err := labelHash("eth")
	if err != nil {
		t.Fatal(err)
	}

	if !chain.Succeed(registry.SetSubnodeOwner(accts[0].Auth, [32]byte{}, ethLabel, accts[0].Addr)) {
		t.Fatal("unable to create eth tld")
	}

	ethHash, err := nameHash("eth")
	if err != nil {
		t.Fatal(err)
	}

	nameLabel, err := labelHash("name")
	if err != nil {
		t.Fatal(err)
	}

	if !chain.Succeed(registry.SetSubnodeOwner(accts[0].Auth, ethHash, nameLabel, accts[1].Addr)) {
		t.Fatal("unable to create eth tld")
	}

	// TODO deploy resolver

	t.Run("nameNoExist", func(t *testing.T) {
	})

	t.Run("noResolver", func(t *testing.T) {
	})

	t.Run("resolverNoText", func(t *testing.T) {
	})

	t.Run("resolverNoEmail", func(t *testing.T) {
	})

	t.Run("resolverInvalidEmail", func(t *testing.T) {
	})

	t.Run("email", func(t *testing.T) {
	})
}

func TestNameHash(t *testing.T) {
	for _, test := range []struct {
		input, hexOutput string
	}{
		{"", "0000000000000000000000000000000000000000000000000000000000000000"},
		{"eth", "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"},
		{"foo.eth", "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"},
		{"FoO.eTh", "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"},
	} {
		t.Run(test.input, func(t *testing.T) {
			out, err := nameHash(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if exp := common.HexToHash(test.hexOutput); bytes.Compare(out[:], exp[:]) != 0 {
				t.Errorf("namehash(%s): want: %x, got: %x", test.input, exp, out)
			}
		})
	}
}
