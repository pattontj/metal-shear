package server

type Streamer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Channel     string `json:"channel"`
	Affiliation string `json:"affiliation"`
}

type Clip struct {
	ID         string   `json:"id"`
	Link       string   `json:"link"`
	TsBegin    string   `json:"tsBegin"`
	TsEnd      string   `json:"tsEnd"`
	StreamerID string   `json:"streamerID"`
	Streamer   Streamer `json:"vtuber"`
}
