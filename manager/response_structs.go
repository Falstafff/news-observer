package manager

import "time"

type UpbitResponse struct {
	Success bool `json:"success"`
	Data    struct {
		TotalCount int `json:"total_count"`
		TotalPages int `json:"total_pages"`
		List       []struct {
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			ID        int       `json:"id"`
			Title     string    `json:"title"`
			ViewCount int       `json:"view_count"`
		} `json:"list"`
		FixedNotices []struct {
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			ID        int       `json:"id"`
			Title     string    `json:"title"`
			ViewCount int       `json:"view_count"`
		} `json:"fixed_notices"`
	} `json:"data"`
}

type KucoinResponse struct {
	Success   bool   `json:"success"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
	TotalNum  int    `json:"totalNum"`
	Items     []struct {
		ID             int           `json:"id"`
		Title          string        `json:"title"`
		Summary        string        `json:"summary"`
		Path           string        `json:"path"`
		Tags           []interface{} `json:"tags"`
		Images         []string      `json:"images"`
		Hot            int           `json:"hot"`
		Stick          int           `json:"stick"`
		PublishAt      string        `json:"publish_at"`
		FirstPublishAt int           `json:"first_publish_at"`
		IsNew          int           `json:"is_new"`
		Categories     []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Path string `json:"path"`
		} `json:"categories"`
		PublishTs int `json:"publish_ts"`
	} `json:"items"`
}
