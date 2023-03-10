package api

import (
	"gorm.io/gorm"

	db "github.com/ZaressaR/Capstone/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	RX     *db.RX
	router *gin.Engine
}

func NewServer(rx *db.RX) *Server {
	server := &Server{RX: rx}
	router := gin.Default()

	router.GET("/", server.createPatientProfile)
	router.POST("/patient", server.createPatientProfile)
	router.POST("/medicationList", server.createMedication)
	router.POST("/medication", server.createMedication)
	router.GET("/medication", server.createMedication)
	router.DELETE("/patient/:firstName", server.deletePatient)

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
	r.POST("/medication", HandlerFunc)
	r.DELETE("/patient", patientPostHandler)

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

func HandlerFunc(c *gin.Context) {
	data := []db.Medication{}

	medicationData := []db.MedicationData{}
	for _, m := range data {
		md := m.Data()

		medicationData = append(medicationData, md)
	}
	c.JSON(200, medicationData)

}
