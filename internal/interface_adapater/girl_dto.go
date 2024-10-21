package dto

type GirlDTO struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Cup       string `json:"cup"`
	Hip       int    `json:"hip"`
	Height    int    `json:"height"`
}
