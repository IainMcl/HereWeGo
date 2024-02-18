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
	// pollsRes.PUT("/:id", UpdatePoll)
	// pollsRes.DELETE("/:id", DeletePoll)
	// pollsRes.POST("/:id/options", CreatePollOption)
	// pollsRes.DELETE("/:id/options/:option_id", DeletePollOption)
	// pollsRes.POST("/:id/options/:option_id/vote", VoteFor)
	// pollsRes.DELETE("/:id/options/:option_id/vote", VoteAgainst)
	return pollsAnon, pollsRes
}

type NewPoll struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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
//	@Param			body	body	models.NamedPoll	true	"Poll"
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
	user, err := util.ClaimsFromToken(c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	p := &models.NamedPoll{
		Title:       np.Title,
		Description: np.Description,
		CreatedById: user.UserId,
	}
	if err := p.CreatePoll(db); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	return c.JSON(http.StatusCreated, p)
}
