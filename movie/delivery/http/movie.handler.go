package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/minhthong582000/go-movies-api/domain"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type MovieHandler struct {
	MovUsecase domain.MovieUsecase
}

func NewMovieHandler(r *gin.Engine, us domain.MovieUsecase) {
	handler := &MovieHandler{
		MovUsecase: us,
	}

	moviesRouter := r.Group("/movies")
	{
		moviesRouter.GET("", handler.FetchMovie)
		moviesRouter.GET("/:id", handler.GetByID)
		moviesRouter.POST("", handler.Store)
		moviesRouter.PUT("/:id", handler.Update)
		moviesRouter.DELETE("", handler.Delete)
	}
}

func (m *MovieHandler) FetchMovie(c *gin.Context) {
	listMovies, err := m.MovUsecase.Fetch()
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, listMovies)
}

func (m *MovieHandler) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
		return
	}
	id := int64(idP)

	mov, err := m.MovUsecase.GetByID(id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, mov)
}

func isRequestValid(m *domain.Movie) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *MovieHandler) Store(c *gin.Context) {
	var mov domain.Movie

	err := c.ShouldBindJSON(&mov)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Validate request obj
	if ok, err := isRequestValid(&mov); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = m.MovUsecase.Store(&mov)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mov)
}

func (m *MovieHandler) Update(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
		return
	}
	id := int64(idP)

	var mov domain.Movie

	err = c.Bind(&mov)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Validate request obj
	if ok, err := isRequestValid(&mov); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedRow, err := m.MovUsecase.Update(id, &mov)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"updated": updatedRow})
}

func (m *MovieHandler) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
		return
	}

	id := int64(idP)

	err = m.MovUsecase.Delete(id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
