package models

type Device struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	State     string `json:"state"`
	Connected bool   `json:"connected"`
}
