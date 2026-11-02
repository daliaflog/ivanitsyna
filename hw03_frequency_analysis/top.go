package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
)

// Top возвращает слова по убыванию частоты и алфавиту.
func Top(text string, count ...int) []string {
	return topWords(strings.Fields(text), count...)
}

// TopNormalized игнорирует регистр и краевую пунктуацию (кроме повторных дефисов).
func TopNormalized(text string, count ...int) []string {
	words := strings.Fields(text)
	for i, word := range words {
		if len(word) > 1 && strings.Trim(word, "-") == "" {
			continue
		}
		words[i] = strings.ToLower(strings.TrimFunc(word, unicode.IsPunct))
	}

	return topWords(words, count...)
}

func topWords(words []string, count ...int) []string {
	limit := 10
	if len(count) > 0 {
		limit = count[0]
	}

	if limit <= 0 {
		return []string{}
	}

	counts := make(map[string]int, len(words))
	for _, word := range words {
		if word != "" {
			counts[word]++
		}
	}

	result := make([]string, 0, len(counts))
	for word := range counts {
		result = append(result, word)
	}
	sort.Slice(result, func(i, j int) bool {
		if counts[result[i]] == counts[result[j]] {
			return result[i] < result[j]
		}
		return counts[result[i]] > counts[result[j]]
	})
	if len(result) > limit {
		result = result[:limit]
	}
	return result
}
