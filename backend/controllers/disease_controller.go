package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/disease"
)

// DiseaseController defines the struct for the disease controller
type DiseaseController struct {
	client *ent.Client
	router gin.IRouter
}

// DiseaseCreate handles POST requests for adding disease entities
// @Summary Create disease
// @Description Create disease
// @ID create-disease
// @Accept   json
// @Produce  json
// @Param disease body ent.Disease true "Disease entity"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [post]

func (ctl *DiseaseController) DiseaseCreate(c *gin.Context) {
	obj := ent.Disease{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Disease binding failed",
		})
		return
	}

	ds, err := ctl.client.Disease.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, ds)
}

// GetDisease handles GET requests to retrieve a disease entity
// @Summary Get a disease entity by ID
// @Description get disease by ID
// @ID get-disease
// @Produce  json
// @Param id path int true "Disease ID"
// @Success 200 {object} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases/{id} [get]
func (ctl *DiseaseController) GetDisease(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	d, err := ctl.client.Disease.
		Query().
		Where(disease.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, ds)
}

// ListDisease handles request to get a list of disease entities
// @Summary List disease entities
// @Description list disease entities
// @ID list-disease
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Disease
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /diseases [get]
func (ctl *DiseaseController) ListDisease(c *gin.Context) {
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

	diseases, err := ctl.client.Disease.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, diseases)
}

// NewDiseaseController creates and registers handles for the disease controller
func NewDiseaseController(router gin.IRouter, client *ent.Client) *DiseaseController {
	dsc := &DiseaseController{
		client: client,
		router: router,
	}
	dsc.register()
	return dsc
}

// InitDiseaseController registers routes to the main engine
func (ctl *DiseaseController) register() {
	diseases := ctl.router.Group("/diseases")

	diseases.POST("", ctl.DiseaseCreate)
	diseases.GET(":id", ctl.GetDisease)
	diseases.GET("", ctl.ListDisease)
}
