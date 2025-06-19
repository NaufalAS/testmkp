package jadwalcontroller

import (
	"mkptest/model"
	"mkptest/model/web"
	jadwalservice "mkptest/service/jadwal"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JadwalControllerImpl struct {
	jadwalService jadwalservice.JadwalService
}

func NewJadwalController(jadwalService jadwalservice.JadwalService) *JadwalControllerImpl {
	return &JadwalControllerImpl{
		jadwalService: jadwalService,
	}
}

func (controller *JadwalControllerImpl) JadwalController(c echo.Context) error {
	jadwal := new(web.CreateScheduleRequest)
	if err := c.Bind(jadwal); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}
	if err := c.Validate(jadwal); err != nil {
		return err
	}
	saverjadwal, errSavejadwal := controller.jadwalService.SaveJadwal(*jadwal)
	if errSavejadwal != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errSavejadwal.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Register Sukses", saverjadwal))
}

func (controller *JadwalControllerImpl) DeleteJadwal(c echo.Context) error {
	idParam := c.Param("id")
	if idParam == "" {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "ID jadwal wajib diisi", nil))
	}
	// Convert id ke int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, "ID tidak valid", nil))
	}
	// Panggil service
	err = controller.jadwalService.DeleteJadwal(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.ResponseToClient(http.StatusNotFound, false, err.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Jadwal berhasil dihapus", nil))
}


func (controller *JadwalControllerImpl) GetListJadwalController(c echo.Context) error {
	getJadwalController, errGetJadwalController := controller.jadwalService.GetJadwalList()

	if errGetJadwalController != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "error", errGetJadwalController.Error()))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK,  "berhasil melihat seluruh list User", getJadwalController))
}

func (controller *JadwalControllerImpl) GetJadwalByIdController(c echo.Context) error{
	IdUser := c.Param("id")
	id, _ := strconv.Atoi(IdUser)
	getJadwalIdController, errGetJadwalController := controller.jadwalService.GetJadwalById(id)

	if errGetJadwalController != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, "error", errGetJadwalController.Error()))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK,  "berhasil melihat satu id", getJadwalIdController))
}