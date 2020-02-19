//+build !js

package greeting_test

import (
	"fmt"
	"testing"

	"github.com/crhntr/goes/examples/greeting"
)

func TestReverse(t *testing.T) {
	table := []struct {
		Message, Reversed string
	}{
		{"01", "10"},
		{"abcΩdef", "fedΩcba"},
		{"a", "a"},
		{"Ω", "Ω"},
		{"😈", "😈"},
	}

	test := func(msg, rev string) func(*testing.T) {
		return func(t *testing.T) {
			if got := greeting.Reverse(msg); got != rev {
				t.Fail()
			}
		}
	}

	for index, row := range table {
		t.Run(
			fmt.Sprintf("failed %d reverse(%q)==%q", index, row.Message, row.Reversed),
			test(row.Message, row.Reversed),
		)
	}
}
