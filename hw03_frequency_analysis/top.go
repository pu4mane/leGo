package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}
	type wordCount struct {
		word  string
		count int
	}
	var counts []wordCount
	for word, count := range m {
		counts = append(counts, wordCount{word, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count == counts[j].count {
			return counts[i].word < counts[j].word
		}
		return counts[i].count > counts[j].count
	})

	var res []string
	for i := 0; i < len(counts) && i < 10; i++ {
		res = append(res, counts[i].word)
	}
	return res
}
