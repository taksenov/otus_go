package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	dmn := strings.ToLower(domain)
	u := &User{}
	result := make(DomainStat)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := u.UnmarshalJSON(scanner.Bytes())
		if err != nil {
			return result, err
		}

		uE := strings.ToLower(u.Email)
		if ok := strings.HasSuffix(uE, dmn); ok {
			d := strings.ToLower(strings.SplitN(u.Email, "@", 2)[1])
			result[d]++
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
