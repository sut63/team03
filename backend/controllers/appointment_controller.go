package controllers

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/appointment"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/room"
)

//AppointmentController defind the struct for Appointment controller
type AppointmentController struct {
	client *ent.Client
	router gin.IRouter
}

//Appointment defind the struct for the Appointment controller
type Appointment struct {
	AppointID	string
	Patient     int
	Dentist     int
	Detail 		string
	Datetime    string
	Room   		int
	Remark		string
	
}

// AppointmentCreate handles POST requests for adding  Appointment entities
// @Summary Create Appointment
// @Description Create Appointment
// @ID create-Appointment
// @Accept   json
// @Produce  json
// @Param  Appointment body ent.Appointment true "Appointment entity"
// @Success 200 {object} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments [post]
func (ctl *AppointmentController) AppointmentCreate(c *gin.Context) {
	obj := Appointment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Appointment binding failed",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "room not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Datetime)
	ap, err := ctl.client.Appointment.
		Create().
		SetAppointID(obj.AppointID).
		SetPatient(p).
		SetRoom(r).
		SetDetail(obj.Detail).
		SetDatetime(time).
		SetRemark(obj.Remark).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   ap,
	})
}

// GetAppointment handles GET requests to retrieve a Appointment entity
// @Summary Get a Appointment entity by ID
// @Description get Appointment by ID
// @ID get-Appointment
// @Produce  json
// @Param id path int true "Appointment ID"
// @Success 200 {object} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments/{id} [get]
func (ctl *AppointmentController) GetAppointment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ap, err := ctl.client.Appointment.
		Query().
		WithDentist().
		WithPatient().
		WithRoom().
		WithNurse().
		Where(appointment.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ap)
}


// GetSearchAppointment handles GET requests to retrieve a Appointment entity
// @Summary Get a Appointment entity by Search
// @Description get Appointment by Search
// @ID get-Appointment-by-search
// @Produce  json
// @Param Appointment query string false "Search Appointment"
// @Success 200 {object} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchappointments [get]
func (ctl *AppointmentController) GetSearchAppointment(c *gin.Context) {
	appointmentsearch := c.Query("appointment")

	aps, err := ctl.client.Appointment.
		Query().
		WithPatient().
		WithRoom().
		Where(appointment.AppointIDContains(appointmentsearch)).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": aps,
	})
}

// ListAppointment handles request to get a list of Appointment entities
// @Summary List Appointment entities
// @Description list Appointment entities
// @ID list-Appointment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Appointment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /appointments [get]
func (ctl *AppointmentController) ListAppointment(c *gin.Context) {
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

	appointments, err := ctl.client.Appointment.
		Query().
		WithPatient().
		WithRoom().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, appointments)
}

// NewAppointmentController creates and registers handles for the appointment controller
func NewAppointmentController(router gin.IRouter, client *ent.Client) *AppointmentController {
	apc := &AppointmentController{
		client: client,
		router: router,
	}

	apc.register()
	return apc

}

func (ctl *AppointmentController) register() {
	appointments := ctl.router.Group("/appointments")
	appointments.GET(":id", ctl.GetAppointment)

	searchappointments := ctl.router.Group("/searchappointments")
	searchappointments.GET("",ctl.GetSearchAppointment)

	// CRUD
	appointments.POST("", ctl.AppointmentCreate)
	appointments.GET("", ctl.ListAppointment)

}
