package api

import (
	"gorm.io/gorm"

	db "github.com/ZaressaR/Capstone/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Medication represents a row in medication table.

type Server struct {
	RX     *db.RX
	router *gin.Engine
}

func NewServer(rx *db.RX) *Server {
	server := &Server{RX: rx}
	router := gin.Default()

	router.GET("/", server.createPatientProfile)
	router.POST("/patient", server.createPatientProfile)

	server.router = router
	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorRespone(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func connectDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	r.LoadHTMLGlob("templates/*.gohtml")
	r.GET("/", patientGetHandler)
	r.POST("/patient", patientPostHandler)
	// r.Use(connectDB(db)
	// r.GET("/patients", getPatients)
	// r.GET("/patients/:id", getPatient)
	// r.POST("/patients", createPatient)
	// r.PUT("/patients/:id", updatePatient)
	// r.DELETE("/patients/:id", deletePatient)
	// r.GET("/medications", getMedications)
	// r.GET("/medications/:id", getMedication)
	// r.POST("/medications", createMedication)
	// r.PUT("/medications/:id", updateMedication)
	// r.DELETE("/medications/:id", deleteMedication)
}

func patientPostHandler(c *gin.Context) {
	person := &db.CreateMedicationParams{}
	if err := c.ShouldBindJSON(person); err != nil {
		c.JSON(400, errorRespone(err))
		return
	}
	c.JSON(200, person)

}

func patientGetHandler(c *gin.Context) {
	c.HTML(200, "index.gohtml", gin.H{})
}
