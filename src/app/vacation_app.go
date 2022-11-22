package app

import (
	"net/http"
	domain "shake-shake/src/domain/vacation"
	"shake-shake/src/service"

	"github.com/gin-gonic/gin"
)

type VacationApp struct {
	vacationSvc *service.VacationService
}

func CreateVacationApp(router *gin.Engine) (*VacationApp, error) {
	vacationSvc, err := service.CreateVacationService()

	if err != nil {
		return nil, err
	}

	app := &VacationApp{vacationSvc: vacationSvc}

	router.POST("/vacation", app.createVacation)

	router.GET("/vacations", app.getVacations)

	router.DELETE("/vacations/:vacationId", app.deleteVacation)

	return app, nil

}

// @Summary Create vacation.
// @Description Create new vacation w/ memberId.
// @Success 200
// @Failure 400
// @Failure 500
// @Router /vacation [post]
func (app *VacationApp) createVacation(c *gin.Context) {
	body := new(domain.Vacation)
	err := c.Bind(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = app.vacationSvc.Create(c, body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// @Summary Get vacations.
// @Description Get vacations array.
// @Success 200 {array} domain.Vacation
// @Failure 500
// @Router /vacations [get]
func (app *VacationApp) getVacations(c *gin.Context) {
	vacations, err := app.vacationSvc.ReadMany(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, vacations)
}

// @Summary Delete vacation.
// @Description Delete vacation w/ vacationId.
// @Param vacationId path string true "Vacation id to delete."
// @Success 200
// @Failure 500
// @Router /vacations/{vacationId} [delete]
func (app *VacationApp) deleteVacation(c *gin.Context) {
	vacationId := c.Params.ByName("vacationId")
	err := app.vacationSvc.Delete(c, vacationId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
