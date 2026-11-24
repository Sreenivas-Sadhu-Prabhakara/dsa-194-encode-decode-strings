package main

import (
	"reflect"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	cases := [][]string{
		{"neet", "code", "love", "you"},
		{"we", "say", ":", "yes"},
		{""},
		{},
		{"", "", ""},
		{"co#de", "4#x", "##"}, // delimiters inside payloads
		{"123#456"},            // digits and delimiter together
		{"a"},
	}
	for _, in := range cases {
		got := Decode(Encode(in))
		want := in
		if len(want) == 0 {
			want = []string{}
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("round trip %q -> %q", in, got)
		}
	}
}

func TestDistinctInputsDistinctEncodings(t *testing.T) {
	a := Encode([]string{"ab", "c"})
	b := Encode([]string{"a", "bc"})
	if a == b {
		t.Error("different lists must encode differently")
	}
}
