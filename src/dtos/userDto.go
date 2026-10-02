package dtos

type StoreUserRequestDto struct {
	Nik        string `json:"nik" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Departemen string `json:"departemen" validate:"required"`
	Position   string `json:"position" validate:"required"`
}

type UpdateUserRequestDto struct {
	Nik        string `json:"nik" validate:"omitempty"`
	Name       string `json:"name" validate:"omitempty"`
	Departemen string `json:"departemen" validate:"omitempty"`
	Position   string `json:"position" validate:"omitempty"`
}

type GetUserResponseDto struct {
	ID         int    `json:"id"`
	Nik        string `json:"nik"`
	Name       string `json:"name"`
	Departemen string `json:"departemen"`
	Position   string `json:"position"`
}
