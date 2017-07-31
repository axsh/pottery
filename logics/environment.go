package logics

import (
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	"github.com/qb0C80aE/pottery/models"
	"net/url"
)

type environmentLogic struct {
	*clayLogics.BaseLogic
}

func newEnvironmentLogic() *environmentLogic {
	logic := &environmentLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func (logic *environmentLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {

	environment := &models.Environment{}

	if err := db.Select(queryFields).First(environment, 1).Error; err != nil {
		return nil, err
	}

	return environment, nil
}

func (logic *environmentLogic) Update(db *gorm.DB, _ string, _ url.Values, data interface{}) (interface{}, error) {

	environment := data.(*models.Environment)
	environment.ID = models.PresentEnvironmentID

	if err := db.Save(&environment).Error; err != nil {
		return nil, err
	}

	return environment, nil

}

func (logic *environmentLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	environments := []*models.Environment{}
	if err := db.Select("*").Find(&environments).Error; err != nil {
		return "", nil, err
	}
	return extensions.RegisteredResourceName(models.SharedEnvironmentModel()), environments, nil
}

func (logic *environmentLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Delete(models.SharedEnvironmentModel()).Error
}

func (logic *environmentLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.Environment{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content[extensions.RegisteredResourceName(models.SharedEnvironmentModel())]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, environment := range container {
			environment.TestRunnerScript = nil
			if err := db.Create(environment).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

var uniqueEnvironmentLogic = newEnvironmentLogic()

func UniqueEnvironmentLogic() extensions.Logic {
	return uniqueEnvironmentLogic
}

func init() {
	extensions.RegisterDesignAccessor(uniqueEnvironmentLogic)
	extensions.RegisterTemplateParameterGenerator(models.SharedEnvironmentModel(), uniqueEnvironmentLogic)
}
