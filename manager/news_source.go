package manager

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NewsSource struct {
	Name                 string
	ApiUrl               string
	NewsTemplates        []NewsTemplate
	ResponsePreprocessor func(res *http.Response) ([]SimpleLNews, error)
}

func (ns *NewsSource) ExtractLatestNews() ([]LatestNews, error) {
	simpleNews, err := ns.RequestNews()

	if err != nil {
		return nil, err
	}

	matchedNews := ns.FindMatchedNewsWithCoinSymbols(simpleNews)

	j, _ := json.Marshal(matchedNews)

	fmt.Println(string(j))

	if err != nil {
		return nil, err
	}

	latestNews, err := ns.MapToLatestNews(matchedNews)

	if err != nil {
		return nil, err
	}

	return latestNews, nil
}

func (ns *NewsSource) RequestNews() ([]SimpleLNews, error) {
	res, err := http.Get(ns.ApiUrl)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ns.ResponsePreprocessor(res)
}

func (ns *NewsSource) FindMatchedNewsWithCoinSymbols(news []SimpleLNews) []SimpleLNews {
	matchedNews := make([]SimpleLNews, 0)

	for _, item := range news {
		if matchedTemplate := ns.FindTemplateMatch(&item); matchedTemplate != nil {
			item.Platform = ns.Name

			// TODO if coin symbol is empty, no need to include

			item.CoinSymbols = ns.ExtractCoinSymbols(&item, matchedTemplate)
			item.Type = matchedTemplate.Type
			item.Source = "poo"
			matchedNews = append(matchedNews, item)
		}
	}

	return matchedNews
}

func (ns *NewsSource) FindTemplateMatch(newsItem *SimpleLNews) *NewsTemplate {
	for _, template := range ns.NewsTemplates {
		if template.IsTextMatchWithTemplate(newsItem.Text) {
			return &template
		}
	}

	return nil
}

func (ns *NewsSource) ExtractCoinSymbols(news *SimpleLNews, template *NewsTemplate) []string {
	scrapper := CoinSymbolScrapper{
		MatchRegex:   template.GetMatchedRegex(),
		ReplaceRegex: ReplaceRegex,
	}

	return scrapper.ExtractCoinSymbols(news.Text)
}

func (ns *NewsSource) MapToLatestNews(news []SimpleLNews) ([]LatestNews, error) {
	return nil, nil
}
