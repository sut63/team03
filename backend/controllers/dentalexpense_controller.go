package controllers

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/dentalexpense"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/pricetype"
)

//DentalexpenseController defind the struct for Dentalexpense controller
type DentalexpenseController struct {
	client *ent.Client
	router gin.IRouter
}

//Dentalexpense defind the struct for the Dentalexpense controller
type Dentalexpense struct {
	Medicalfile 	int
	Pricetype  	    int
	Name   	        string
	AddedTime  	    string
	Tax		        string
	Amount          string
	Phone           string
	
}

// DentalexpenseCreate handles POST requests for adding  Dentalexpense entities
// @Summary Create Dentalexpense
// @Description Create Dentalexpense
// @ID create-Dentalexpense
// @Accept   json
// @Produce  json
// @Param  Dentalexpense body ent.Dentalexpense true "Dentalexpense entity"
// @Success 200 {object} ent.Dentalexpense
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalexpenses [post]
func (ctl *DentalexpenseController) DentalexpenseCreate(c *gin.Context) {
	obj := Dentalexpense{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Dentalexpense binding failed",
		})
		return
	}

	m, err := ctl.client.Medicalfile.
		Query().
		Where(medicalfile.IDEQ(int(obj.Medicalfile))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "medicalfile not found",
		})
		return
	}

	pt, err := ctl.client.Pricetype.
		Query().
		Where(pricetype.IDEQ(int(obj.Pricetype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "pricetype not found",
		})
		return
	}

	
	time, err := time.Parse(time.RFC3339, obj.AddedTime)
	am, err := strconv.Atoi(obj.Amount)
	de, err := ctl.client.Dentalexpense.
	    Create().
	    SetMedicalfile(m).
	    SetPricetype(pt).
	    SetName(obj.Name).
	    SetTax(obj.Tax).
	    SetPhone(obj.Phone).
	    SetAmount(am).
	    SetAddedTime(time).
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
		"data":   de,
	})
}

// GetDentalexpense handles GET requests to retrieve a Dentalexpense entity
// @Summary Get a Dentalexpense entity by ID
// @Description get Dentalexpense by ID
// @ID get-Dentalexpense
// @Produce  json
// @Param id path int true "Dentalexpense ID"
// @Success 200 {object} ent.Dentalexpense
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalexpenses/{id} [get]
func (ctl *DentalexpenseController) GetDentalexpense(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	de, err := ctl.client.Dentalexpense.
		Query().
		WithMedicalfile().
		WithPricetype().
		
		WithNurse().
		Where(dentalexpense.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, de)
}


// GetSearchDentalexpense handles GET requests to retrieve a Dentalexpense entity
// @Summary Get a Dentalexpense entity by Search
// @Description get Dentalexpense by Search
// @ID get-Dentalexpense-by-search
// @Produce  json
// @Param Dentalexpense query string false "Search Dentalexpense"
// @Success 200 {object} ent.Dentalexpense
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchdentalexpenses [get]
func (ctl *DentalexpenseController) GetSearchDentalexpense(c *gin.Context) {
	dentalexpensesearch := c.Query("dentalexpense")

	des, err := ctl.client.Dentalexpense.
		Query().
		WithMedicalfile().
		WithPricetype().
		Where(dentalexpense.TaxContains(dentalexpensesearch)).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": des,
	})
}

// ListDentalexpense handles request to get a list of Dentalexpense entities
// @Summary List Dentalexpense entities
// @Description list Dentalexpense entities
// @ID list-Dentalexpense
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Dentalexpense
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalexpenses [get]
func (ctl *DentalexpenseController) ListDentalexpense(c *gin.Context) {
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

	dentalexpenses, err := ctl.client.Dentalexpense.
		Query().
		WithMedicalfile().
		WithPricetype().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dentalexpenses)
}

// NewDentalexpenseController creates and registers handles for the dentalexpense controller
func NewDentalexpenseController(router gin.IRouter, client *ent.Client) *DentalexpenseController {
	dec := &DentalexpenseController{
		client: client,
		router: router,
	}

	dec.register()
	return dec

}

func (ctl *DentalexpenseController) register() {
	dentalexpenses := ctl.router.Group("/dentalexpenses")
	dentalexpenses.GET(":id", ctl.GetDentalexpense)

	searchdentalexpenses := ctl.router.Group("/searchdentalexpenses")
	searchdentalexpenses.GET("",ctl.GetSearchDentalexpense)

	// CRUD
	dentalexpenses.POST("", ctl.DentalexpenseCreate)
	dentalexpenses.GET("", ctl.ListDentalexpense)

}
