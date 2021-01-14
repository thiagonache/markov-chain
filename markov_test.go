package markov_test

import (
	"markov"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildChain(t *testing.T) {
	input := "a b a c"
	want := markov.Chain{
		"": []string{"a"},
		"a": []string{"b", "c"},
		"b": []string{"a"},
	}
	got := markov.BuildChain(strings.NewReader(input))
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGenerate(t *testing.T) {
	chain := markov.Chain{
		"": []string{"a"},
		"a": []string{"b", "c"},
		"b": []string{"a"},
	}
	want := "a c a "
	got := markov.Generate(chain, 4)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}