package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/aldhonie/go-simple-clean-crud/car/delivery/http/middleware"
	"github.com/aldhonie/go-simple-clean-crud/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CarHandler  represent the httphandler for car
type CarHandler struct {
	CarUsecase domain.CarUsecase
}

// NewCarHandler will initialize the cars/ resources endpoint
func NewCarHandler(e *echo.Echo, cu domain.CarUsecase) {
	handler := &CarHandler{
		CarUsecase: cu,
	}
	middleware := middleware.InitMiddleware()

	e.GET("/car/search", handler.SearchCar)
	e.GET("/cars", handler.FetchCars)
	e.GET("/car/:id", handler.GetByID)
	e.POST("/car", handler.Store, middleware.AuthHeader)
	e.POST("/car/edit/:id", handler.Update, middleware.AuthHeader)
	e.DELETE("/car/:id", handler.Delete, middleware.AuthHeader)
}

// SearchCar will search car by car name
func (a *CarHandler) SearchCar(c echo.Context) error {

	resListCar, err := a.CarUsecase.FetchByKeyword(c.Request().Context(), c.QueryParam("q"))

	if err != nil {
		return c.JSON(http.StatusOK, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resListCar)
}

// FetchCars will fetch the car based on given params
func (a *CarHandler) FetchCars(c echo.Context) error {

	listCar, _, err := a.CarUsecase.Fetch(c.Request().Context())
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listCar)
}

// GetByID will get car by given id
func (a *CarHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	id := int64(idP)
	ctx := c.Request().Context()

	car, err := a.CarUsecase.GetByID(ctx, id)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, car)
}

func isRequestValid(dc *domain.Car) (bool, error) {
	validate := validator.New()
	err := validate.Struct(dc)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the car by given request body
func (a *CarHandler) Store(c echo.Context) (err error) {
	var car domain.Car
	err = c.Bind(&car)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&car); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.CarUsecase.Store(ctx, &car)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, car)
}

// Update will delete car by given id
func (a *CarHandler) Update(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	car, err := a.CarUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	fmt.Println(car)

	var carEntity domain.Car

	err = c.Bind(&carEntity)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&carEntity); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = a.CarUsecase.Update(ctx, &carEntity)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, id)
}

// Delete will delete car by id
func (a *CarHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = a.CarUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseError{Message: fmt.Sprintf("id %d is deleted.", id)})
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
