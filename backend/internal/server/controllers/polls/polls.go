package polls

import (
	"net/http"
	"strconv"

	models "github.com/IainMcl/HereWeGo/internal/models/poll"
	"github.com/IainMcl/HereWeGo/internal/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var db *sqlx.DB

func Setup(d *sqlx.DB, a, r *echo.Group) (*echo.Group, *echo.Group) {
	db = d
	// Add restricted routes under /polls
	pollsRes := r.Group("/polls")
	// Add anonymous routes under /polls
	pollsAnon := a.Group("/polls")
	// pollsAnon.GET("", GetPolls)
	pollsAnon.GET("/:id", GetPoll)
	pollsRes.POST("", CreatePoll)
	pollsRes.PUT("/:id", UpdatePoll)
	// pollsRes.DELETE("/:id", DeletePoll)
	pollsRes.POST("/:id/options", CreatePollOption)
	// pollsRes.DELETE("/:id/options/:option_id", DeletePollOption)
	pollsRes.POST("/:id/options/:option_id/vote", VoteFor)
	// Remove vote
	pollsRes.DELETE("/:id/options/:option_id/vote", RemoveVote)

	// pollsRes.DELETE("/:id/options/:option_id/vote", VoteAgainst)
	return pollsAnon, pollsRes
}

type NewPoll struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NewPollOption struct {
	Name string `json:"name"`
}

// GetPoll godoc
//
//	@Summary		Get a poll
//	@Description	Get a poll
//	@Tags			polls
//	@Produce		json
//	@Param			id			path		int	true	"Poll ID"
//	@Success		200			{object}	models.NamedPoll
//	@Router			/polls/{id}	[get]
//
//	@Security		ApiKeyAuth
func GetPoll(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	poll, err := models.GetPollById(db, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, poll)
}

// CreatePoll godoc
//
//	@Summary		Create a poll
//	@Description	Create a poll
//	@Tags			polls
//	@Accept			json
//	@Param			body	body	NewPoll	true	"Poll"
//	@Produce		json
//	@Success		201		{object}	models.NamedPoll
//	@Router			/polls	[post]
//
//	@Security		ApiKeyAuth
func CreatePoll(c echo.Context) error {
	var np NewPoll
	if err := c.Bind(&np); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	userId, err := util.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	p := &models.NamedPoll{
		Title:       np.Title,
		Description: np.Description,
		CreatedById: userId,
	}
	if err := p.CreatePoll(db); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, p)
}

// UpdatePoll godoc
//
//	@Summary		Update a poll
//	@Description	Update a poll
//	@Tags			polls
//	@Accept			json
//	@Param			id		path	int		true	"Poll ID"
//	@Param			body	body	NewPoll	true	"Poll"
//	@Produce		json
//	@Success		200			{object}	models.NamedPoll
//	@Router			/polls/{id}	[put]
//
//	@Security		ApiKeyAuth
func UpdatePoll(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	var np NewPoll
	if err := c.Bind(&np); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	p := &models.NamedPoll{
		ID:          id,
		Title:       np.Title,
		Description: np.Description,
	}
	if err := p.UpdatePoll(db); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusOK, p)
}

// Create poll option godoc
//
//	@Summary		Create a poll option
//	@Description	Create a poll option
//	@Tags			polls
//	@Accept			json
//	@Param			id		path	int				true	"Poll ID"
//	@Param			body	body	NewPollOption	true	"Poll Option"
//	@Produce		json
//	@Success		201					{object}	models.NamedPollOption
//	@Router			/polls/{id}/options	[post]
//
//	@Security		ApiKeyAuth
func CreatePollOption(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	var np NewPollOption
	if err := c.Bind(&np); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	var po models.NamedPollOption = models.NamedPollOption{
		Name:   np.Name,
		PollId: id,
	}
	if err := po.CreatePollOption(db); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, po)
}

// VoteFor godoc
//
//	@Summary		Vote for a poll option
//	@Description	Vote for a poll option
//	@Tags			polls
//	@Param			option_id	path	int	true	"Poll Option ID"
//	@Produce		json
//	@Success		200										{string}	string
//	@Router			/polls/{id}/options/{option_id}/vote	[post]
//
//	@Security		ApiKeyAuth
func VoteFor(c echo.Context) error {
	idStr := c.Param("option_id")
	optionId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	userId, err := util.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	err = models.VoteFor(db, userId, optionId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Voted"})
}

// RemoveVote godoc
//
//	@Summary		Remove vote for a poll option
//	@Description	Remove vote for a poll option
//	@Tags			polls
//	@Param			option_id	path	int	true	"Poll Option ID"
//	@Produce		json
//	@Success		200										{string}	string
//	@Router			/polls/{id}/options/{option_id}/vote	[delete]
//
//	@Security		ApiKeyAuth
func RemoveVote(c echo.Context) error {
	idStr := c.Param("option_id")
	optionId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	userId, err := util.GetUserId(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	err = models.RemoveVote(db, userId, optionId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Vote removed"})
}
