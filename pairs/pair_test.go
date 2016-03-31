package pairs

import "testing"

type pairTest struct {
	name     string
	source   []string
	expected [][]string
}

var (
	tests = []pairTest{
		pairTest{
			name:     "empty",
			source:   []string{},
			expected: [][]string{}},
		pairTest{
			name:   "even",
			source: []string{"Darth Vader", "Darth Maul", "Palpatine", "Maurice"},
			expected: [][]string{
				[]string{"Darth Vader", "Darth Maul"},
				[]string{"Palpatine", "Maurice"}}},
		pairTest{
			name:   "odd",
			source: []string{"How", "much", "wood", "would", "a", "woodchuck", "chuck"},
			expected: [][]string{
				[]string{"How", "much"},
				[]string{"wood", "would"},
				[]string{"a", "woodchuck", "chuck"}}}}
)

func TestPairs(t *testing.T) {
	for _, test := range tests {
		testPair(t, test)
	}
}

func testPair(t *testing.T, test pairTest) {
	result := GeneratePairs(test.source)
	t.Log(test.name, "expected is:", test.expected)
	t.Log(test.name, "result   is:", result)

	if len(result) != len(test.expected) {
		t.Error(test.name, "Array length are not the same", result, test.expected)
	}

	for i := range test.expected {
		for j := range test.expected[i] {
			if test.expected[i][j] != result[i][j] {
				t.Error(test.name, test.expected[i][j], "!=", result[i][j])
			}
		}
	}
}

func TestShuffle(t *testing.T) {
	src := []string{"Obi-Wan", "Han", "Luke", "Anakin"}
	result := Shuffle(src, 2)
	expectedWithSeed2 := []string{"Han", "Luke", "Anakin", "Obi-Wan"}

	if len(result) != len(src) {
		t.Error("Array length are not the same")
	}
	for i := range src {
		if result[i] != expectedWithSeed2[i] {
			t.Errorf("Shuffle is wrong, result[%v] != expected[%v]", i, i)
		}
	}
}
