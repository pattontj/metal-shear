package server 




type Vtuber struct {
	ID 			string `json:"id"` 
	Name 		string `json:"name"`
	Channel		string `json:"channel"`
	Affiliation string `json:"affiliation"`
}


type Clip struct {
	ID 			string `json:"id"`
	Link 		string `json:"link"`
	TsBegin 	string `json:"tsBegin"`
	TsEnd 		string `json:"tsEnd"`
	VtuberID 	string `json:"vtuberID"`
	Vtuber 		Vtuber `json:"vtuber"`
}
