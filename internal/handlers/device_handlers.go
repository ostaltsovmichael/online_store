package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"online_shop/internal/middlewares/chec_id"
	"online_shop/internal/services/device_services"

	"github.com/go-chi/chi/v5"
)

var deviceModel = device_services.Device{}
var deviceInfoModel = device_services.DeviceInfo{}
var deviceTypeBrand = device_services.DeviceTypeBrand{}

func CreateDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewDecoder(r.Body).Decode(&deviceModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Info("Плохой запрос")
		log.Println(err)
	}

	newDevice, err := deviceModel.CreateDevice()
	if err != nil {
		slog.Error("Ошибка при создании девайса")
	}
	_ = newDevice
	json.NewEncoder(w).Encode(deviceModel)
}

func GetDeviceByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := chi.URLParam(r, "id")

	newDevice, err := deviceModel.GetDeviceByID(id)
	if err != nil {
		slog.Error(fmt.Sprint(err))
	}

	param := chec_id.Param_chec_id{
		Id: id,
		ModelId: newDevice.Id,
		Name: "device",
		Writer: w,
		ModelSecsees: newDevice,
	}
	chec_id.ChecId(param)



}


func GetDeviceBrandTypeByIdDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := chi.URLParam(r, "id")
	newDevice, err := deviceTypeBrand.GetDeviceTypeBrandByID(id)
	if err != nil {
		slog.Error("Не удается получить девайс по id")
		log.Println(err)
	}
	json.NewEncoder(w).Encode(newDevice)
	slog.Info("Девайс GetDeviceBrandTypeByIdDevice получен по id: " + id)

}

func CreateDiviceInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewDecoder(r.Body).Decode(&deviceInfoModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Info("Плохой запрос")
		log.Println(err)
	}

	newDevice, err := deviceInfoModel.CreateDiviceInfo()
	if err != nil {
		slog.Error("Ошибка при создании информации о девайсе")
	}
	_ = newDevice
	json.NewEncoder(w).Encode(deviceInfoModel)
}
