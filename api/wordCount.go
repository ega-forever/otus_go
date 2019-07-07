package api

import (
	"sort"
	"strings"
)

type kv struct {
	key   string
	value int
}

func WordCount(text string, last int) []kv {

	var dict = make(map[string]int)

	var words = strings.Fields(text)

	for _, v := range words {
		dict[v] += 1
	}

	var wordsCountArray = make([]kv, 0, len(dict))

	for k, v := range dict {
		wordsCountArray = append(wordsCountArray, kv{k, v})
	}

	sort.Slice(wordsCountArray, func(i, j int) bool {
		return wordsCountArray[i].value > wordsCountArray[j].value
	})

	return wordsCountArray[:last]
}
