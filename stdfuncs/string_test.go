package stdfuncs

import "testing"

func TestReplace(t *testing.T) {
	tests := []struct {
		name     string
		old      string
		new      string
		src      string
		expected string
	}{
		{
			name:     "simple replacement",
			old:      "world",
			new:      "Go",
			src:      "hello world",
			expected: "hello Go",
		},
		{
			name:     "multiple occurrences",
			old:      "a",
			new:      "b",
			src:      "banana",
			expected: "bbnbnb",
		},
		{
			name:     "no match",
			old:      "x",
			new:      "y",
			src:      "hello world",
			expected: "hello world",
		},
		{
			name:     "empty old string",
			old:      "",
			new:      "x",
			src:      "test",
			expected: "xtxexsxtx",
		},
		{
			name:     "empty src string",
			old:      "a",
			new:      "b",
			src:      "",
			expected: "",
		},
		{
			name:     "replace with empty string",
			old:      "o",
			new:      "",
			src:      "hello world",
			expected: "hell wrld",
		},
		{
			name:     "replace entire string",
			old:      "test",
			new:      "replacement",
			src:      "test",
			expected: "replacement",
		},
		{
			name:     "case sensitive",
			old:      "HELLO",
			new:      "hi",
			src:      "hello HELLO",
			expected: "hello hi",
		},
		{
			name:     "special characters",
			old:      ".",
			new:      "!",
			src:      "hello.world.test",
			expected: "hello!world!test",
		},
		{
			name:     "whitespace replacement",
			old:      " ",
			new:      "_",
			src:      "hello world test",
			expected: "hello_world_test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := replace(tt.old, tt.new, tt.src)
			if result != tt.expected {
				t.Errorf("replace(%q, %q, %q) = %q; want %q",
					tt.old, tt.new, tt.src, result, tt.expected)
			}
		})
	}
}
