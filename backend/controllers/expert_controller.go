package controllers

import (
	"context"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/expert"
)

//ExpertController defines the struct for the Expert controller

type ExpertController struct {
	client *ent.Client
	router gin.IRouter
}

// ExpertCreatee handles POST requests for adding  expert entities
// @Summary Create expert
// @Description Create expert
// @ID create-expert
// @Accept   json
// @Produce  json
// @Param expert  body ent.Expert true " Expert entity"
// @Success 200 {object} ent.Expert
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /experts [post]
func (ctl *ExpertController) ExpertCreate(c *gin.Context) {
	obj := ent.Expert{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "expert binding failed",
		})
		return
	}

	e, err := ctl.client.Expert.
		Create().
		SetName(obj.Name).
		
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, e)
}

// GetExpert handles GET requests to retrieve a expert entity
// @Summary Get a expert entity by ID
// @Description get expert by ID
// @ID get-expert
// @Produce  json
// @Param id path int true "Expert ID"
// @Success 200 {object} ent.Expert
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /experts/{id} [get]
func (ctl *ExpertController) GetExpert(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	e, err := ctl.client.Expert.
		Query().
		Where(expert.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, e)
}

// ListExpert handles request to get a list of expert entities
// @Summary List expert entities
// @Description list expert entities
// @ID list-expert
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Expert
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /experts [get]
func (ctl *ExpertController) ListExpert(c *gin.Context) {
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

	experts, err := ctl.client.Expert.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, experts)
}

// NewExpertController creates and registers handles for the expert controller
func NewExpertController(router gin.IRouter, client *ent.Client) *ExpertController {
	ec := &ExpertController{
		client: client,
		router: router,
	}
	ec.register()
	return ec
}

// InitExpertController registers routes to the main engine
func (ctl *ExpertController) register() {
	experts := ctl.router.Group("/experts")

	experts.GET("", ctl.ListExpert)
	// CRUD
	experts.POST("", ctl.ExpertCreate)
	experts.GET(":id", ctl.GetExpert)

}
