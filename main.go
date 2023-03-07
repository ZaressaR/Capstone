package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"strings"
	"time"

	db "github.com/ZaressaR/Capstone/db/sqlc"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql:///patient_profile?sslmode=disable"
	dbName        = "patient_profile"
	serverAddress = "localhost:8080"
)

var ctx = context.Background()

type PatientForm struct {
	FirstName string
	LastName  string
}

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	rx := db.NewRX(conn)

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	})
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.gohtml")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.gohtml", &PatientForm{})
	})

	r.POST("/patient", func(c *gin.Context) {
		rxname := c.PostForm("rxname")
		rxname = strings.Replace(rxname, " ", "", -1)
		firstName := c.PostForm("firstName")
		lastName := c.PostForm("lastName")
		//admissionDate := c.PostForm("admissionDate")

		medication := db.CreateMedicationParams{
			RxName:       rxname,
			Administered: time.Now(),
		}
		medicationRecord, err := rx.CreateMedication(ctx, medication)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		patient := db.CreatePatientParams{
			FirstName: firstName,
			LastName:  lastName,
		}
		patientRecord, err := rx.CreatePatient(ctx, patient)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.HTML(200, "success.gohtml", medicationRecord)
		c.HTML(200, "success.gohtml", patientRecord)
	})
	r.GET("/success", func(c *gin.Context) {
		c.HTML(200, "success.gohtml", nil)
	})

	if err := r.Run(serverAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}

}
