package parser

import "testing"

func TestParseLine(t *testing.T) {
	type expected struct {
		Type    EntryType
		Content string
	}
	tests := map[string]struct {
		input    string
		expected *expected
	}{
		"info line": {
			input: "iWelcome to Floodgap Systems' official gopher server.	error.host      1",
			expected: &expected{
				Type:    INFO,
				Content: "Welcome to Floodgap Systems' official gopher server.",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := New(test.input)
			res := p.Line()

			if res.Type != test.expected.Type {
				t.Errorf("wrong type, want=%q, got=%q\n", test.expected.Type, res.Type)
			}

			if res.Content != test.expected.Content {
				t.Errorf("wrong content, want=%q, got=%q\n", test.expected.Content, res.Content)
			}
		})
	}

}
