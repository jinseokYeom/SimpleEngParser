package dict

import (
	"bufio"
	"errors"
	"os"
	"path"
	"strings"
)

// syntax dictionary
type SyntDict struct {
	dict map[string]string
}

func New() *SyntDict {
	return &SyntDict{
		dict: make(map[string]string),
	}
}

// load dictionary
func (s *SyntDict) Load(filename) error {
	if path.Ext(filename) != ".dict" {
		err := errors.New("Wrong dictionary file format")
		return err
	}

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return err
	}
	// file scanner
	fs := bufio.NewScanner(f)
	for fs.Scan() {
		line := fs.Text()
		tokens := strings.Split(line, ": ")
		if len(tokens) == 2 {
			word := strings.ToUpper(tokens[0])
			satom := strings.ToUpper(token[1])
			s.dict[word] = satom
		}
	}
	return nil
}
