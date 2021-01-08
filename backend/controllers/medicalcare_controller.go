package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/medicalcare"
)

// MedicalCareController defines the struct for the medicalcare controller
type MedicalCareController struct {
	client *ent.Client
	router gin.IRouter
}

// MedicalCareCreate handles POST requests for adding medicalcare entities
// @Summary Create medicalcare
// @Description Create medicalcare
// @ID create-medicalcare
// @Accept   json
// @Produce  json
// @Param medicalcare body ent.MedicalCare true "MedicalCare entity"
// @Success 200 {object} ent.MedicalCare
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalcares [post]

func (ctl *MedicalCareController) MedicalCareCreate(c *gin.Context) {
	obj := ent.MedicalCare{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "MedicalCare binding failed",
		})
		return
	}

	mc, err := ctl.client.MedicalCare.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, mc)
}

// GetMedicalCare handles GET requests to retrieve a medicalcare entity
// @Summary Get a medicalcare entity by ID
// @Description get medicalcare by ID
// @ID get-medicalcare
// @Produce  json
// @Param id path int true "MedicalCare ID"
// @Success 200 {object} ent.MedicalCare
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalcares/{id} [get]
func (ctl *MedicalCareController) GetMedicalCare(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	mc, err := ctl.client.MedicalCare.
		Query().
		Where(medicalcare.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, mc)
}

// ListMedicalCare handles request to get a list of medicalcare entities
// @Summary List medicalcare entities
// @Description list medicalcare entities
// @ID list-medicalcare
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.MedicalCare
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalcares [get]
func (ctl *MedicalCareController) ListMedicalCare(c *gin.Context) {
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

	medicalcares, err := ctl.client.MedicalCare.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, medicalcares)
}

// NewMedicalCareController creates and registers handles for the medicalcare controller
func NewMedicalCareController(router gin.IRouter, client *ent.Client) *MedicalCareController {
	mcc := &MedicalCareController{
		client: client,
		router: router,
	}
	mcc.register()
	return mcc
}

// InitMedicalCareController registers routes to the main engine
func (ctl *MedicalCareController) register() {
	medicalcares := ctl.router.Group("/medicalcares")
	medicalcares.GET("", ctl.ListMedicalCare)
	// CRUD
	medicalcares.POST("", ctl.MedicalCareCreate)
	medicalcares.GET(":id", ctl.GetMedicalCare)
}
