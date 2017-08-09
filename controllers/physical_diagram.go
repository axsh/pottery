package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type physicalDiagramController struct {
	*clayControllers.BaseController
}

type physicalDiagramNodeController struct {
	*clayControllers.BaseController
}

func newPhysicalDiagramController() extensions.Controller {
	controller := &physicalDiagramController{
		BaseController: clayControllers.NewBaseController(
			models.SharedDiagramModel(),
			logics.UniquePhysicalDiagramLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func newPhysicalDiagramNodeController() extensions.Controller {
	controller := &physicalDiagramNodeController{
		BaseController: clayControllers.NewBaseController(
			models.SharedDiagramNodeModel(),
			logics.UniquePhysicalDiagramNodeLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *physicalDiagramController) RouteMap() map[int]map[string]gin.HandlerFunc {
	url := fmt.Sprintf("%s/physical", controller.ResourceName())
	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			url: controller.GetSingle,
		},
	}
	return routeMap
}

func (controller *physicalDiagramNodeController) RouteMap() map[int]map[string]gin.HandlerFunc {
	singleUrl := fmt.Sprintf("%s/physical/%s/:id", uniquePhysicalDiagramController.ResourceName(), controller.ResourceName())
	multiUrl := fmt.Sprintf("%s/physical/%s", uniquePhysicalDiagramController.ResourceName(), controller.ResourceName())
	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			singleUrl: controller.GetSingle,
			multiUrl: controller.GetMulti,
		},
		extensions.MethodPut: {
			singleUrl: controller.Update,
		},
		extensions.MethodDelete: {
			singleUrl: controller.Delete,
		},
	}
	return routeMap
}

var uniquePhysicalDiagramController = newPhysicalDiagramController()
var uniquePhysicalDiagramNodeController = newPhysicalDiagramNodeController()

func init() {
	extensions.RegisterController(uniquePhysicalDiagramController)
	extensions.RegisterController(uniquePhysicalDiagramNodeController)
}
