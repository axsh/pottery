package controllers

import (
	"github.com/gin-gonic/gin"
	clayControllers "github.com/qb0C80aE/clay/controllers"
	"github.com/qb0C80aE/clay/extension"
	"github.com/qb0C80aE/pottery/logics"
	"github.com/qb0C80aE/pottery/models"
)

type ProtocolController struct {
	clayControllers.BaseController
}

type ServiceController struct {
	clayControllers.BaseController
}

type ConnectionController struct {
	clayControllers.BaseController
}

type RequirementController struct {
	clayControllers.BaseController
}

func init() {
	extension.RegisterController(NewProtocolController())
	extension.RegisterController(NewServiceController())
	extension.RegisterController(NewConnectionController())
	extension.RegisterController(NewRequirementController())
}

func NewProtocolController() *ProtocolController {
	controller := &ProtocolController{}
	controller.Initialize()
	return controller
}

func NewServiceController() *ServiceController {
	controller := &ServiceController{}
	controller.Initialize()
	return controller
}

func NewConnectionController() *ConnectionController {
	controller := &ConnectionController{}
	controller.Initialize()
	return controller
}

func NewRequirementController() *RequirementController {
	controller := &RequirementController{}
	controller.Initialize()
	return controller
}

func (this *ProtocolController) Initialize() {
	this.ResourceName = "protocol"
	this.Model = models.ProtocolModel
	this.Logic = logics.ProtocolLogicInstance
	this.Outputter = this
}

func (this *ServiceController) Initialize() {
	this.ResourceName = "service"
	this.Model = models.ServiceModel
	this.Logic = logics.ServiceLogicInstance
	this.Outputter = this
}

func (this *ConnectionController) Initialize() {
	this.ResourceName = "connection"
	this.Model = models.ConnectionModel
	this.Logic = logics.ConnectionLogicInstance
	this.Outputter = this
}

func (this *RequirementController) Initialize() {
	this.ResourceName = "requirement"
	this.Model = models.RequirementModel
	this.Logic = logics.RequirementLogicInstance
	this.Outputter = this
}

func (this *ProtocolController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
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

func (this *ServiceController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
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

func (this *ConnectionController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
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

func (this *RequirementController) GetRouteMap() map[int]map[string]gin.HandlerFunc {
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
