// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package triegen_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"unicode"

	"golang.org/x/text/internal/triegen"
)

const seed = 0x12345

var genWriter = ioutil.Discard

func randomRunes() map[rune]uint8 {
	rnd := rand.New(rand.NewSource(seed))
	m := map[rune]uint8{}
	for len(m) < 100 {
		// Only set our random rune if it is a valid Unicode code point.
		if r := rune(rnd.Int31n(unicode.MaxRune + 1)); []rune(string(r))[0] == r {
			m[r] = 1
		}
	}
	return m
}

// Example_build shows how to build a simple trie. It assigns the value 1 to
// 100 random runes generated by randomRunes.
func Example_build() {
	t := triegen.NewTrie("rand")

	for r, _ := range randomRunes() {
		t.Insert(r, 1)
	}
	sz, err := t.Gen(genWriter)

	fmt.Printf("Trie size: %d bytes\n", sz)
	fmt.Printf("Error:     %v\n", err)

	// Output:
	// Trie size: 9280 bytes
	// Error:     <nil>
}

// Example_lookup demonstrates how to use the trie generated by Example_build.
func Example_lookup() {
	trie := newRandTrie(0)

	// The same set of runes used by Example_build.
	runes := randomRunes()

	// Verify the right value is returned for all runes.
	for r := rune(0); r <= unicode.MaxRune; r++ {
		// Note that the return type of lookup is uint8.
		if v, _ := trie.lookupString(string(r)); v != runes[r] {
			fmt.Println("FAILURE")
			return
		}
	}
	fmt.Println("SUCCESS")

	// Output:
	// SUCCESS
}

// runeValues generates some random values for a set of interesting runes.
func runeValues() map[rune]uint64 {
	rnd := rand.New(rand.NewSource(seed))
	m := map[rune]uint64{}
	for p := 4; p <= unicode.MaxRune; p <<= 1 {
		for d := -1; d <= 1; d++ {
			m[rune(p+d)] = uint64(rnd.Int63())
		}
	}
	return m
}

// ExampleGen_build demonstrates the creation of multiple tries sharing common
// blocks. ExampleGen_lookup demonstrates how to use the generated tries.
func ExampleGen_build() {
	var tries []*triegen.Trie

	rv := runeValues()
	for _, c := range []struct {
		include func(rune) bool
		name    string
	}{
		{func(r rune) bool { return true }, "all"},
		{func(r rune) bool { return r < 0x80 }, "ASCII only"},
		{func(r rune) bool { return r < 0x80 }, "ASCII only control"},
		{func(r rune) bool { return r <= 0xFFFF }, "BMP only"},
		{func(r rune) bool { return r > 0xFFFF }, "No BMP"},
	} {
		t := triegen.NewTrie(c.name)
		tries = append(tries, t)

		for r, v := range rv {
			if c.include(r) {
				t.Insert(r, v)
			}
		}
	}
	sz, err := triegen.Gen(genWriter, "multi", tries)

	fmt.Printf("Trie size: %d bytes\n", sz)
	fmt.Printf("Error:     %v\n", err)

	// Output:
	// Trie size: 18250 bytes
	// Error:     <nil>
}

// ExampleGen_lookup shows how to look up values in the trie generated by
// ExampleGen_build.
func ExampleGen_lookup() {
	rv := runeValues()
	for i, include := range []func(rune) bool{
		func(r rune) bool { return true },        // all
		func(r rune) bool { return r < 0x80 },    // ASCII only
		func(r rune) bool { return r < 0x80 },    // ASCII only control
		func(r rune) bool { return r <= 0xFFFF }, // BMP only
		func(r rune) bool { return r > 0xFFFF },  // No BMP
	} {
		t := newMultiTrie(i)

		for r := rune(0); r <= unicode.MaxRune; r++ {
			x := uint64(0)
			if include(r) {
				x = rv[r]
			}
			// As we convert from a valid rune, we know it is safe to use
			// lookupStringUnsafe.
			if v := t.lookupStringUnsafe(string(r)); x != v {
				fmt.Println("FAILURE")
				return
			}
		}
	}
	fmt.Println("SUCCESS")

	// Output:
	// SUCCESS
}
