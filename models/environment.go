package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayModels "github.com/qb0C80aE/clay/models"
	loamModels "github.com/qb0C80aE/loam/models"
)

const (
	PresentEnvironmentID                               = 1
	NodeExtraAttributeField_ServerType                 = 3
	NodeExtraAttributeField_DeviceInitializationConfig = 4
	NodeExtraAttributeField_DeviceConfig               = 5
)

type Environment struct {
	ID                        int                  `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	TemplateID                int                  `json:"template_id" gorm:"index" sql:"type:integer references templates(id) on delete set null"`
	Template                  *clayModels.Template `json:"template"`
	TestRunnerScriptID        int                  `json:"test_runner_script_id" gorm:"index" sql:"type:integer references templates(id) on delete set null"`
	TestRunnerScript          *clayModels.Template `json:"test_runner_script"`
	GitRepositoryURI          string               `json:"git_repository_uri" gorm:"not null"`
	GitUserName               string               `json:"git_user_name" gorm:"not null"`
	GitUserEmail              string               `json:"git_user_email" gorm:"not null"`
	DesignFileName            string               `json:"design_file_name" gorm:"not null"`
	TemplateFileName          string               `json:"template_file_name" gorm:"not null"`
	TestCaseDirectoryName     string               `json:"test_case_directory_name" gorm:"not null"`
	ServerConfigDirectoryName string               `json:"server_config_directory_name" gorm:"not null"`
	DeviceConfigDirectoryName string               `json:"device_config_directory_name" gorm:"not null"`
}

func NewEnvironmentModel() *Environment {
	return &Environment{}
}

var sharedEnvironmentModel = NewEnvironmentModel()

func SharedEnvironmentModel() *Environment {
	return sharedEnvironmentModel
}

func (environment *Environment) SetupInitialData(db *gorm.DB) error {
	nodeExtraAttributeFields := []*loamModels.NodeExtraAttributeField{
		{ID: NodeExtraAttributeField_ServerType, Name: "server_type"},
		{ID: NodeExtraAttributeField_DeviceInitializationConfig, Name: "device_initialization_config"},
		{ID: NodeExtraAttributeField_DeviceConfig, Name: "device_config"},
	}

	for _, nodeExtraAttributeField := range nodeExtraAttributeFields {
		if err := db.Save(nodeExtraAttributeField).Error; err != nil {
			return err
		}
	}

	portExtraAttributeFields := []*loamModels.PortExtraAttributeField{
		{ID: 3, Name: "pass_through"},
	}

	for _, portExtraAttributeField := range portExtraAttributeFields {
		if err := db.Save(portExtraAttributeField).Error; err != nil {
			return err
		}
	}

	db.Exec(`
		create trigger if not exists DeleteServerInitializationConfigTemplate delete on node_extra_attribute_options when old.node_extra_attribute_field_id = 3
		begin
		 	delete from templates where id = old.value_int;
		end;
	`)

	return nil
}

func init() {
	extensions.RegisterInitialDataLoader(sharedEnvironmentModel)
	extensions.RegisterModel(sharedEnvironmentModel)
}
