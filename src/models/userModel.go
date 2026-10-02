package models

type User struct {
	ID         int    `json:"id"`
	Nik        string `json:"nik"`
	Name       string `json:"name"`
	Departemen string `json:"departemen"`
	Position   string `json:"position"`
}
