package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	const MaxLength = 10

	words := strings.Fields(text)

	frequenciesByWords := map[string]int{}

	for _, word := range words {
		_, ok := frequenciesByWords[word]

		if ok {
			frequenciesByWords[word]++
		} else {
			frequenciesByWords[word] = 1
		}
	}

	wordsByFrequencies := make(map[int][]string)

	for word, frequency := range frequenciesByWords {
		wordsByFrequencies[frequency] = append(wordsByFrequencies[frequency], word)
	}

	frequenciesList := []int{}

	for frequency := range wordsByFrequencies {
		frequenciesList = append(frequenciesList, frequency)
	}

	sort.Slice(frequenciesList, func(i, j int) bool { return frequenciesList[i] > frequenciesList[j] })

	result := []string{}

	for _, freq := range frequenciesList {
		words = wordsByFrequencies[freq]
		sort.Slice(words, func(i, j int) bool { return words[i] < words[j] })

		for _, word := range words {
			result = append(result, word)

			if len(result) == MaxLength {
				return result
			}
		}
	}

	return result
}
