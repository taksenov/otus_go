package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	u := &User{}

	i := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		err := u.UnmarshalJSON(scanner.Bytes())
		if err != nil {
			return result, err
		}

		result[i] = (*u)
		i++
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		if ok := strings.Contains(user.Email, domain); ok {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}

	return result, nil
}
