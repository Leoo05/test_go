package entities

type Student struct {
	IDStudent uint   `json:"idStudent"`
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Document  int    `json:"document"`
	Phone     int    `json:"phone"`
	Address   string `json:"address"`
}
