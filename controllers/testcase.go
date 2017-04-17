package controllers

import (
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extensions"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type testCommandController struct {
	*clayControllers.BaseController
}

type testPatternController struct {
	*clayControllers.BaseController
}

type testCaseController struct {
	*clayControllers.BaseController
}

func newTestCommandController() extensions.Controller {
	controller := &testCommandController{
		BaseController: clayControllers.NewBaseController(
			"test_command",
			models.SharedTestCommandModel(),
			logics.UniqueTestCommandLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func newTestPatternController() extensions.Controller {
	controller := &testPatternController{
		BaseController: clayControllers.NewBaseController(
			"test_pattern",
			models.SharedTestPatternModel(),
			logics.UniqueTestPatternLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func newTestCaseController() extensions.Controller {
	controller := &testCaseController{
		BaseController: clayControllers.NewBaseController(
			"test_case",
			models.SharedTestCaseModel(),
			logics.UniqueTestCaseLogic(),
		),
	}
	controller.SetOutputter(controller)
	return controller
}

func (controller *testCommandController) RouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleURL := extensions.BuildResourceSingleURL(controller.ResourceName())
	resourceMultiURL := extensions.BuildResourceMultiURL(controller.ResourceName())

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

func (controller *testPatternController) RouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleURL := extensions.BuildResourceSingleURL(controller.ResourceName())
	resourceMultiURL := extensions.BuildResourceMultiURL(controller.ResourceName())

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

func (controller *testCaseController) RouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleURL := extensions.BuildResourceSingleURL(controller.ResourceName())
	resourceMultiURL := extensions.BuildResourceMultiURL(controller.ResourceName())

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
		extensions.MethodPatch: {
			resourceSingleURL: controller.Patch,
		},
	}
	return routeMap
}

func (controller *testCaseController) OutputPatch(c *gin.Context, code int, result interface{}) {
	text := result.(string)
	c.String(code, text)
}

var uniqueTestCommandController = newTestCommandController()
var uniqueTestPatternController = newTestPatternController()
var uniqueTestCaseController = newTestCaseController()

func init() {
	extensions.RegisterController(uniqueTestCommandController)
	extensions.RegisterController(uniqueTestPatternController)
	extensions.RegisterController(uniqueTestCaseController)
}
