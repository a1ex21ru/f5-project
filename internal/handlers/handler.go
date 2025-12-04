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

// CreateNote godoc
// @Summary      Create a new note
// @Description  Create a new note with the provided data
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        note  body      models.Note  true  "Note data"
// @Success      201   {object}  models.NoteResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /api/notes [post]
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
		c.JSON(http.StatusInternalServerError,
			models.ErrorResponse{
				Message: "error creating note",
			})
		return
	}

	c.JSON(http.StatusCreated, models.NoteResponse{
		Message: "note created",
	})
}

// UpdateNote godoc
// @Summary      Update a note
// @Description  Update an existing note by ID
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id    path      int                 true  "Note ID"
// @Param        note  body      models.UpdateNote  true  "Updated note data"
// @Success      200   {object}  models.NoteResponse
// @Failure      400   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /api/notes/{id} [patch]
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	const fn = "UpdateNote"

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "invalid id",
		})
		return
	}

	var input models.Note
	if err = c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "invalid input",
		})
		return
	}

	if err = h.service.Update(id, &input); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error updating note",
		})
	}

	c.JSON(http.StatusOK, models.NoteResponse{
		Message: "note updated",
	})
	return
}

// DeleteNote godoc
// @Summary      Delete a note
// @Description  Delete a note by ID
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Note ID"
// @Success      200 {object}  models.NoteResponse
// @Failure      400 {object}  models.ErrorResponse
// @Router       /api/notes/{id} [delete]
func (h *NoteHandler) DeleteNote(c *gin.Context) {

	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error deleting note",
		})
		return
	}

	c.JSON(http.StatusOK, models.NoteResponse{
		Message: "note deleted",
	})
	return
}

// GetNoteByID godoc
// @Summary      Get note by ID
// @Description  Get details of a specific note by its ID
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Note ID"
// @Success      200  {object}  models.Note
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/notes/{id} [get]
func (h *NoteHandler) GetNoteByID(c *gin.Context) {

	id := c.Param("id")

	note, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error getting note",
		})
		return
	}

	c.JSON(http.StatusOK, models.NoteResponse{
		Message: "note retrieved",
		Data:    *note,
	})
	return
}

// GetAll GetAllNotes godoc
// @Summary      Get all notes
// @Description  Retrieve a list of all notes
// @Tags         notes
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Notes
// @Failure      500  {object}  models.ErrorResponse
// @Router       /api/notes [get]
func (h *NoteHandler) GetAll(c *gin.Context) {
	notes, err := h.service.GetAll()
	if err != nil {
		log.Printf("Error getting all notes: %v", err)
		return
	}

	c.JSON(http.StatusOK, models.Notes{
		Notes: notes,
	})
	return
}

func (h *NoteHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.NoteResponse{
		Message: "OK",
	})
	return
}

func (h *NoteHandler) GetIndexPage(c *gin.Context) {
	c.File("./static/index.html")
}

func (h *NoteHandler) GetEditPage(c *gin.Context) {
	c.File("./static/edit.html")
}

func (h *NoteHandler) GetCreatePage(c *gin.Context) {
	c.File("./static/create.html")
}
