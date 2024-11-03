package model

type Item struct {
	Name        string `json:"name"`
	Code        int    `json:"code"`
	PostingDate string `json:"posting_date"`
}
