package ens

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

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
			out, err := NameHash(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if exp := common.HexToHash(test.hexOutput); bytes.Compare(out[:], exp[:]) != 0 {
				t.Errorf("namehash(%s): want: %x, got: %x", test.input, exp, out)
			}
		})
	}
}
