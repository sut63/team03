package controllers

import (
	"context"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/degree"
)

//DegreeController defines the struct for the Degree controller

type DegreeController struct {
	client *ent.Client
	router gin.IRouter
}

// DegreeCreatee handles POST requests for adding  degree entities
// @Summary Create degree
// @Description Create degree
// @ID create-degree
// @Accept   json
// @Produce  json
// @Param degree  body ent.Degree true " Degree entity"
// @Success 200 {object} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees [post]
func (ctl *DegreeController) DegreeCreate(c *gin.Context) {
	obj := ent.Degree{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "degree binding failed",
		})
		return
	}

	dg, err := ctl.client.Degree.
		Create().
		SetName(obj.Name).
		
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, dg)
}

// GetDegree handles GET requests to retrieve a degree entity
// @Summary Get a Degree entity by ID
// @Description get degree by ID
// @ID get-degree
// @Produce  json
// @Param id path int true "Degree ID"
// @Success 200 {object} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees/{id} [get]
func (ctl *DegreeController) GetDegree(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	dg, err := ctl.client.Degree.
		Query().
		Where(degree.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dg)
}

// ListDegree handles request to get a list of degree entities
// @Summary List degree entities
// @Description list degree entities
// @ID list-degree
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Degree
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /degrees [get]
func (ctl *DegreeController) ListDegree(c *gin.Context) {
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

	degrees, err := ctl.client.Degree.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, degrees)
}

// NewDegreeController creates and registers handles for the degree controller
func NewDegreeController(router gin.IRouter, client *ent.Client) *DegreeController {
	dgc := &DegreeController{
		client: client,
		router: router,
	}
	ec.register()
	return dgc
}

// InitDegreeController registers routes to the main engine
func (ctl *DegreeController) register() {
	degrees := ctl.router.Group("/degrees")

	degrees.GET("", ctl.ListDegree)
	// CRUD
	degrees.POST("", ctl.DegreeCreate)
	degrees.GET(":id", ctl.GetDegree)

}
