package manager

import (
	"regexp"
)

type CoinSymbolScrapper struct {
	MatchRegex   string
	ReplaceRegex string
}

func (ts *CoinSymbolScrapper) ExtractCoinSymbols(text string) []string {
	matchRegex := regexp.MustCompile(ts.MatchRegex)
	replaceRegex := regexp.MustCompile(ts.ReplaceRegex)
	potentialSymbolMatches := matchRegex.FindAllString(text, -1)
	finalSymbolMatches := make([]string, 0)

	for _, potentialSymbol := range potentialSymbolMatches {
		finalSymbol := replaceRegex.ReplaceAllString(potentialSymbol, "")

		if finalSymbol != "" {
			finalSymbolMatches = append(finalSymbolMatches, finalSymbol)
		}
	}

	return finalSymbolMatches
}
