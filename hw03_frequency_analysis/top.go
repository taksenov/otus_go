// Package hw03frequencyanalysis -- Otus Go. HW03.
package hw03frequencyanalysis

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	maximumWords int = 10
)

type wordData struct {
	count int
	word  string
}

// Top10 -- подсчет частоты появления слов в тексте.
func Top10(s string) []string {
	words := strings.Fields(s)
	cache := map[string]wordData{}
	warnCache := make(map[string]wordData, maximumWords)

	for _, val := range words {
		if checkExist(&cache, val) {
			t := get(&cache, val)
			t.count++
			add(&cache, val, t)
			handleWarnCache(&warnCache, val, t)
		} else {
			t := wordData{count: 1, word: val}
			add(&cache, val, t)
			handleWarnCache(&warnCache, val, t)
		}
	}

	res := buildSortedWords(&warnCache)

	return res
}

func checkExist(c *map[string]wordData, k string) bool {
	_, res := (*c)[k]
	return res
}

func add(c *map[string]wordData, k string, v wordData) {
	(*c)[k] = v
}

func get(c *map[string]wordData, k string) wordData {
	res := (*c)[k]
	return res
}

func handleWarnCache(c *map[string]wordData, k string, v wordData) {
	add(c, k, v)

	fmt.Println("handleWarnCache : len((*c)) = ", len((*c)))

	if len((*c)) > maximumWords {
		wordWithMinCount := wordData{count: int(math.Inf(1))}
		for _, val := range *c {
			if val.count < wordWithMinCount.count {
				wordWithMinCount = wordData{count: val.count, word: val.word}
			}
		}
		delete((*c), wordWithMinCount.word)

		fmt.Println("2 handleWarnCache : len((*c)) = ", len((*c)), "WORD =", wordWithMinCount.word)
		value, ok := (*c)[wordWithMinCount.word]
		fmt.Println("2 handleWarnCache : VALUE = ", value, "OK =", ok)
	}
}

func buildSortedWords(c *map[string]wordData) []string {
	if len((*c)) == 0 {
		res := make([]string, 0)
		return res
	}

	tS := make([]wordData, 0)

	for _, val := range *c {
		tS = append(tS, val)
	}

	sort.Slice(tS, func(i, j int) bool {
		cond := tS[i].count == tS[j].count
		if cond {
			return tS[i].word < tS[j].word
		}

		return tS[i].count > tS[j].count
	})

	res := make([]string, int(math.Abs(float64(maximumWords-(maximumWords+len((*c)))))))

	for idx, val := range tS {
		res[idx] = val.word
	}

	return res
}
