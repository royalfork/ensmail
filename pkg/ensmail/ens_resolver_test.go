package ensmail

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/royalfork/ensmail/pkg/ens"
)

type Thing struct{}

func TestResolveEmail(t *testing.T) {
	testENS, err := ens.NewTest()
	if err != nil {
		t.Fatal(err)
	}

	r, err := NewENSResolver(testENS.RegistryAddr, testENS.Chain)
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

		if _, err := testENS.Register(testENS.Accts[1].Addr, label); err != nil {
			t.Fatal(err)
		}

		if _, err := r.Email(context.Background(), label); err != ErrNoResolver {
			t.Errorf("want err: %s, got: %s", ErrNoResolver, err)
		}
	})

	t.Run("resolverNoText", func(t *testing.T) {
		label := "badresolver"

		node, err := testENS.Register(testENS.Accts[1].Addr, label)
		if err != nil {
			t.Fatal(err)
		}

		if !testENS.Chain.Succeed(testENS.Registry.SetResolver(testENS.Accts[1].Auth, node, testENS.RegistryAddr)) {
			t.Fatal("unable to set resolver")
		}

		if _, err := r.Email(context.Background(), label); err != vm.ErrExecutionReverted {
			t.Errorf("want err: %s, got: %s", vm.ErrExecutionReverted, err)
		}
	})

	t.Run("resolverNoEmail", func(t *testing.T) {
		label := "noemailtext"

		node, err := testENS.Register(testENS.Accts[1].Addr, label)
		if err != nil {
			t.Fatal(err)
		}

		if !testENS.Chain.Succeed(testENS.Registry.SetResolver(testENS.Accts[1].Auth, node, testENS.ResolverAddr)) {
			t.Fatal("unable to set resolver")
		}

		if _, err := r.Email(context.Background(), label); err != ErrNoEmail {
			t.Errorf("want err: %s, got: %s", ErrNoEmail, err)
		}
	})

	t.Run("success", func(t *testing.T) {
		label := "hasemail"
		email := "test@example.com"

		node, err := testENS.Register(testENS.Accts[1].Addr, label)
		if err != nil {
			t.Fatal(err)
		}

		if !testENS.Chain.Succeed(testENS.Registry.SetResolver(testENS.Accts[1].Auth, node, testENS.ResolverAddr)) {
			t.Fatal("unable to set resolver")
		}

		if !testENS.Chain.Succeed(testENS.Resolver.SetText(testENS.Accts[1].Auth, node, "email", email)) {
			t.Fatal("unable to set resolver")
		}

		if got, err := r.Email(context.Background(), label); err != nil {
			t.Error("unexpected err:", err)
		} else if got != email {
			t.Errorf("want email: %s, got: %s", email, got)
		}
	})
}
