package manager

import (
	"fmt"
	"regexp"
	"strings"
)

type NewsTemplate struct {
	Type              string
	CoinSymbolRegexes []string
	ExpectedWords     []string
	MatchPercentage   float32
	MatchedIndex      int
}

func (nt *NewsTemplate) IsTextMatchWithTemplate(text string) bool {
	return strings.TrimSpace(text) != "" && nt.IsTextContainsExpectedWords(text) && nt.IsTextContainsCoinSymbol(text)
}

func (nt *NewsTemplate) IsTextContainsCoinSymbol(text string) bool {
	for i, regex := range nt.CoinSymbolRegexes {
		isMatched, err := regexp.MatchString(regex, text)

		if err != nil {
			fmt.Println(err)
		}

		if isMatched {
			nt.MatchedIndex = i
			return true
		}
	}

	return false
}

func (nt *NewsTemplate) IsTextContainsExpectedWords(text string) bool {
	wordsPresenceCount := 0

	for _, word := range nt.ExpectedWords {
		if strings.Contains(strings.ToLower(text), strings.ToLower(word)) {
			wordsPresenceCount++
		}
	}

	return float32(wordsPresenceCount)/float32(len(nt.ExpectedWords)) >= nt.MatchPercentage
}

func (nt *NewsTemplate) GetMatchedRegex() string {
	return nt.CoinSymbolRegexes[nt.MatchedIndex]
}
