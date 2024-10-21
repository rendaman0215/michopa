package girl

type GirlDTO struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Cup       string `json:"cup"`
	Hip       int    `json:"hip"`
	Height    int    `json:"height"`
}

func ToDTO(girl Girl) GirlDTO {
	return GirlDTO{
		ID:        girl.Id().Value(),
		Firstname: girl.Name().Firstname(),
		Lastname:  girl.name.Lastname(),
		Age:       girl.Age().Value(),
		Cup:       girl.Cup().Value(),
		Hip:       girl.Hip().Value(),
		Height:    girl.Height().Value(),
	}
}
