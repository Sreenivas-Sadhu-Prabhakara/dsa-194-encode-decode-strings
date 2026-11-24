package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Encode packs any list of strings into one string: each item is written as
// <length>#<bytes>. Length prefixes make delimiters inside the data harmless.
func Encode(strs []string) string {
	var b strings.Builder
	for _, s := range strs {
		b.WriteString(strconv.Itoa(len(s)))
		b.WriteByte('#')
		b.WriteString(s)
	}
	return b.String()
}

// Decode reverses Encode exactly: read the length up to '#', then take that
// many bytes verbatim — never scanning the payload for separators.
func Decode(s string) []string {
	out := []string{}
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		start := j + 1
		out = append(out, s[start:start+length])
		i = start + length
	}
	return out
}

func main() {
	packed := Encode([]string{"neet", "co#de", "", "4#x"})
	fmt.Printf("%q\n", packed)
	fmt.Printf("%q\n", Decode(packed))
}
