package handler

import (
	"musicLibrary/internal/models"
	"musicLibrary/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(s service.Service) *Handler {
    return &Handler{service: s}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/songs", h.CreateSong)
		api.PUT("/songs/:id", h.UpdateSong)
		api.DELETE("/songs/:id", h.DeleteSong)
		api.GET("/songs/:id", h.GetSong)
		api.GET("/songs", h.ListSongs)		
	}
}

func (h *Handler) CreateSong(c *gin.Context) {
	var song models.Song

	if err := c.BindJSON(&song); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	if err := h.service.CreateSong(&song); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

	c.JSON(201, song)
}

// 	   ID          int64     `json:"id"`
//     Group       string    `json:"group"`
//     SongName    string    `json:"song"`
//     ReleaseDate string    `json:"releaseDate"`
//     Text        string    `json:"text"`
//     Link        string    `json:"link"`

func (h *Handler) UpdateSong(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// group := c.Param("group")
	// songName := c.Param("song")
	// releaseDate := c.Param("releaseDate")
	// text := c.Param("text")
	// link := c.Param("link")
	
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	var song models.Song
	if err := c.BindJSON(&song); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	song.ID = id
	// song.Group = group
	// song.SongName = songName
	// song.ReleaseDate = releaseDate
	// song.Text = text
	// song.Link = link

	if err := h.service.UpdateSong(&song); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
        return
	}

	c.JSON(201, song)
}

func (h *Handler) DeleteSong(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.DeleteSongById(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}

func (h *Handler) GetSong(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    song, err := h.service.GetSongById(id)

    if err!= nil {
        c.JSON(404, gin.H{"error": "Song not found"})
        return
    }

    c.JSON(200, song)
}

func (h *Handler) ListSongs(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	songs, err := h.service.ListSongs(offset, limit)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, songs)
}