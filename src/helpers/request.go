package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ReadRequest(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := int64(1024 * 1024 * 10)
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}

	if err := validate.Struct(data); err != nil {
		return err
	}

	return nil
}
