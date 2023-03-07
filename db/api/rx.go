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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	arg := db.CreatePatientParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	patient, err := server.RX.CreatePatient(c.Request.Context(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, patient)
}

func (server *Server) createMedicationProfile(c *gin.Context) {

	var input createPatientInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	arg2 := db.CreateMedicationParams{
		RxName:       input.RxName,
		Administered: input.Administered,
	}

	medication, err := server.CreateMedication(c.Request.Context(), arg2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, medication)
}
