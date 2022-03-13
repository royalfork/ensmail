package ensmail

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/royalfork/ensmail/pkg/ens"
	"github.com/royalfork/soltest"
)

func TestResolveEmail(t *testing.T) {
	// Before subtests run, the following setup occurs:
	// - Registry is deployed, with accts[0] owning root
	// - .eth TLD registered, owned by accts[0]
	// - public resolver deployed
	// - NewENSResolver created

	chain, accts := soltest.New()

	registryAddr, _, registry, err := ens.DeployENSRegistry(accts[0].Auth, chain)
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

	resolverAddr, _, resolver, err := ens.DeployPublicResolver(accts[0].Auth, chain, registryAddr)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	// registerLabel registers label.eth and returns the registered node.
	ethHash, err := nameHash("eth")
	if err != nil {
		t.Fatal(err)
	}
	registerLabel := func(owner common.Address, label string) ([32]byte, error) {
		lh, err := labelHash(label)
		if err != nil {
			return [32]byte{}, err
		}

		if !chain.Succeed(registry.SetSubnodeOwner(accts[0].Auth, ethHash, lh, accts[1].Addr)) {
			return [32]byte{}, errors.New("unable to register label")
		}
		return nameHash(label + ".eth")
	}

	r, err := NewENSResolver(registryAddr, chain)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("nameNotRegistered", func(t *testing.T) {
		if _, err := r.Email(context.Background(), "noexist"); err != ErrNoResolver {
			t.Errorf("want err: %s, got: %s", ErrNoResolver, err)
		}
	})

	t.Run("noResolver", func(t *testing.T) {
		label := "noresolver"

		if _, err := registerLabel(accts[1].Addr, label); err != nil {
			t.Fatal(err)
		}

		if _, err := r.Email(context.Background(), label); err != ErrNoResolver {
			t.Errorf("want err: %s, got: %s", ErrNoResolver, err)
		}
	})

	t.Run("resolverNoText", func(t *testing.T) {
		label := "badresolver"

		node, err := registerLabel(accts[1].Addr, label)
		if err != nil {
			t.Fatal(err)
		}

		if !chain.Succeed(registry.SetResolver(accts[1].Auth, node, registryAddr)) {
			t.Fatal("unable to set resolver")
		}

		if _, err := r.Email(context.Background(), label); err != vm.ErrExecutionReverted {
			t.Errorf("want err: %s, got: %s", vm.ErrExecutionReverted, err)
		}
	})

	t.Run("resolverNoEmail", func(t *testing.T) {
		label := "noemailtext"

		node, err := registerLabel(accts[1].Addr, label)
		if err != nil {
			t.Fatal(err)
		}

		if !chain.Succeed(registry.SetResolver(accts[1].Auth, node, resolverAddr)) {
			t.Fatal("unable to set resolver")
		}

		if _, err := r.Email(context.Background(), label); err != ErrNoEmail {
			t.Errorf("want err: %s, got: %s", ErrNoEmail, err)
		}
	})

	t.Run("success", func(t *testing.T) {
		label := "hasemail"
		email := "test@example.com"

		node, err := registerLabel(accts[1].Addr, label)
		if err != nil {
			t.Fatal(err)
		}

		if !chain.Succeed(registry.SetResolver(accts[1].Auth, node, resolverAddr)) {
			t.Fatal("unable to set resolver")
		}

		if !chain.Succeed(resolver.SetText(accts[1].Auth, node, "email", email)) {
			t.Fatal("unable to set resolver")
		}

		if got, err := r.Email(context.Background(), label); err != nil {
			t.Error("unexpected err:", err)
		} else if got != email {
			t.Errorf("want email: %s, got: %s", email, got)
		}
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
