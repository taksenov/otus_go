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
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return res, nil
}

type users [100_000]User

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
			num := result[strings.ToLower(strings.SplitN(u.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(u.Email, "@", 2)[1])] = num
		}

	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
