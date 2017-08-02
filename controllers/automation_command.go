package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	dbpkg "github.com/qb0C80aE/clay/db"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type automationCommandController struct {
	*clayControllers.BaseController
}

func newAutomationCommandController() extensions.Controller {
	controller := &automationCommandController{
		BaseController: clayControllers.NewBaseController(
			models.SharedAutomationCommandModel(),
			logics.UniqueAutomationCommandLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *automationCommandController) RouteMap() map[int]map[string]gin.HandlerFunc {
	url := fmt.Sprintf("%s/execution", controller.ResourceName())
	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodPut: {
			url: controller.Update,
		},
	}
	return routeMap
}

func (controller *automationCommandController) Update(c *gin.Context) {
	dbpkg.Instance(c).Exec("pragma foreign_keys = off;")
	controller.BaseController.Update(c)
	dbpkg.Instance(c).Exec("pragma foreign_keys = on;")
}

var uniqueAutomationCommandController = newAutomationCommandController()

func init() {
	extensions.RegisterController(uniqueAutomationCommandController)
}
