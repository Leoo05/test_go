package entities

type Professor struct {
	IDProfessor uint   `json:"idProfessor"`
	Name        string `json:"name"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Document    int    `json:"document"`
	Phone       int    `json:"phone"`
}
