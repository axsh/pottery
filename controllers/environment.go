package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type environmentController struct {
	*clayControllers.BaseController
}

func newEnvironmentController() extensions.Controller {
	controller := &environmentController{
		BaseController: clayControllers.NewBaseController(
			models.SharedEnvironmentModel(),
			logics.UniqueEnvironmentLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *environmentController) RouteMap() map[int]map[string]gin.HandlerFunc {
	url := fmt.Sprintf("%s/present", controller.ResourceName())
	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			url: controller.GetSingle,
		},
		extensions.MethodPut: {
			url: controller.Update,
		},
	}
	return routeMap
}

var uniqueEnvironmentController = newEnvironmentController()

func init() {
	extensions.RegisterController(uniqueEnvironmentController)
}
