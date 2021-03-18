package controllers

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/team03/app/ent/disease"
	"github.com/team03/app/ent/gender"
	"github.com/team03/app/ent/medicalcare"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent"
	"github.com/gin-gonic/gin"
)

// PatientController defines the struct for the patient controller
type PatientController struct {
	client *ent.Client
	router gin.IRouter
}

type Patient struct {
	PatientID   string
	Name        string
	CardID      string
	Gender      int
	Tel         string
	Birthday    string
	Age         int
	Disease     int
	MedicalCare int
}

// PatientCreate handles POST requests for adding patient entities
// @Summary Create patient
// @Description Create patient
// @ID create-patient
// @Accept   json
// @Produce  json
// @Param Patient body ent.Patient true "Patient entity"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [post]
func (ctl *PatientController) PatientCreate(c *gin.Context) {
	obj := Patient{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "patient binding failed",
		})
		return 
	}

	g, err := ctl.client.Gender.
		Query().
		Where(gender.IDEQ(int(obj.Gender))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Gender not found",
		})
		return
	}

	ds, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(obj.Disease))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Disease not found",
		})
		return
	}

	mc, err := ctl.client.MedicalCare.
		Query().
		Where(medicalcare.IDEQ(int(obj.MedicalCare))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Medicalcare not found",
		})
		return
	}


	time, err := time.Parse(time.RFC3339, obj.Birthday+"T00:00:00Z")  //format time ตามรูปแบบของRFC3339 เอาแค่วันเดือนปี ไม่เอาเวลา
	p, err := ctl.client.Patient.
		Create().
		SetPatientID(obj.PatientID).
		SetName(obj.Name).
		SetCardID(obj.CardID).
		SetGender(g).
		SetTel(obj.Tel).
		SetBirthday(time).
		SetAge(obj.Age).
		SetDisease(ds).
		SetMedicalcare(mc).
		Save(context.Background())
	
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   p,
	})

}

// GetPatient handles GET requests to retrieve a patient entity
// @Summary Get a patient entity by ID
// @Description get patient by ID
// @ID get-patient
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients/{id} [get]
func (ctl *PatientController) GetPatient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		WithGender().
		WithDisease().
		WithMedicalcare().
		WithNurse().
		Where(patient.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// GetSearchPatient handles GET requests to retrieve a patient entity
// @Summary Get a Patient entity by Search
// @Description get patient by Search
// @ID get-Patient-by-search
// @Produce  json
// @Param Patient query string false "Search Patient"
// @Success 200 {object} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchpatients [get]
func (ctl *PatientController) GetSearchPatient(c *gin.Context) {
	patientsearch := c.Query("patient")

	ps, err := ctl.client.Patient.
		Query().
		WithGender().
		WithDisease().
		WithMedicalcare().
		Where(patient.PatientIDContains(patientsearch)).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": ps,
	})
}

// ListPatient handles request to get a list of patient entities
// @Summary List patient entities
// @Description list patient entities
// @ID list-patient
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Patient
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patients [get]
func (ctl *PatientController) ListPatient(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	patients, err := ctl.client.Patient.
		Query().
		WithGender().
		WithDisease().
		WithMedicalcare().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, patients)
}

// NewPatientController creates and patients handles for the patient controller
func NewPatientController(router gin.IRouter, client *ent.Client) *PatientController {
	pc := &PatientController{
		client: client,
		router: router,
	}
	pc.register()
	return pc
}

// InitUserController patient routes to the main engine
func (ctl *PatientController) register() {
	patients := ctl.router.Group("/patients")
	
	patients.POST("", ctl.PatientCreate)	
	patients.GET(":id", ctl.GetPatient)
	patients.GET("", ctl.ListPatient)

	searchpatients := ctl.router.Group("/searchpatients")
	searchpatients.GET("",ctl.GetSearchPatient)

}