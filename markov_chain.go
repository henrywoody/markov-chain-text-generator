package main

import (
	"math/rand"
	"regexp"
	"strings"
)

type MarkovChain struct {
	wordCounts        map[string]map[string]int
	wordProbabilities map[string]map[string]float64
}

func NewMarkovChain() *MarkovChain {
	return &MarkovChain{
		wordCounts: make(map[string]map[string]int),
	}
}

func (m *MarkovChain) Sentence() string {
	return m.SentenceStartingWith("")
}

func (m *MarkovChain) SentenceStartingWith(firstWord string) string {
	words := []string{}
	if firstWord != "" {
		words = append(words, firstWord)
	}
	for nextWord := m.nextWord(firstWord); nextWord != ""; nextWord = m.nextWord(nextWord) {
		words = append(words, nextWord)
	}
	return strings.Join(words, " ") + "."
}

func (m *MarkovChain) nextWord(word string) string {
	if probabilities, ok := m.wordProbabilities[word]; ok {
		roll := rand.Float64()
		totalProbability := 0.0
		for nextWord, probability := range probabilities {
			totalProbability += probability
			if roll < totalProbability {
				return nextWord
			}
		}
	}
	return ""
}

var sentenceSplitRe = regexp.MustCompile(`\.\s*`)
var spaceRe = regexp.MustCompile(`\s+`)

func (m *MarkovChain) ReadText(text string) {
	sentences := sentenceSplitRe.Split(text, -1)
	for _, sentence := range sentences {
		cleanSentence := strings.ReplaceAll(sentence, ",", "")
		if cleanSentence == "" {
			continue
		}
		words := spaceRe.Split(cleanSentence, -1)

		for i, word := range words {
			var precedingWord string
			if i == 0 {
				precedingWord = ""
			} else {
				precedingWord = words[i-1]
			}

			m.addWord(precedingWord, word)
		}

		m.addWord(words[len(words)-1], "")
	}
}

func (m *MarkovChain) MakeProbabilities() {
	m.wordProbabilities = make(map[string]map[string]float64)
	for word, counts := range m.wordCounts {
		total := 0
		for _, count := range counts {
			total += count
		}

		probabilities := make(map[string]float64, len(counts))
		for nextWord, count := range counts {
			probabilities[nextWord] = float64(count) / float64(total)
		}
		m.wordProbabilities[word] = probabilities
	}
}

func (m *MarkovChain) addWord(word, nextWord string) {
	counts, ok := m.wordCounts[word]
	if !ok {
		counts = make(map[string]int)
	}

	count := counts[nextWord]
	counts[nextWord] = count + 1
	m.wordCounts[word] = counts
}
