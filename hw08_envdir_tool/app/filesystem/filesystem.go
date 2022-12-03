// Package filesystem -- utilities for working with the file system.
package filesystem

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"unicode"
)

// ReadFileFirstLine -- reading first line of file.
func ReadFileFirstLine(file string) (string, error) {
	res := ""

	var err error
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	lB, _, err := r.ReadLine()
	if err == io.EOF {
		lB = handleNullish(lB)
		res = handleTailSpaces(bytes.NewBuffer(lB).String())
		return res, err
	} else if err != nil {
		return "", err
	}

	lB = handleNullish(lB)
	res = handleTailSpaces(bytes.NewBuffer(lB).String())

	return res, nil
}

func handleNullish(b []byte) []byte {
	var res []byte

	for _, v := range b {
		if v != 0x00 {
			res = append(res, v)
		} else {
			res = append(res, '\n')
		}
	}

	return res
}

func handleTailSpaces(s string) string {
	res := ""

	res = strings.TrimRightFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	return res
}
