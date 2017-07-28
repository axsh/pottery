package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type environmentCommitmentController struct {
	*clayControllers.BaseController
}

func newEnvironmentCommitmentController() extensions.Controller {
	controller := &environmentCommitmentController{
		BaseController: clayControllers.NewBaseController(
			models.SharedEnvironmentModel(),
			logics.UniqueEnvironmentCommitmentLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *environmentCommitmentController) RouteMap() map[int]map[string]gin.HandlerFunc {
	url := fmt.Sprintf("%s/present/commitment", controller.ResourceName())

	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodPut: {
			url: controller.Update,
		},
	}
	return routeMap
}

func (controller *environmentCommitmentController) OutputGetSingle(c *gin.Context, code int, result interface{}, fields map[string]interface{}) {
	text := result.(string)
	c.String(code, text)
}

var uniqueEnvironmentCommitmentController = newEnvironmentCommitmentController()

func init() {
	extensions.RegisterController(uniqueEnvironmentCommitmentController)
}
