package main

import (
	"math/rand"
)

type MarkovChain struct {
	tokenCounts        map[string]map[string]int
	tokenProbabilities map[string]map[string]float64
}

func NewMarkovChain() *MarkovChain {
	return &MarkovChain{
		tokenCounts: make(map[string]map[string]int),
	}
}

func (m *MarkovChain) Sentence() string {
	return m.SentenceStartingWith("")
}

func (m *MarkovChain) SentenceStartingWith(firstWord string) string {
	sentence := firstWord
	for nextToken := m.nextToken(firstWord); ; nextToken = m.nextToken(nextToken) {
		if len(sentence) > 0 && !isPunctuation(nextToken) {
			sentence += " "
		}
		sentence += nextToken
		if isSentenceTermination(nextToken) {
			break
		}
	}
	return sentence
}

func (m *MarkovChain) nextToken(token string) string {
	if probabilities, ok := m.tokenProbabilities[token]; ok {
		roll := rand.Float64()
		totalProbability := 0.0
		for nextToken, probability := range probabilities {
			totalProbability += probability
			if roll < totalProbability {
				return nextToken
			}
		}
	}
	return ""
}

func (m *MarkovChain) ReadText(text string) {
	tokens := splitTextToTokens(text)
	for i, token := range tokens {
		var precedingToken string
		if i > 0 {
			precedingToken = tokens[i-1]
		}
		if isSentenceTermination(precedingToken) {
			precedingToken = ""
		}

		m.addToken(precedingToken, token)

		if i == len(tokens)-1 && !isSentenceTermination(token) {
			m.addToken(token, ".")
		}
	}
}

func (m *MarkovChain) addToken(token, nextToken string) {
	counts, ok := m.tokenCounts[token]
	if !ok {
		counts = make(map[string]int)
	}

	count := counts[nextToken]
	counts[nextToken] = count + 1
	m.tokenCounts[token] = counts
}

func (m *MarkovChain) UpdateProbabilities() {
	m.tokenProbabilities = make(map[string]map[string]float64)
	for token, counts := range m.tokenCounts {
		total := 0
		for _, count := range counts {
			total += count
		}

		probabilities := make(map[string]float64, len(counts))
		for nextWord, count := range counts {
			probabilities[nextWord] = float64(count) / float64(total)
		}
		m.tokenProbabilities[token] = probabilities
	}
}
