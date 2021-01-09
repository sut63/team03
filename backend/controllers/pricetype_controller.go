package controllers

import (
	"context"
	"strconv"
    "github.com/team03/app/ent/pricetype"
    "github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
)
// PriceTypeController defines the struct for the pricetype controller
type PriceTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// PriceTypeCreate handles POST requests for adding pricetype entities
// @Summary Create pricetype
// @Description Create pricetype
// @ID create-pricetype
// @Accept   json
// @Produce  json
// @Param pricetype body ent.PriceType true "PriceType entity"
// @Success 200 {object} ent.PriceType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pricetypes [post]
func (ctl *PriceTypeController) PriceTypeCreate(c *gin.Context) {
	obj := ent.PriceType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "pricetype binding failed",
		})
		return
	}

	pt, err := ctl.client.PriceType.
		Create().
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pt)
}

// GetPriceType handles GET requests to retrieve a pricetype entity
// @Summary Get a pricetype entity by ID
// @Description get pricetype by ID
// @ID get-pricetype
// @Produce  json
// @Param id path int true "PriceType ID"
// @Success 200 {object} ent.PriceType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pricetypes/{id} [get]
func (ctl *PriceTypeController) GetPriceType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	pt, err := ctl.client.PriceType.
		Query().
		Where(pricetype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pt)
}

// ListPriceType handles request to get a list of pricetype entities
// @Summary List pricetype entities
// @Description list pricetype entities
// @ID list-pricetype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PriceType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /pricetypes [get]
func (ctl *PriceTypeController) ListPriceType(c *gin.Context) {
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

	pricetypes, err := ctl.client.PriceType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, pricetypes)
}



// NewPriceTypeController creates and registers handles for the pricetype controller
func NewPriceTypeController(router gin.IRouter, client *ent.Client) *PriceTypeController {
	ptc := &PriceTypeController{
		client: client,
		router: router,
	}

	ptc.register()

	return ptc

}

func (ctl *PriceTypeController) register() {
	pricetypes := ctl.router.Group("/pricetypes")

	pricetypes.GET("", ctl.ListPriceType)

	// CRUD
	pricetypes.POST("", ctl.PriceTypeCreate)
	pricetypes.GET(":id", ctl.GetPriceType)
	
}
