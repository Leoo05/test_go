package entities

type Student struct {
	IDStudent uint   `json:"idStudent,omitempty"`
	Name      string `json:"name"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Document  int    `json:"document,omitempty"`
	Phone     int    `json:"phone,omitempty"`
	Address   string `json:"address,omitempty"`
}
