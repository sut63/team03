package controllers

import (
	"context"
	"strconv"
	"time"
    "github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/medicalfile"
	"github.com/team03/app/ent/pricetype"
	"github.com/team03/app/ent/nurse"
	"github.com/team03/app/ent/dentalexpense"
)

type DentalExpenseController struct {
	client *ent.Client
	router gin.IRouter
}

type DentalExpense struct {
	Medicalfile   int
	PriceType     int
	Nurse         int
	Added         string
}

// DentalExpenseCreate handles POST requests for adding dentalexpense entities
// @Summary Create dentalexpense
// @Description Create dentalexpense
// @ID create-dentalexpense
// @Accept   json
// @Produce  json
// @Param dentalexpense body ent.DentalExpense true "DentalExpense entity"
// @Success 200 {object} ent.DentalExpense
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalexpenses [post]
func (ctl *DentalExpenseController) DentalExpenseCreate(c *gin.Context) {
	obj := DentalExpense{}
	if err := c.ShouldBind(&obj); 
	err != nil {
		c.JSON(400, gin.H{
			"error": "dentalexpense binding failed",
		})
		return
	}

	mf, err := ctl.client.Medicalfile.
		Query().
		Where(medicalfile.IDEQ(int(obj.Medicalfile))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "medicalfile not found",
		})
		return
	}

	pt, err := ctl.client.PriceType.
		Query().
		Where(pricetype.IDEQ(int(obj.PriceType))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "pricetype not found",
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

	time, err := time.Parse(time.RFC3339, obj.Added)
	de, err := ctl.client.DentalExpense.
		Create().
		SetAddedTime(time).
		SetMedicalfile(mf).
		SetPricetype(pt).
		SetNurse(n).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   de,
	})
}
// GetDentalExpense handles GET requests to retrieve a dentalexpense entity
// @Summary Get a dentalexpense entity by ID
// @Description get dentalexpense by ID
// @ID get-dentalexpense
// @Produce  json
// @Param id path int true "DentalExpense ID"
// @Success 200 {object} ent.DentalExpense
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalexpenses/{id} [get]
func (ctl *DentalExpenseController) GetDentalExpense(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	de, err := ctl.client.DentalExpense.
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
// ListDentalExpense handles request to get a list of dentalexpense entities
// @Summary List dentalexpense entities
// @Description list dentalexpense entities
// @ID list-dentalexpense
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DentalExpense
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentalexpenses [get]
func (ctl *DentalExpenseController) ListDentalExpense(c *gin.Context) {
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

	dentalExpenses, err := ctl.client.DentalExpense.
		Query().
		WithMedicalfile().
		WithPricetype().
		WithNurse().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dentalExpenses)
}

// NewDentalExpenseController creates and registers handles for the dentalexpense controller
func NewDentalExpenseController(router gin.IRouter, client *ent.Client) *DentalExpenseController {
	dec := &DentalExpenseController{
		client: client,
		router: router,
	}

	dec.register()

	return dec

}

func (ctl *DentalExpenseController) register() {
	dentalExpenses := ctl.router.Group("/dentalexpenses")
    dentalExpenses.GET(":id", ctl.GetDentalExpense)
	dentalExpenses.POST("", ctl.DentalExpenseCreate)
	dentalExpenses.GET("", ctl.ListDentalExpense)

}
