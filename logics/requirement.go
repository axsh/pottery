package logics

import (
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extension"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	"github.com/qb0C80aE/pottery/models"
	"strconv"
)

type ProtocolLogic struct {
}

type ServiceLogic struct {
}

type ConnectionLogic struct {
}

type RequirementLogic struct {
}

func (_ *ProtocolLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	protocol := &models.Protocol{}

	if err := db.Select(queryFields).First(protocol, id).Error; err != nil {
		return nil, err
	}

	return protocol, nil

}

func (_ *ProtocolLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

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

func (_ *ProtocolLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)

	if err := db.Create(&protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func (_ *ProtocolLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)
	protocol.ID, _ = strconv.Atoi(id)

	if err := db.Save(protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func (_ *ProtocolLogic) Delete(db *gorm.DB, id string) error {

	protocol := &models.Protocol{}

	if err := db.First(&protocol, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&protocol).Error; err != nil {
		return err
	}

	return nil

}

func (_ *ProtocolLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *ProtocolLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *ProtocolLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	protocols := []*models.Protocol{}
	if err := db.Select("*").Find(&protocols).Error; err != nil {
		return "", nil, err
	}
	return "protocols", protocols, nil
}

func (_ *ProtocolLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from protocols;").Error
}

func (_ *ProtocolLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
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

func (_ *ServiceLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	service := &models.Service{}

	if err := db.Select(queryFields).First(service, id).Error; err != nil {
		return nil, err
	}

	return service, nil

}

func (_ *ServiceLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

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

func (_ *ServiceLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	service := data.(*models.Service)

	if err := db.Create(&service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (_ *ServiceLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	service := data.(*models.Service)
	service.ID, _ = strconv.Atoi(id)

	if err := db.Save(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (_ *ServiceLogic) Delete(db *gorm.DB, id string) error {

	service := &models.Service{}

	if err := db.First(&service, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&service).Error; err != nil {
		return err
	}

	return nil

}

func (_ *ServiceLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *ServiceLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *ServiceLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	services := []*models.Service{}
	if err := db.Select("*").Find(&services).Error; err != nil {
		return "", nil, err
	}
	return "services", services, nil
}

func (_ *ServiceLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from services;").Error
}

func (_ *ServiceLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
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

func (_ *ConnectionLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	connection := &models.Connection{}

	if err := db.Select(queryFields).First(connection, id).Error; err != nil {
		return nil, err
	}

	return connection, nil

}

func (_ *ConnectionLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

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

func (_ *ConnectionLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)

	if err := db.Create(&connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func (_ *ConnectionLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)
	connection.ID, _ = strconv.Atoi(id)

	if err := db.Save(connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func (_ *ConnectionLogic) Delete(db *gorm.DB, id string) error {

	connection := &models.Connection{}

	if err := db.First(&connection, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&connection).Error; err != nil {
		return err
	}

	return nil

}

func (_ *ConnectionLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *ConnectionLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *ConnectionLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	connections := []*models.Connection{}
	if err := db.Select("*").Find(&connections).Error; err != nil {
		return "", nil, err
	}
	return "connections", connections, nil
}

func (_ *ConnectionLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from connections;").Error
}

func (_ *ConnectionLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
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

func (_ *RequirementLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	requirement := &models.Requirement{}

	if err := db.Select(queryFields).First(requirement, id).Error; err != nil {
		return nil, err
	}

	return requirement, nil

}

func (_ *RequirementLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

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

func (_ *RequirementLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)

	if err := db.Create(&requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func (_ *RequirementLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)
	requirement.ID, _ = strconv.Atoi(id)

	if err := db.Save(requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func (_ *RequirementLogic) Delete(db *gorm.DB, id string) error {

	requirement := &models.Requirement{}

	if err := db.First(&requirement, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&requirement).Error; err != nil {
		return err
	}

	return nil

}

func (_ *RequirementLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *RequirementLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *RequirementLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	requirements := []*models.Requirement{}
	if err := db.Select("*").Find(&requirements).Error; err != nil {
		return "", nil, err
	}
	return "requirements", requirements, nil
}

func (_ *RequirementLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from requirements;").Error
}

func (_ *RequirementLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
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

var ProtocolLogicInstance = &ProtocolLogic{}
var ServiceLogicInstance = &ServiceLogic{}
var ConnectionLogicInstance = &ConnectionLogic{}
var RequirementLogicInstance = &RequirementLogic{}

func init() {
	extension.RegisterDesignAccessor(ProtocolLogicInstance)
	extension.RegisterDesignAccessor(ServiceLogicInstance)
	extension.RegisterDesignAccessor(ConnectionLogicInstance)
	extension.RegisterDesignAccessor(RequirementLogicInstance)
}
