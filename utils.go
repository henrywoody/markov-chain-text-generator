package main

import "regexp"

var tokenRe = regexp.MustCompile(`([\p{L}\d'_-]+|[.,!?]+)`)
var punctuationRe = regexp.MustCompile(`^[.,!?]+$`)
var sentenceTerminationRe = regexp.MustCompile(`^[.!?]+$`)

func splitTextToTokens(text string) []string {
	return tokenRe.FindAllString(text, -1)
}

func isPunctuation(token string) bool {
	return token == "" || punctuationRe.MatchString(token)
}

func isSentenceTermination(token string) bool {
	return token == "" || sentenceTerminationRe.MatchString(token)
}
