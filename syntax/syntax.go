package syntax

import ()

// Syntactic Atom
type SAtom struct {
	atom     string
	optional bool
}

// Syntactic rule
type Syntax struct {
	first *SAtom
	next  []*SAtom
}
