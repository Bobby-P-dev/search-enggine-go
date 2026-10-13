package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Bobby-P-dev/search-enggine-go/src/helpers"
	"github.com/Bobby-P-dev/search-enggine-go/src/models"
)

type LampPublisher interface {
	PublishLampCommand(state string) error
}

type DeviceController interface {
	GetDevice(w http.ResponseWriter, r *http.Request)
	SetLamp(w http.ResponseWriter, r *http.Request)
}

type DeviceControllerImpl struct {
	publisher LampPublisher
}

func NewDeviceControllerImpl(
	publisher LampPublisher,
) DeviceController {
	return &DeviceControllerImpl{
		publisher: publisher,
	}
}

func (d *DeviceControllerImpl) GetDevice(w http.ResponseWriter, r *http.Request) {
	device := []models.Device{
		{
			ID:        "1",
			Name:      "Device 1",
			Type:      "Device Type 1",
			State:     "State 1",
			Connected: true,
		},
	}

	helpers.SuccesResponse(w, http.StatusOK, "Berhasil Mengambil Data", device)
}

type LampCommandRequest struct {
	State string `json:"state"`
}

func (d *DeviceControllerImpl) SetLamp(
	w http.ResponseWriter,
	r *http.Request,
) {
	var input LampCommandRequest

	r.Body = http.MaxBytesReader(w, r.Body, 1024)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		helpers.ErrorResponse(
			w, http.StatusBadRequest, "Invalid JSON",
		)
		return
	}

	if input.State != "ON" && input.State != "OFF" {
		helpers.ErrorResponse(
			w, http.StatusBadRequest, "State must be ON or OFF",
		)
		return
	}

	if err := d.publisher.PublishLampCommand(input.State); err != nil {
		helpers.ErrorResponse(
			w, http.StatusServiceUnavailable,
			"Failed to publish MQTT command",
		)
		return
	}

	helpers.SuccesResponse(
		w,
		http.StatusAccepted,
		"Command published to MQTT broker",
		map[string]string{"state": input.State},
	)
}
