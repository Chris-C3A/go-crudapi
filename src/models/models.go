package models

type Book struct {
	ID      int    `json:"$id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Msg struct {
	Text string `json:"text"`
}
