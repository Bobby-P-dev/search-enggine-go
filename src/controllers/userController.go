package controllers

import (
	"log"
	"net/http"

	"github.com/Bobby-P-dev/search-enggine-go/src/dtos"
	"github.com/Bobby-P-dev/search-enggine-go/src/helpers"
	"github.com/Bobby-P-dev/search-enggine-go/src/models"
	"github.com/Bobby-P-dev/search-enggine-go/src/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserController interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	UR repositories.UserRepository
}

func NewUserControllerImpl(db *pgxpool.Pool, ur repositories.UserRepository) UserController {
	return &UserControllerImpl{UR: ur}
}

func (u *UserControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method tidak diizinkan, silakan gunakan GET")
		return
	}

	users, err := u.UR.GetUser()
	if err != nil {
		log.Printf("Gagal query database: %v\n", err)
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Gagal Mengambil Data")
		return
	}

	helpers.SuccesResponse(w, http.StatusOK, "Berhasil Mengambil Data", users)
}

func (u *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorResponse(w, http.StatusMethodNotAllowed, "Method tidak diizinkan, silakan gunakan POST")
		return
	}

	var req dtos.StoreUserRequestDto
	if err := helpers.ReadRequest(w, r, &req); err != nil {
		log.Printf("Gagal membaca request: %v\n", err)
		helpers.ErrorResponse(w, http.StatusBadRequest, "Gagal Membaca Request")
		return
	}

	result, err := u.UR.CreateUser(r.Context(), models.User{
		Nik:        req.Nik,
		Name:       req.Name,
		Departemen: req.Departemen,
		Position:   req.Position,
	})
	if err != nil {
		log.Printf("Gagal query database: %v\n", err)
		helpers.ErrorResponse(w, http.StatusInternalServerError, "Gagal Mengambil Data")
		return
	}

	res := dtos.GetUserResponseDto{
		ID:         result.ID,
		Nik:        result.Nik,
		Name:       result.Name,
		Departemen: result.Departemen,
		Position:   result.Position,
	}

	helpers.SuccesResponse(w, http.StatusCreated, "Berhasil Menambahkan Data", res)
}
