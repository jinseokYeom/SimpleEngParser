/*
Copyright (c) 2016 Jae-Il Yeom, Jinseok Yeom

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation
the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// syntactic atoms
const (
	S  = "S"  // Sentence
	NP = "NP" // Noun Phrase
	VP = "VP" // Verb Phrase
	AP = "AP" // Adjective Phrase
	PP = "PP" // Preposition Phrase
	N  = "N"  // Noun
	V  = "V"  // Verb
	A  = "A"  // Adjectvie
	P  = "P"  // Preposition
	D  = "D"  // Determiner
)

// help
func help() {
	fmt.Printf("\n")
	fmt.Printf("SYNTASK\n")
	fmt.Printf("Example usage:\n")
	fmt.Printf("  syntask [ENGLISH SENTENCE]\n")
	fmt.Printf("  syntask -t [ENGLISH SENTENCE]\n")
	fmt.Printf("\n")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

}
