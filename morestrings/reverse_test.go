package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, i := range cases {
		output := ReverseRunes(i.in)
		if output != i.want {
			t.Errorf("ReverseRune(%q) == %q; want %q", i.in, output, i.want)
		}
	}
}
