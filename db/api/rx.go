package api

import (
	"net/http"
	"time"

	db "github.com/ZaressaR/Capstone/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPatientInput struct {
	FirstName    string    `json:"firstName" binding:"required"`
	LastName     string    `json:"lastName" binding:"required"`
	RxName       string    `json:"rxname" binding:"required"`
	Administered time.Time `form: "administered" binding:"required" time_format:"weekday"`
}

func (server *Server) createPatientProfile(c *gin.Context) {
	var input createPatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	arg := db.CreatePatientParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	patient, err := server.RX.CreatePatient(c.Request.Context(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (server *Server) createMedication(c *gin.Context) {
	var input createPatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	arg := db.CreateMedicationParams{
		RxName:       input.RxName,
		Administered: input.Administered,
	}
	medication, err := server.RX.CreateMedication(c.Request.Context(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, medication)
}

func (server *Server) deletePatient(c *gin.Context) {
	var input createPatientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	arg := db.DeletePatientParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	err := server.RX.DeletePatient(c.Request.Context(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Patient Deleted")
}
