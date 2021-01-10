package controllers

import (
	"context"
	
	"strconv"
	"time"

	"github.com/team03/app/ent"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/degree"
	"github.com/team03/app/ent/expert"
	"github.com/team03/app/ent/gender"
	"github.com/gin-gonic/gin"
)
//DentistController defines the struct for the Dentist controller
type DentistController struct {
	client *ent.Client
	router gin.IRouter
}

type Dentist struct {
	Nurse         int
	Degree		  int
	Expert		  int
	Gender		  int
	Birthday	string
}

// CreateDentist handles POST requests for adding dentist entities
// @Summary Create dentist
// @Description Create dentist
// @ID create-dentist
// @Accept   json
// @Produce  json
// @Param dentist body Dentist true "Dentist entity"
// @Success 200 {object} Dentist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentists [post]
func (ctl *DentistController) CreateDentist(c *gin.Context) {
	obj := Dentist{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dentist binding failed",
		})
		return
	}

	dg, err := ctl.client.Degree.
		Query().
		Where(degree.IDEQ(int(obj.Degree))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "degree not found",
		})
		return
	}

	e, err := ctl.client.Expert.
	Query().
	Where(expert.IDEQ(int(obj.Expert))).
	Only(context.Background())

	if err != nil {
	c.JSON(400, gin.H{
		"error": "degree not found",
	})
	return
	}

	g, err := ctl.client.Gender.
	Query().
	Where(gender.IDEQ(int(obj.Gender))).
	Only(context.Background())

	if err != nil {
	c.JSON(400, gin.H{
		"error": "gender not found",
	})
	return
	}


	n, err := ctl.client.Nurse.
		Query().
		Where(nurse.IDEQ(int(obj.Nurse))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "nurse not found",
		})
		return
	}


	times, err := time.Parse(time.RFC3339, obj.Birthday)

	d, err := ctl.client.Dentist.
		Create().
		SetNurse(n).
		SetDegree(dg).
		SetExpert(e).
		SetGender(g).
		SetBirthday(times).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}


	c.JSON(200, gin.H{
		"status": true,
		"data":   d,
	})
}


// GetDentist handles GET requests to retrieve a dentist entity
// @Summary Get a dentist entity by ID
// @Description get dentist by ID
// @ID get-dentist
// @Produce  json
// @Param id path int true "Dentist ID"
// @Success 200 {object} ent.Dentist
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentists/{id} [get]
func (ctl *DentistController) GetDentist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	d, err := ctl.client.Dentist.
		Query().
		Where(dentist.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, d)
 }
 

// ListDentist handles request to get a list of dentist entities
// @Summary List dentist entities
// @Description list dentist entities
// @ID list-dentist
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Dentist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentists [get]
func (ctl *DentistController) ListDentist(c *gin.Context) {
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

	dentists, err := ctl.client.Dentist.
		Query().
		WithNurse().
		WithDegree().
		WithExpert().
		WithGender().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dentists)
}

// NewDentistController creates and registers handles for the dentist controller
func NewDentistController(router gin.IRouter, client *ent.Client) *DentistController {
	dc := &DentistController{
		client: client,
		router: router,
	}

	dc.register()

	return dc

}

// InitDentistController registers routes to the main engine
func (ctl *DentistController) register() {
	dentists := ctl.router.Group("/dentists")

	dentists.POST("", ctl.CreateDentist)
	dentists.GET("", ctl.ListDentist)
	//queues.DELETE(":id", ctl.DeleteQueue)
  //test
}
