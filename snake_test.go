package strcase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnakeCase(t *testing.T) {
	cases := [][]string{
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"Test Case", "test_case"},
		{" Test Case", "test_case"},
		{"Test Case ", "test_case"},
		{" Test Case ", "test_case"},
		{"test", "test"},
		{"test_case", "test_case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"AnyKind of_string", "any_kind_of_string"},
		{"numbers2and55with000", "numbers_2_and_55_with_000"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := SnakeCase(in)
		require.Equal(t, result, out)
	}
}

func TestSnakeCaseScreaming(t *testing.T) {
	cases := [][]string{
		{"testCase", "TEST_CASE"},
		{"TestCase", "TEST_CASE"},
		{"Test Case", "TEST_CASE"},
		{" Test Case", "TEST_CASE"},
		{"Test Case ", "TEST_CASE"},
		{" Test Case ", "TEST_CASE"},
		{"test", "TEST"},
		{"test_case", "TEST_CASE"},
		{"Test", "TEST"},
		{"", ""},
		{"ManyManyWords", "MANY_MANY_WORDS"},
		{"manyManyWords", "MANY_MANY_WORDS"},
		{"AnyKind of_string", "ANY_KIND_OF_STRING"},
		{"numbers2and55with000", "NUMBERS_2_AND_55_WITH_000"},
		{"JSONData", "JSON_DATA"},
		{"userID", "USER_ID"},
		{"AAAbbb", "AA_ABBB"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := SnakeCaseScreaming(in)
		require.Equal(t, result, out)
	}

}

func TestCaseDelimitedScreaming(t *testing.T) {
	cases := [][]string{
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"Test Case", "test_case"},
		{" Test Case", "test_case"},
		{"Test Case ", "test_case"},
		{" Test Case ", "test_case"},
		{"test", "test"},
		{"test_case", "test_case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"AnyKind of_string", "any_kind_of_string"},
		{"numbers2and55with000", "numbers_2_and_55_with_000"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
		{"test-case", "TEST_CASE", "true"},
		{"testCase", "TEST_CASE", "true"},
		{"TestCase", "TEST_CASE", "true"},
		{"Test Case", "TEST_CASE", "true"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		screaming := false
		if len(i) >= 3 {
			screaming = "true" == i[2]
		}
		result := CaseDelimitedScreaming(in, '_', screaming)
		require.Equal(t, result, out)
	}
}

func TestKebabCase(t *testing.T) {
	cases := [][]string{
		{"testCase", "test-case"},
		{"TestCase", "test-case"},
		{"Test Case", "test-case"},
		{" Test Case", "test-case"},
		{"Test Case ", "test-case"},
		{" Test Case ", "test-case"},
		{"test", "test"},
		{"test_case", "test-case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many-many-words"},
		{"manyManyWords", "many-many-words"},
		{"AnyKind of_string", "any-kind-of-string"},
		{"numbers2and55with000", "numbers-2-and-55-with-000"},
		{"JSONData", "json-data"},
		{"userID", "user-id"},
		{"AAAbbb", "aa-abbb"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := KebabCase(in)
		require.Equal(t, result, out)
	}
}

func TestKebabCaseScreaming(t *testing.T) {
	cases := [][]string{
		{"testCase", "TEST-CASE"},
		{"TestCase", "TEST-CASE"},
		{"Test Case", "TEST-CASE"},
		{" Test Case", "TEST-CASE"},
		{"Test Case ", "TEST-CASE"},
		{" Test Case ", "TEST-CASE"},
		{"test", "TEST"},
		{"test_case", "TEST-CASE"},
		{"Test", "TEST"},
		{"", ""},
		{"ManyManyWords", "MANY-MANY-WORDS"},
		{"manyManyWords", "MANY-MANY-WORDS"},
		{"AnyKind of_string", "ANY-KIND-OF-STRING"},
		{"numbers2and55with000", "NUMBERS-2-AND-55-WITH-000"},
		{"JSONData", "JSON-DATA"},
		{"userID", "USER-ID"},
		{"AAAbbb", "AA-ABBB"},
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := KebabCaseScreaming(in)
		require.Equal(t, result, out)
	}
}
