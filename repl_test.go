package main
import(
	"testing"
)

func TestCleanInput(t *testing.T) {

	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	
}
	// cs CASE
	for _, cs := range cases {
		actual := CleanInput(cs.input)
		
		if len(actual) != len(cs.expected) {
			t.Errorf("Length does not match expected length!")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := cs.expected[i]

			if word != expectedWord {
				t.Errorf("Word does not match to %s", expectedWord)
			}

		}
}

}