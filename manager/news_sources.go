package manager

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewsBithumbSource() NewsSource {
	return NewsSource{
		Name:   "bithumb",
		ApiUrl: "https://cafe.bithumb.com/view/boards/43",
		NewsTemplates: []NewsTemplate{
			{
				Type:              CoinListing,
				CoinSymbolRegexes: []string{CoinSymbolRegex},
				ExpectedWords:     []string{"마켓 추가", "이벤트", "이벤트 안내"},
				MatchPercentage:   1,
			},
		},
		ResponsePreprocessor: func(res *http.Response) ([]SimpleLNews, error) {
			doc, err := goquery.NewDocumentFromReader(res.Body)

			if err != nil {
				return nil, err
			}

			simpleNews := make([]SimpleLNews, 0)

			doc.Find("#dataTables a").Each(func(i int, selection *goquery.Selection) {
				simpleNews = append(simpleNews, SimpleLNews{
					Text: strings.TrimSpace(selection.Text()),
					Slug: "",
				})
			})

			return simpleNews, nil
		},
	}
}

func NewUpbitSource() NewsSource {
	return NewsSource{
		Name:   "upbit",
		ApiUrl: "https://api-manager.upbit.com/api/v1/notices?page=1&per_page=20",
		NewsTemplates: []NewsTemplate{
			{
				Type:              CoinListing,
				CoinSymbolRegexes: []string{CoinSymbolRegex},
				ExpectedWords:     []string{"거래", "자산", "디지털", "추가"},
				MatchPercentage:   0.75,
			},
		},
		ResponsePreprocessor: func(res *http.Response) ([]SimpleLNews, error) {
			body, err := ioutil.ReadAll(res.Body)

			if err != nil {
				return nil, err
			}

			var result UpbitResponse

			if err := json.Unmarshal(body, &result); err != nil {
				return nil, err
			}

			simpleNews := make([]SimpleLNews, len(result.Data.List))

			for i, item := range result.Data.List {
				simpleNews[i].Text = item.Title
				simpleNews[i].Slug = string(rune(item.ID))
			}

			return simpleNews, nil
		},
	}
}

func NewKucoinSource() NewsSource {
	return NewsSource{
		Name:   "kucoin",
		ApiUrl: "https://www.kucoin.com/_api/cms/articles?page=1&pageSize=3&category=listing&lang=en_US",
		NewsTemplates: []NewsTemplate{
			{
				Type:              CoinListing,
				CoinSymbolRegexes: []string{CoinSymbolRegex},
				ExpectedWords:     []string{"kucoin", "get", "list"},
				MatchPercentage:   0.8,
			},
		},
		ResponsePreprocessor: func(res *http.Response) ([]SimpleLNews, error) {
			body, err := ioutil.ReadAll(res.Body)

			if err != nil {
				return nil, err
			}

			var result KucoinResponse

			if err := json.Unmarshal(body, &result); err != nil {
				return nil, err
			}

			simpleNews := make([]SimpleLNews, len(result.Items))

			for i, item := range result.Items {
				simpleNews[i].Text = item.Title
				simpleNews[i].Slug = item.Path
			}

			return simpleNews, nil
		},
	}
}
