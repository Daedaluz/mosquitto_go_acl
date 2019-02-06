package main

import(
	"testing"
)

type SubCase struct {
	Topic string
	Subsc string
	Expect bool
}

var SubCases = []SubCase{
	{"hello/world", "hello/world", true},
	{"hello/world", "#", true},
	{"hello/world", "+/world", true},
	{"hello/world", "hello/+", true},
	{"hello/world", "hello/#", true},
	{"hello/world", "hello/world/#", true},
	{"///", "///", true},
	{"///", "///#", true},
	{"///", "//#", true},

	{"hello/world", "hello/world/toolong", false},
	{"hello/world", "hello/world/+", false},
	{"hello/world", "hello/world/", false},
	{"hello/world", "hello", false},
	{"hello/world", "hello/world/hello", false},
	{"hello/world", "hello/", false},
	{"hello/world", "hello/mismatch", false},
	{"hello/world", "hello/+/mismatch", false},
	{"hello/world", "hello/+/mismatch", false},

	{"/", "", false},
	{"", "/", false},
	{"///", "//", false},
	{"///", "////", false},
}

func TestSubs(t *testing.T) {
	for _, test := range SubCases {
		result := Match(test.Topic, test.Subsc)
		if result != test.Expect {
			t.Errorf("\"%s\" - \"%s\" should return %v", test.Topic, test.Subsc, test.Expect)
		}
	}
}
