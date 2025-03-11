package go_project_test

import "testing"

var testCases = map[string]struct {
	input string
}{
	"empty string": {
		input: "",
	},
	"non-empty string": {
		input: "hello",
	},
	"empty string with spaces": {
		input: "   ",
	},
	"non-empty string with spaces": {
		input: "  hello  ",
	},
	"empty string with tabs": {
		input: "\t\t\t",
	},
	"non-empty string with tabs": {
		input: "\t\thello\t\t",
	},
	"empty string with newlines": {
		input: "\n\n\n",
	},
	"non-empty string with newlines": {
		input: "\n\nhello\n\n",
	},
	"empty string with mixed whitespace": {
		input: " \t\n \t\n",
	},
	"non-empty string with mixed whitespace": {
		input: " \t\nhello \t\n",
	},
}

// Benchmark using direct string comparison.
func BenchmarkDirectComparison(b *testing.B) {
	for name, tc := range testCases {
		b.Run(name, func(b *testing.B) {
			s := tc.input
			for i := 0; i < b.N; i++ {
				_ = s == ""
			}
		})
	}
}

// Benchmark using length check.
func BenchmarkLengthComparison(b *testing.B) {
	for name, tc := range testCases {
		b.Run(name, func(b *testing.B) {
			s := tc.input
			for i := 0; i < b.N; i++ {
				_ = len(s) == 0
			}
		})
	}
}
