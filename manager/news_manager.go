package manager

import "fmt"

type NewsManager struct {
	NewsSources []NewsSource
}

func (nm *NewsManager) GetLastNews() {
	for news := range nm.GetNewsQueue() {
		fmt.Println("---------------->")
		fmt.Println(news)
	}
}

func (nm *NewsManager) GetNewsQueue() chan []LatestNews {
	newsQueue := make(chan []LatestNews)

	for _, newsSource := range nm.NewsSources {
		go func(newsSource NewsSource) {
			latestNews, err := newsSource.ExtractLatestNews()

			if err != nil {
				fmt.Println(err)
				return
			}

			newsQueue <- latestNews
		}(newsSource)
	}

	return newsQueue
}
