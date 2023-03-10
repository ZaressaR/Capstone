package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"time"

	db "github.com/ZaressaR/Capstone/db/sqlc"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql:///patient_profile?sslmode=disable"
	dbName        = "patient_profile"
	serverAddress = "localhost:3030"
)

var ctx = context.Background()

type PatientForm struct {
	FirstName string
	LastName  string
}

type MedicationForm struct {
	No     int
	RxName string
	Date   time.Time
	Action string
	Status string
}

type MedicationListData struct {
	Patient    db.Patient
	Medication []db.Medication
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

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	r.LoadHTMLGlob("templates/*.gohtml")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.gohtml", &PatientForm{}) //&PatientForm is where the data is stored
	})

	r.POST("/patient", func(c *gin.Context) {
		rxname := c.PostForm("rxname")
		firstName := c.PostForm("firstName")
		lastName := c.PostForm("lastName")
		administeredStr := c.PostForm("date")
		administered, err := time.Parse("Monday", administeredStr)

		medication := db.CreateMedicationParams{
			RxName:       rxname,
			Administered: administered,
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

		data := struct {
			Patient      db.Patient
			Medication   []db.Medication
			Submitted    bool
			RxName       string
			Administered time.Time
		}{
			Patient:      patientRecord,
			Medication:   []db.Medication{medicationRecord},
			Submitted:    true,
			RxName:       rxname,
			Administered: administered,
		}
		c.HTML(200, "medicationList.gohtml", data)
	})

	r.DELETE("/patient/:firstName", func(c *gin.Context) {
		method := c.Query("_method")
		if method == "DELETE" {
			firstName := c.Param("firstName")
			err := rx.DeletePatient(ctx, firstName)
			if err != nil {
				c.AbortWithError(500, err)

				return
			}
			method := c.Query("_method")
			if method == "DELETE" {

				c.HTML(200, "success.gohtml", method)

			}
		}
	})

	if err := r.Run(serverAddress); err != nil {
		log.Fatal("cannot start server:", err)

	}
}
