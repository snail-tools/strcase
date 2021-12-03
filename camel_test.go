package strcase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCamelCase(t *testing.T) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test.case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
		{"ID", "ID"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := CamelCase(in)
		require.Equal(t, result, out)
	}
}

func TestCamelCaseLower(t *testing.T) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
		{"AnyKind.of-string", "anyKindOfString"},
		{"ID", "ID"},
		{"some string", "someString"},
		{" some string", "someString"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := CamelCaseLower(in)
		require.Equal(t, result, out)
	}
}
