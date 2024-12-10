package models

type Song struct {
	ID          int64  `json:"id"`
	Group       string `json:"group"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
