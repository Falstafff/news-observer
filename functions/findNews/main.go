package main

import "github.com/Projects/news-observer-go/manager"

func main() {
	newsManager := &manager.NewsManager{
		NewsSources: []manager.NewsSource{
			manager.NewsBithumbSource(),
			manager.NewUpbitSource(),
			manager.NewKucoinSource(),
		},
	}

	newsManager.GetLastNews()
}
