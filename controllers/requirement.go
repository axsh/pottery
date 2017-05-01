package controllers

import (
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type protocolController struct {
	*clayControllers.BaseController
}

type serviceController struct {
	*clayControllers.BaseController
}

type connectionController struct {
	*clayControllers.BaseController
}

type requirementController struct {
	*clayControllers.BaseController
}

func newProtocolController() extensions.Controller {
	controller := &protocolController{
		BaseController: clayControllers.NewBaseController(
			models.SharedProtocolModel(),
			logics.UniqueProtocolLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func newServiceController() extensions.Controller {
	controller := &serviceController{
		BaseController: clayControllers.NewBaseController(
			models.SharedServiceModel(),
			logics.UniqueServiceLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func newConnectionController() extensions.Controller {
	controller := &connectionController{
		BaseController: clayControllers.NewBaseController(
			models.SharedConnectionModel(),
			logics.UniqueConnectionLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func newRequirementController() extensions.Controller {
	controller := &requirementController{
		BaseController: clayControllers.NewBaseController(
			models.SharedRequirementModel(),
			logics.UniqueRequirementLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *protocolController) RouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleURL := controller.ResourceSingleURL()
	resourceMultiURL := controller.ResourceMultiURL()

	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			resourceSingleURL: controller.GetSingle,
			resourceMultiURL:  controller.GetMulti,
		},
		extensions.MethodPost: {
			resourceMultiURL: controller.Create,
		},
		extensions.MethodPut: {
			resourceSingleURL: controller.Update,
		},
		extensions.MethodDelete: {
			resourceSingleURL: controller.Delete,
		},
	}
	return routeMap
}

func (controller *serviceController) RouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleURL := controller.ResourceSingleURL()
	resourceMultiURL := controller.ResourceMultiURL()

	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			resourceSingleURL: controller.GetSingle,
			resourceMultiURL:  controller.GetMulti,
		},
		extensions.MethodPost: {
			resourceMultiURL: controller.Create,
		},
		extensions.MethodPut: {
			resourceSingleURL: controller.Update,
		},
		extensions.MethodDelete: {
			resourceSingleURL: controller.Delete,
		},
	}
	return routeMap
}

func (controller *connectionController) RouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleURL := controller.ResourceSingleURL()
	resourceMultiURL := controller.ResourceMultiURL()

	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			resourceSingleURL: controller.GetSingle,
			resourceMultiURL:  controller.GetMulti,
		},
		extensions.MethodPost: {
			resourceMultiURL: controller.Create,
		},
		extensions.MethodPut: {
			resourceSingleURL: controller.Update,
		},
		extensions.MethodDelete: {
			resourceSingleURL: controller.Delete,
		},
	}
	return routeMap
}

func (controller *requirementController) RouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleURL := controller.ResourceSingleURL()
	resourceMultiURL := controller.ResourceMultiURL()

	routeMap := map[int]map[string]gin.HandlerFunc{
		extensions.MethodGet: {
			resourceSingleURL: controller.GetSingle,
			resourceMultiURL:  controller.GetMulti,
		},
		extensions.MethodPost: {
			resourceMultiURL: controller.Create,
		},
		extensions.MethodPut: {
			resourceSingleURL: controller.Update,
		},
		extensions.MethodDelete: {
			resourceSingleURL: controller.Delete,
		},
	}
	return routeMap
}

var uniqueProtocolController = newProtocolController()
var uniqueServiceController = newServiceController()
var uniqueConnectionController = newConnectionController()
var uniqueRequirementController = newRequirementController()

func init() {
	extensions.RegisterController(uniqueProtocolController)
	extensions.RegisterController(uniqueServiceController)
	extensions.RegisterController(uniqueConnectionController)
	extensions.RegisterController(uniqueRequirementController)
}
