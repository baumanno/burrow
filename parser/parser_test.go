package parser

import "testing"

func TestParseLine(t *testing.T) {
	type expected struct {
		Type     EntryType
		UserName string
		Selector string
		Server   string
		Port     string
	}
	tests := map[string]struct {
		input    string
		expected *expected
	}{
		// canonical types as defined in https://tools.ietf.org/html/rfc1436
		"text file": {
			input: "0Does this gopher menu look correct?\t/gopher/proxy\tgopher.floodgap.com\t70",
			expected: &expected{
				Type:     TextFile,
				UserName: "Does this gopher menu look correct?",
				Selector: "/gopher/proxy",
				Server:   "gopher.floodgap.com",
				Port:     "70",
			},
		},
		"directory": {
			input: "1Super-Dimensional Fortress: SDF Gopherspace\t\tsdf.org\t70",
			expected: &expected{
				Type:     Directory,
				UserName: "Super-Dimensional Fortress: SDF Gopherspace",
				Selector: "",
				Server:   "sdf.org",
				Port:     "70",
			},
		},
		"CSO phone-book server": {
			input: "2Floodgap CSO/ph phonebook server\t\tgopher.floodgap.com\t105",
			expected: &expected{
				Type:     PhoneBook,
				UserName: "Floodgap CSO/ph phonebook server",
				Selector: "",
				Server:   "gopher.floodgap.com",
				Port:     "105",
			},
		},
		"Index-Search server": {
			input: "7Search Veronica-2\t/v2/vs\tgopher.floodgap.com\t70",
			expected: &expected{
				Type:     IndexSearch,
				UserName: "Search Veronica-2",
				Selector: "/v2/vs",
				Server:   "gopher.floodgap.com",
				Port:     "70",
			},
		},

		// non-canonical, but broadly supported types
		"HTML file": {
			input: "hFloodgap.com (Web pages)\tURL:http://www.floodgap.com/\tgopher.floodgap.com\t70",
			expected: &expected{
				Type:     Html,
				UserName: "Floodgap.com (Web pages)",
				Selector: "URL:http://www.floodgap.com/",
				Server:   "gopher.floodgap.com",
				Port:     "70",
			},
		},
		"informational message": {
			input: "iWelcome to Floodgap Systems' official gopher server.\t\terror.host\t1",
			expected: &expected{
				Type:     Info,
				UserName: "Welcome to Floodgap Systems' official gopher server.",
				Selector: "",
				Server:   "error.host",
				Port:     "1",
			},
		},
		// custom types used in this client implementation
		"invalid line": {
			input: "XWelcome to Floodgap Systems' official gopher server.\t\terror.host\t1",
			expected: &expected{
				Type:     Invalid,
				UserName: "XWelcome to Floodgap Systems' official gopher server.\t\terror.host\t1",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := New(test.input)
			res := p.NextLine()

			if res.Type != test.expected.Type {
				t.Errorf("wrong type, want=%q, got=%q\n", test.expected.Type, res.Type)
			}

			if res.UserName != test.expected.UserName {
				t.Errorf("wrong userName, want=%q, got=%q\n", test.expected.UserName, res.UserName)
			}
			if res.Selector != test.expected.Selector {
				t.Errorf("wrong selector, want=%q, got=%q\n", test.expected.Selector, res.Selector)
			}

			if res.Server != test.expected.Server {
				t.Errorf("wrong server, want=%q, got=%q\n", test.expected.Server, res.Server)
			}

			if res.Port != test.expected.Port {
				t.Errorf("wrong port, want=%q, got=%q\n", test.expected.Port, res.Port)
			}
		})
	}

}
