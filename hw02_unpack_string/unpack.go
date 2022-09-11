// Package hw02unpackstring -- OTUS HW2 Unpack string.
package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

type runeAdvanced struct {
	isDigit bool
	val     rune
}

type runeDigit struct {
	isSet bool
	val   rune
}

type (
	stack struct {
		top    *node
		length int
	}
	node struct {
		value runeAdvanced
		prev  *node
	}
)

// createStack -- создает стек.
func createStack() *stack {
	return &stack{nil, 0}
}

// lenStk -- возвращает размер.
func (t *stack) lenStk() int {
	return t.length
}

// pop -- снимает элемент с вершины.
func (t *stack) pop() runeAdvanced {
	if t.length == 0 {
		return runeAdvanced{isDigit: true, val: rune(0)}
	}

	n := t.top
	t.top = n.prev
	t.length--
	return n.value
}

// push -- добавляет элемент на вершину.
func (t *stack) push(value runeAdvanced) {
	n := &node{value, t.top}
	t.top = n
	t.length++
}

// buildStr -- создает строку из стека рун.
func buildStr(st *stack) string {
	var b strings.Builder
	for st.lenStk() > 0 {
		b.WriteString(string(st.pop().val))
	}

	res := b.String()
	return res
}

// checkIsDigit -- проверяет является ли руна числом.
func checkIsDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// rTI -- конвертирует rune в int.
func rTI(r rune) int {
	return int(r - '0')
}

// ErrInvalidString -- custom error.
var ErrInvalidString = errors.New("invalid string")

// Unpack -- input string.
func Unpack(input string) (string, error) {
	res := ""
	runes := []rune(input)

	if len(runes) == 0 {
		return res, nil
	}

	firstRune, _ := utf8.DecodeRuneInString(input)
	if unicode.IsDigit(firstRune) {
		return "", ErrInvalidString
	}

	stack := createStack()

	for _, val := range runes {
		rA := runeAdvanced{val: val}
		rA.isDigit = checkIsDigit(rA.val)

		stack.push(rA)
	}

	sLocal := createStack()
	tmpDigit := runeDigit{isSet: false, val: rune(0)}

	for stack.lenStk() > 0 {
		rA := stack.pop()

		isDigit := checkIsDigit(rA.val)
		if isDigit && tmpDigit.isSet {
			return "", ErrInvalidString
		}

		if isDigit {
			tmpDigit = runeDigit{isSet: true, val: rA.val}
			continue
		}

		if !isDigit {
			if tmpDigit.isSet {
				c := rTI(tmpDigit.val)

				for c > 0 {
					sLocal.push(rA)
					c--
				}

				tmpDigit = runeDigit{isSet: false, val: rune(0)}
			} else {
				sLocal.push(rA)
			}
		}
	}

	res = buildStr(sLocal)
	return res, nil
}
