package routes

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"f5-project/internal/handlers"
)

type Handler struct {
	*handlers.NoteHandler
}

func NewHandler(handler *handlers.NoteHandler) *Handler {
	return &Handler{handler}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(static.Serve("/static", static.LocalFile("./static", true)))
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	note := router.Group("/api/notes")
	{
		note.POST("/", h.CreateNote)
		note.DELETE("/:id", h.DeleteNote)
		note.PATCH("/:id", h.UpdateNote)
		note.GET("/:id", h.GetNote)
		note.GET("", h.GetAll)
	}

	healthCheck := router.Group("/healthcheck")
	{
		healthCheck.GET("/", h.HealthCheck)
	}

	return router
}
