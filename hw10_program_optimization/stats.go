package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	res, err := getSomethingGrandlyOptimized(r, domain)
	if err != nil {
		// NB: `get users error` -- вывод ошибки становится бессмысленным.
		// Оставлено, чтобы не нарушать сигнатуру функции. 
		return nil, fmt.Errorf("get users error: %w", err) 
	}
	return res, nil
}

func getSomethingGrandlyOptimized(r io.Reader, domain string) (DomainStat, error) {
	u := &User{}
	result := make(DomainStat)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := u.UnmarshalJSON(scanner.Bytes())
		if err != nil {
			return result, err
		}

		if ok := strings.Contains(u.Email, domain); ok {
			d := strings.ToLower(strings.SplitN(u.Email, "@", 2)[1])
			result[d] = result[d] + 1
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
