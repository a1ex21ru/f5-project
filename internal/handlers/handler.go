package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"f5-project/internal/models"
	"f5-project/internal/service"
)

type NoteHandler struct {
	service *service.NoteService
}

func NewHandler(service *service.NoteService) *NoteHandler {
	return &NoteHandler{service: service}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	const fn = "CreateNote"

	var note models.Note

	err := c.BindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		log.Printf("%s - Error binding note: %v", fn, err)
		return
	}

	if err = h.service.Create(&note); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "error creating note",
			})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "note created",
	})
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	const fn = "UpdateNote"

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var input models.Note
	if err = c.BindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid input",
		})
		return
	}

	if err = h.service.Update(id, &input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error updating note",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "note updated",
	})
	return
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting note",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "note deleted",
	})
	return
}

func (h *NoteHandler) GetNote(c *gin.Context) {

	id := c.Param("id")

	note, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting note",
		})
		return
	}

	c.JSON(http.StatusOK, note)
	return
}

func (h *NoteHandler) GetAll(c *gin.Context) {
	notes, err := h.service.GetAll()
	if err != nil {
		log.Printf("Error getting all notes: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": notes,
	})
	return
}

func (h *NoteHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}
