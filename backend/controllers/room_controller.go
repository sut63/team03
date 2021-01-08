package controllers

import (
	"context"
	"strconv"
	"github.com/team03/app/ent"
	"github.com/team03/app/ent/room"
	"github.com/gin-gonic/gin"
 )

//RoomController defines the struct for the Room controller
type RoomController struct {
	client *ent.Client
	router gin.IRouter
}

// RoomCreate handles POST requests for adding Room entities
// @Summary Create Room
// @Description Create Room
// @ID create-Room
// @Accept   json
// @Produce  json
// @Param Room body ent.Room true "Room entity"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms [post]
func (ctl *RoomController) RoomCreate(c *gin.Context) {
	obj := ent.Room{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	r, err := ctl.client.Room.
		Create().
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// GetRoom handles GET requests to retrieve a Room entity
// @Summary Get a Room entity by ID
// @Description get Room by ID
// @ID get-Room
// @Produce  json
// @Param id path int true "Room ID"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms/{id} [get]
func (ctl *RoomController) GetRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListRoom handles request to get a list of Room entities
// @Summary List Room entities
// @Description list Room entities
// @ID list-Room
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms [get]
func (ctl *RoomController) ListRoom(c *gin.Context) {
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

	rooms, err := ctl.client.Room.
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

	c.JSON(200, rooms)
}


// NewRoomController creates and registers handles for the Room controller
func NewRoomController(router gin.IRouter, client *ent.Client) *RoomController {
	rc := &RoomController{
		client: client,
		router: router,
	}

	rc.register()
	return rc

}

func (ctl *RoomController) register() {
	rooms := ctl.router.Group("/rooms")
	rooms.GET("", ctl.ListRoom)

	// CRUD
	rooms.POST("", ctl.RoomCreate)
	rooms.GET(":id", ctl.GetRoom)
}
