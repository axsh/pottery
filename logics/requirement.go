package logics

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	"github.com/qb0C80aE/pottery/models"
	"net/url"
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

type firewallTestScriptGenerationLogic struct {
	*clayLogics.BaseLogic
}

type firewallFirewallTestProgramLogic struct {
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


func newFirewallTestScriptGenerationLogic() *firewallTestScriptGenerationLogic {
	logic := &firewallTestScriptGenerationLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func newFirewallTestProgramLogic() *firewallFirewallTestProgramLogic {
	logic := &firewallFirewallTestProgramLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func (logic *protocolLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {

	protocol := &models.Protocol{}

	if err := db.Select(queryFields).First(protocol, id).Error; err != nil {
		return nil, err
	}

	return protocol, nil

}

func (logic *protocolLogic) GetMulti(db *gorm.DB, _ url.Values, queryFields string) (interface{}, error) {

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

func (logic *protocolLogic) Create(db *gorm.DB, _ url.Values, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)

	if err := db.Create(&protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func (logic *protocolLogic) Update(db *gorm.DB, id string, _ url.Values, data interface{}) (interface{}, error) {

	protocol := data.(*models.Protocol)
	protocol.ID, _ = strconv.Atoi(id)

	if err := db.Save(protocol).Error; err != nil {
		return nil, err
	}

	return protocol, nil
}

func (logic *protocolLogic) Delete(db *gorm.DB, id string, _ url.Values) error {

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
	return extensions.RegisteredResourceName(models.SharedProtocolModel()), protocols, nil
}

func (logic *protocolLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Delete(models.SharedProtocolModel()).Error
}

func (logic *protocolLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Protocol{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content[extensions.RegisteredResourceName(models.SharedProtocolModel())]; exists {
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

func (logic *serviceLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {

	service := &models.Service{}

	if err := db.Select(queryFields).First(service, id).Error; err != nil {
		return nil, err
	}

	return service, nil

}

func (logic *serviceLogic) GetMulti(db *gorm.DB, _ url.Values, queryFields string) (interface{}, error) {

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

func (logic *serviceLogic) Create(db *gorm.DB, _ url.Values, data interface{}) (interface{}, error) {

	service := data.(*models.Service)

	if err := db.Create(&service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (logic *serviceLogic) Update(db *gorm.DB, id string, _ url.Values, data interface{}) (interface{}, error) {

	service := data.(*models.Service)
	service.ID, _ = strconv.Atoi(id)

	if err := db.Save(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

func (logic *serviceLogic) Delete(db *gorm.DB, id string, _ url.Values) error {

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
	return extensions.RegisteredResourceName(models.SharedServiceModel()), services, nil
}

func (logic *serviceLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Delete(models.SharedServiceModel()).Error
}

func (logic *serviceLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Service{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content[extensions.RegisteredResourceName(models.SharedServiceModel())]; exists {
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

func (logic *connectionLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {

	connection := &models.Connection{}

	if err := db.Select(queryFields).First(connection, id).Error; err != nil {
		return nil, err
	}

	return connection, nil

}

func (logic *connectionLogic) GetMulti(db *gorm.DB, _ url.Values, queryFields string) (interface{}, error) {

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

func (logic *connectionLogic) Create(db *gorm.DB, _ url.Values, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)

	if err := db.Create(&connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func (logic *connectionLogic) Update(db *gorm.DB, id string, _ url.Values, data interface{}) (interface{}, error) {

	connection := data.(*models.Connection)
	connection.ID, _ = strconv.Atoi(id)

	if err := db.Save(connection).Error; err != nil {
		return nil, err
	}

	return connection, nil
}

func (logic *connectionLogic) Delete(db *gorm.DB, id string, _ url.Values) error {

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
	return extensions.RegisteredResourceName(models.SharedConnectionModel()), connections, nil
}

func (logic *connectionLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Delete(models.SharedConnectionModel()).Error
}

func (logic *connectionLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Connection{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content[extensions.RegisteredResourceName(models.SharedConnectionModel())]; exists {
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

func (logic *requirementLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {

	requirement := &models.Requirement{}

	if err := db.Select(queryFields).First(requirement, id).Error; err != nil {
		return nil, err
	}

	return requirement, nil

}

func (logic *requirementLogic) GetMulti(db *gorm.DB, _ url.Values, queryFields string) (interface{}, error) {

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

func (logic *requirementLogic) Create(db *gorm.DB, _ url.Values, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)

	if err := db.Create(&requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func (logic *requirementLogic) Update(db *gorm.DB, id string, _ url.Values, data interface{}) (interface{}, error) {

	requirement := data.(*models.Requirement)
	requirement.ID, _ = strconv.Atoi(id)

	if err := db.Save(requirement).Error; err != nil {
		return nil, err
	}

	return requirement, nil
}

func (logic *requirementLogic) Delete(db *gorm.DB, id string, _ url.Values) error {

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
	return extensions.RegisteredResourceName(models.SharedRequirementModel()), requirements, nil
}

func (logic *requirementLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Delete(models.SharedRequirementModel()).Error
}

func (logic *requirementLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Requirement{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content[extensions.RegisteredResourceName(models.SharedRequirementModel())]; exists {
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

func convertAccessString(access bool) string {
	if access {
		return "allow"
	} else {
		return "deny"
	}
}

func getTestProgram(db *gorm.DB, requirementID int) (*models.FirewallTestProgram, map[string]interface{}, string, error) {

	requirement := &models.Requirement{}
	if err := db.Preload("Service").
		Preload("Service.Connections").
		Preload("SourcePort").
		Preload("SourcePort.Node").
		Preload("DestinationPort").
		Preload("DestinationPort.Node").Select("*").First(requirement, requirementID).Error; err != nil {
		return nil, nil, "", err
	}

	testProgram := &models.FirewallTestProgram{}
	if err := db.Preload("Service").
		Preload("TestScriptTemplate").Select("*").First(testProgram, "service_id = ?", requirement.ServiceID).Error; err != nil {
		return nil, nil, "", err
	}

	templateParameterMap := map[string]interface{}{
		"Service":         requirement.Service,
		"SourcePort":      requirement.SourcePort,
		"DestinationPort": requirement.DestinationPort,
		"Access":          requirement.Access,
	}

	testPattern := fmt.Sprintf("%s_to_%s_%s_%s",
		requirement.SourcePort.Node.Name,
		requirement.DestinationPort.Node.Name,
		convertAccessString(requirement.Access),
		requirement.Service.Name)
	return testProgram, templateParameterMap, testPattern, nil
}

func GenerateFirewallTestScript(db *gorm.DB, requirementID int) (interface{}, string, error) {
	testProgram, templateParameterMap, testPattern, err := getTestProgram(db, requirementID)
	if err != nil {
		return "", "", err
	}

	script, err := clayLogics.GenerateTemplate(db, strconv.Itoa(testProgram.TestScriptTemplateID), templateParameterMap)
	if err != nil {
		return "", "", err
	}

	return script, testPattern, nil
}

func (logic *firewallTestScriptGenerationLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {
	requirementID, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}
	testTemplate, _, err := GenerateFirewallTestScript(db, requirementID)
	return testTemplate, err
}


func (logic *firewallFirewallTestProgramLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {

	testProgram := &models.FirewallTestProgram{}

	if err := db.Select(queryFields).First(testProgram, id).Error; err != nil {
		return nil, err
	}

	return testProgram, nil

}

func (logic *firewallFirewallTestProgramLogic) GetMulti(db *gorm.DB, _ url.Values, queryFields string) (interface{}, error) {

	testPrograms := []*models.FirewallTestProgram{}

	if err := db.Select(queryFields).Find(&testPrograms).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testPrograms))
	for i, data := range testPrograms {
		result[i] = data
	}

	return result, nil

}

func (logic *firewallFirewallTestProgramLogic) Create(db *gorm.DB, _ url.Values, data interface{}) (interface{}, error) {

	testProgram := data.(*models.FirewallTestProgram)

	if err := db.Create(&testProgram).Error; err != nil {
		return nil, err
	}

	return testProgram, nil
}

func (logic *firewallFirewallTestProgramLogic) Update(db *gorm.DB, id string, _ url.Values, data interface{}) (interface{}, error) {

	testProgram := data.(*models.FirewallTestProgram)
	testProgram.ID, _ = strconv.Atoi(id)

	if err := db.Save(testProgram).Error; err != nil {
		return nil, err
	}

	return testProgram, nil
}

func (logic *firewallFirewallTestProgramLogic) Delete(db *gorm.DB, id string, _ url.Values) error {

	testProgram := &models.FirewallTestProgram{}

	if err := db.First(&testProgram, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testProgram).Error; err != nil {
		return err
	}

	return nil
}

func (logic *firewallFirewallTestProgramLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	testPrograms := []*models.FirewallTestProgram{}
	if err := db.Select("*").Find(&testPrograms).Error; err != nil {
		return "", nil, err
	}
	return extensions.RegisteredResourceName(models.SharedFirewallTestProgramModel()), testPrograms, nil
}

func (logic *firewallFirewallTestProgramLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Delete(models.SharedFirewallTestProgramModel()).Error
}

func (logic *firewallFirewallTestProgramLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.FirewallTestProgram{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content[extensions.RegisteredResourceName(models.SharedFirewallTestProgramModel())]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, testProgram := range container {
			if err := db.Create(testProgram).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

var uniqueProtocolLogic = newProtocolLogic()
var uniqueServiceLogic = newServiceLogic()
var uniqueConnectionLogic = newConnectionLogic()
var uniqueRequirementLogic = newRequirementLogic()
var uniqueFirewallTestScriptGenerationLogic = newFirewallTestScriptGenerationLogic()
var uniqueFirewallTestProgramLogic = newFirewallTestProgramLogic()

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


func UniqueFirewallTestScriptGenerationLogic() extensions.Logic {
	return uniqueFirewallTestScriptGenerationLogic
}

func UniqueFirewallTestProgramLogic() extensions.Logic {
	return uniqueFirewallTestProgramLogic
}

func init() {
	extensions.RegisterDesignAccessor(uniqueProtocolLogic)
	extensions.RegisterDesignAccessor(uniqueServiceLogic)
	extensions.RegisterDesignAccessor(uniqueConnectionLogic)
	extensions.RegisterDesignAccessor(uniqueRequirementLogic)
	extensions.RegisterDesignAccessor(uniqueFirewallTestProgramLogic)
	extensions.RegisterTemplateParameterGenerator(models.SharedProtocolModel(), uniqueProtocolLogic)
	extensions.RegisterTemplateParameterGenerator(models.SharedServiceModel(), uniqueServiceLogic)
	extensions.RegisterTemplateParameterGenerator(models.SharedConnectionModel(), uniqueConnectionLogic)
	extensions.RegisterTemplateParameterGenerator(models.SharedRequirementModel(), uniqueRequirementLogic)
	extensions.RegisterTemplateParameterGenerator(models.SharedFirewallTestProgramModel(), uniqueFirewallTestProgramLogic)
}
