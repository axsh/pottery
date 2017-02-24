package controllers

import (
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extension"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type TestCommandController struct {
	clayControllers.BaseController
}

type TestPatternController struct {
	clayControllers.BaseController
}

type TestCaseController struct {
	clayControllers.BaseController
}

func init() {
	extension.RegisterController(NewTestCommandController())
	extension.RegisterController(NewTestPatternController())
	extension.RegisterController(NewTestCaseController())
}

func NewTestCommandController() *TestCommandController {
	controller := &TestCommandController{}
	controller.Initialize()
	return controller
}

func NewTestPatternController() *TestPatternController {
	controller := &TestPatternController{}
	controller.Initialize()
	return controller
}

func NewTestCaseController() *TestCaseController {
	controller := &TestCaseController{}
	controller.Initialize()
	return controller
}

func (this *TestCommandController) Initialize() {
	this.ResourceName = "test_command"
	this.Model = models.TestCommandModel
	this.Logic = logics.TestCommandLogicInstance
	this.Outputter = this
}

func (this *TestPatternController) Initialize() {
	this.ResourceName = "test_pattern"
	this.Model = models.TestPatternModel
	this.Logic = logics.TestPatternLogicInstance
	this.Outputter = this
}

func (this *TestCaseController) Initialize() {
	this.ResourceName = "test_case"
	this.Model = models.TestCaseModel
	this.Logic = logics.TestCaseLogicInstance
	this.Outputter = this
}

func (this *TestCommandController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleUrl := extension.GetResourceSingleUrl(this.ResourceName)
	resourceMultiUrl := extension.GetResourceMultiUrl(this.ResourceName)

	routeMap := map[int]map[string]gin.HandlerFunc{
		extension.MethodGet: {
			resourceSingleUrl: this.GetSingle,
			resourceMultiUrl:  this.GetMulti,
		},
		extension.MethodPost: {
			resourceMultiUrl: this.Create,
		},
		extension.MethodPut: {
			resourceSingleUrl: this.Update,
		},
		extension.MethodDelete: {
			resourceSingleUrl: this.Delete,
		},
	}
	return routeMap
}

func (this *TestPatternController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleUrl := extension.GetResourceSingleUrl(this.ResourceName)
	resourceMultiUrl := extension.GetResourceMultiUrl(this.ResourceName)

	routeMap := map[int]map[string]gin.HandlerFunc{
		extension.MethodGet: {
			resourceSingleUrl: this.GetSingle,
			resourceMultiUrl:  this.GetMulti,
		},
		extension.MethodPost: {
			resourceMultiUrl: this.Create,
		},
		extension.MethodPut: {
			resourceSingleUrl: this.Update,
		},
		extension.MethodDelete: {
			resourceSingleUrl: this.Delete,
		},
	}
	return routeMap
}

func (this *TestCaseController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
	resourceSingleUrl := extension.GetResourceSingleUrl(this.ResourceName)
	resourceMultiUrl := extension.GetResourceMultiUrl(this.ResourceName)

	routeMap := map[int]map[string]gin.HandlerFunc{
		extension.MethodGet: {
			resourceSingleUrl: this.GetSingle,
			resourceMultiUrl:  this.GetMulti,
		},
		extension.MethodPost: {
			resourceMultiUrl: this.Create,
		},
		extension.MethodPut: {
			resourceSingleUrl: this.Update,
		},
		extension.MethodDelete: {
			resourceSingleUrl: this.Delete,
		},
		extension.MethodPatch: {
			resourceSingleUrl: this.Patch,
		},
	}
	return routeMap
}

func (this *TestCaseController) OutputPatch(c *gin.Context, code int, result interface{}) {
	text := result.(string)
	c.String(code, text)
}
