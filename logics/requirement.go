package logics

import (
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	"github.com/qb0C80aE/pottery/models"
	"strconv"
)

type protocolLogic struct {
	*clayLogics.BaseLogic
}

type serviceLogic struct {
	*clayLogics.BaseLogic
}

type connectionLogic struct {
	*clayLogics.BaseLogic
}

type requirementLogic struct {
	*clayLogics.BaseLogic
}

func newProtocolLogic() *protocolLogic {
	logic := &protocolLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func newServiceLogic() *serviceLogic {
	logic := &serviceLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func newConnectionLogic() *connectionLogic {
	logic := &connectionLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func newRequirementLogic() *requirementLogic {
	logic := &requirementLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func (logic *protocolLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	protocol := &models.Protocol{}

	if err := db.Select(queryFields).First(protocol, id).Error; err != nil {
		return nil, err
	}

	return protocol, nil

}

func (logic *protocolLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	protocols := []*models.Protocol{}

	if err := db.Select(queryFields).Find(&protocols).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(protocols))
	for i, data := range protocols {
		result[i] = data
	}

	return result, nil

}

func (logic *protocolLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)

	if err := db.Create(&protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func (logic *protocolLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)
	protocol.ID, _ = strconv.Atoi(id)

	if err := db.Save(protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func (logic *protocolLogic) Delete(db *gorm.DB, id string) error {

	protocol := &models.Protocol{}

	if err := db.First(&protocol, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&protocol).Error; err != nil {
		return err
	}

	return nil

}

func (logic *protocolLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	protocols := []*models.Protocol{}
	if err := db.Select("*").Find(&protocols).Error; err != nil {
		return "", nil, err
	}
	return "protocols", protocols, nil
}

func (logic *protocolLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from protocols;").Error
}

func (logic *protocolLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Protocol{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["protocols"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, protocol := range container {
			if err := db.Create(protocol).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (logic *protocolLogic) GenerateTemplateParameter(db *gorm.DB) (string, interface{}, error) {
	protocols := []*models.Protocol{}
	if err := db.Select("*").Find(&protocols).Error; err != nil {
		return "", nil, err
	}
	return "Protocols", protocols, nil
}

func (logic *serviceLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	service := &models.Service{}

	if err := db.Select(queryFields).First(service, id).Error; err != nil {
		return nil, err
	}

	return service, nil

}

func (logic *serviceLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	services := []*models.Service{}

	if err := db.Select(queryFields).Find(&services).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(services))
	for i, data := range services {
		result[i] = data
	}

	return result, nil

}

func (logic *serviceLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	service := data.(*models.Service)

	if err := db.Create(&service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (logic *serviceLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	service := data.(*models.Service)
	service.ID, _ = strconv.Atoi(id)

	if err := db.Save(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (logic *serviceLogic) Delete(db *gorm.DB, id string) error {

	service := &models.Service{}

	if err := db.First(&service, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&service).Error; err != nil {
		return err
	}

	return nil

}

func (logic *serviceLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	services := []*models.Service{}
	if err := db.Select("*").Find(&services).Error; err != nil {
		return "", nil, err
	}
	return "services", services, nil
}

func (logic *serviceLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from services;").Error
}

func (logic *serviceLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Service{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["services"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, service := range container {
			if err := db.Create(service).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (logic *serviceLogic) GenerateTemplateParameter(db *gorm.DB) (string, interface{}, error) {
	services := []*models.Service{}
	if err := db.Select("*").Find(&services).Error; err != nil {
		return "", nil, err
	}
	return "Services", services, nil
}

func (logic *connectionLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	connection := &models.Connection{}

	if err := db.Select(queryFields).First(connection, id).Error; err != nil {
		return nil, err
	}

	return connection, nil

}

func (logic *connectionLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	connections := []*models.Connection{}

	if err := db.Select(queryFields).Find(&connections).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(connections))
	for i, data := range connections {
		result[i] = data
	}

	return result, nil

}

func (logic *connectionLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)

	if err := db.Create(&connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func (logic *connectionLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)
	connection.ID, _ = strconv.Atoi(id)

	if err := db.Save(connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func (logic *connectionLogic) Delete(db *gorm.DB, id string) error {

	connection := &models.Connection{}

	if err := db.First(&connection, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&connection).Error; err != nil {
		return err
	}

	return nil

}

func (logic *connectionLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	connections := []*models.Connection{}
	if err := db.Select("*").Find(&connections).Error; err != nil {
		return "", nil, err
	}
	return "connections", connections, nil
}

func (logic *connectionLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from connections;").Error
}

func (logic *connectionLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Connection{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["connections"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, connection := range container {
			if err := db.Create(connection).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (logic *connectionLogic) GenerateTemplateParameter(db *gorm.DB) (string, interface{}, error) {
	connections := []*models.Connection{}
	if err := db.Select("*").Find(&connections).Error; err != nil {
		return "", nil, err
	}
	return "Connections", connections, nil
}

func (logic *requirementLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	requirement := &models.Requirement{}

	if err := db.Select(queryFields).First(requirement, id).Error; err != nil {
		return nil, err
	}

	return requirement, nil

}

func (logic *requirementLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	requirements := []*models.Requirement{}

	if err := db.Select(queryFields).Find(&requirements).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(requirements))
	for i, data := range requirements {
		result[i] = data
	}

	return result, nil

}

func (logic *requirementLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)

	if err := db.Create(&requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func (logic *requirementLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)
	requirement.ID, _ = strconv.Atoi(id)

	if err := db.Save(requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func (logic *requirementLogic) Delete(db *gorm.DB, id string) error {

	requirement := &models.Requirement{}

	if err := db.First(&requirement, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&requirement).Error; err != nil {
		return err
	}

	return nil

}

func (logic *requirementLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	requirements := []*models.Requirement{}
	if err := db.Select("*").Find(&requirements).Error; err != nil {
		return "", nil, err
	}
	return "requirements", requirements, nil
}

func (logic *requirementLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from requirements;").Error
}

func (logic *requirementLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Requirement{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["requirements"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, requirement := range container {
			if err := db.Create(requirement).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (logic *requirementLogic) GenerateTemplateParameter(db *gorm.DB) (string, interface{}, error) {
	requirements := []*models.Requirement{}
	if err := db.Select("*").Find(&requirements).Error; err != nil {
		return "", nil, err
	}
	return "Requirements", requirements, nil
}

var uniqueProtocolLogic = newProtocolLogic()
var uniqueServiceLogic = newServiceLogic()
var uniqueConnectionLogic = newConnectionLogic()
var uniqueRequirementLogic = newRequirementLogic()

func UniqueProtocolLogic() extensions.Logic {
	return uniqueProtocolLogic
}

func UniqueServiceLogic() extensions.Logic {
	return uniqueServiceLogic
}

func UniqueConnectionLogic() extensions.Logic {
	return uniqueConnectionLogic
}

func UniqueRequirementLogic() extensions.Logic {
	return uniqueRequirementLogic
}

func init() {
	extensions.RegisterDesignAccessor(uniqueProtocolLogic)
	extensions.RegisterDesignAccessor(uniqueServiceLogic)
	extensions.RegisterDesignAccessor(uniqueConnectionLogic)
	extensions.RegisterDesignAccessor(uniqueRequirementLogic)
	extensions.RegisterTemplateParameterGenerator(uniqueProtocolLogic)
	extensions.RegisterTemplateParameterGenerator(uniqueServiceLogic)
	extensions.RegisterTemplateParameterGenerator(uniqueConnectionLogic)
	extensions.RegisterTemplateParameterGenerator(uniqueRequirementLogic)
}
