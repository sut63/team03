package controllers

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/team03/app/ent"
	"github.com/team03/app/ent/dentist"
	"github.com/team03/app/ent/patient"
	"github.com/team03/app/ent/queue"
	"github.com/gin-gonic/gin"
)

type QueueController struct {
	client *ent.Client
	router gin.IRouter
}

type Queue struct {
	QueueID   string
	Patient   int
	Phone     string
	Dentist   int
	Dental    string
	QueueTime string
}

// QueueCreate handles POST requests for adding queue entities
// @Summary Create queue
// @Description Create queue
// @ID create-queue
// @Accept   json
// @Produce  json
// @Param queue body Queue true "Queue entity"
// @Success 200 {object} Queue
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /queues [post]
func (ctl *QueueController) QueueCreate(c *gin.Context) {
	obj := Queue{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "queue binding failed",
		})
		return
	}

	d, err := ctl.client.Dentist.
		Query().
		Where(dentist.IDEQ(int(obj.Dentist))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "dentist not found",
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

	time, err := time.Parse(time.RFC3339, obj.QueueTime)

	q, err := ctl.client.Queue.
		Create().
		SetQueueID(obj.QueueID).
		SetPatient(p).
		SetPhone(obj.Phone).
		SetDentist(d).
		SetDental(obj.Dental).
		SetQueueTime(time).
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
		"data":   q,
	})
}

// GetQueue handles GET requests to retrieve a queue entity
// @Summary Get a queue entity by ID
// @Description get queue by ID
// @ID get-queue
// @Produce  json
// @Param id path int true "Queue ID"
// @Success 200 {object} ent.Queue
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /queues/{id} [get]
func (ctl *QueueController) GetQueue(c *gin.Context) { 
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	q, err := ctl.client.Queue.
		Query().
		WithDentist().
		WithPatient().
		WithNurse().
		Where(queue.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, q)
}

// GetSearchQueue handles GET requests to retrieve a queue entity
// @Summary Get a Queue entity by Search
// @Description get queue by Search
// @ID get-Queue-by-search
// @Produce  json
// @Param Queue query string false "Search Queue"
// @Success 200 {object} ent.Queue
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchqueues [get]
func (ctl *QueueController) GetSearchQueue(c *gin.Context) {
	queuesearch := c.Query("queue")

	qs, err := ctl.client.Queue.
		Query().
		WithDentist().
		WithPatient().
		Where(queue.QueueIDContains(queuesearch)).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": qs,
	})
}

// ListQueue handles request to get a list of queue entities
// @Summary List queue entities
// @Description list queue entities
// @ID list-queue
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Queue
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /queues [get]
func (ctl *QueueController) ListQueue(c *gin.Context) {
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

	queues, err := ctl.client.Queue.
		Query().
		WithPatient().
		WithDentist().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, queues)
}

// NewQueueController creates and registers handles for the queue controller
func NewQueueController(router gin.IRouter, client *ent.Client) *QueueController {
	qc := &QueueController{
		client: client,
		router: router,
	}

	qc.register()

	return qc

}

func (ctl *QueueController) register() {
	queues := ctl.router.Group("/queues")

	queues.POST("", ctl.QueueCreate)
	queues.GET("", ctl.ListQueue)
	queues.GET(":id", ctl.GetQueue)

	searchqueues := ctl.router.Group("/searchqueues")
	searchqueues.GET("",ctl.GetSearchQueue)
	

}
