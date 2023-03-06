package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql:///patient_profile?sslmode=disable"
	dbName   = "patient_profile"
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
	defer conn.Close()

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
		medicationRecord, err := db.CreateMedication(ctx, medication)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		patientRecord, err := db.CreatePatient(ctx, db.CreatePatientParams{
			FirstName: firstName,
			LastName:  lastName,
		})
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.HTML(200, "success.gohtml", medicationRecord)
		c.HTML(200, "success.gohtml", patientRecord)
	})
	r.Run(":8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
