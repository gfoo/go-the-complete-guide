package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"org.gfoo/api-rest/models"
	"org.gfoo/api-rest/utils"
)

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID param"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event from DB"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event from DB"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID param"})
		return
	}
	_, err = models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event from DB"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": "Could not parse request data (not fit the model)"})
		return
	}

	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event in DB"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func getEventsById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID param"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event from DB"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events from DB"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized,
			gin.H{"message": "Token is required"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"error": "Could not parse request data (not fit the model)"})
		return
	}

	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could save event in DB"})
		return
	}
	context.JSON(http.StatusCreated,
		gin.H{"event": event, "message": "Event created successfully!"},
	)
}
