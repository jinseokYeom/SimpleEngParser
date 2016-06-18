/*

Copyright (c) 2016 Jinseok Yeom, Jae-Il Yeom

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

   Welcome to Syntask!

   Example usage:

       For plain view,

           $ syntask The students are interesting.
           S(NP(D("The")N("students"))VP(V("are")AP(A("interesting"))))

       For tree view,

       $ syntask -t The Students are interesting.

           S
           NP           VP
           D   N        V   AP
                        A
           |   |        |   |
           The students are interesting.

       For this program to work, dict.synt is required.

           dict.synt format:

           [word 1] [syntactic atom 1]
           [word 2] [syntactic atom 2]
           [word 3] [syntactic atom 3]
           ...

*/

package main

import (
	"errors"
	"log"
	"os"
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

}

// load dictionary
func loadDictionary() (map[string]string, error) {
	f, err := os.Open("dict.synt")
	defer f.Close()
	if err != nil {
		return nil, err
	}

	dict := make(map[string]string)
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		line := fs.Text()
		dict[line[0]] = line[1]
	}

	return dict
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	dict, err := loadDictionary()
	if err != nil {
		log.Fatal(err)
	}

	// slice of words
	sentence := os.Args[1:]

}
