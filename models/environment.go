package models

import (
	"github.com/qb0C80aE/clay/extensions"
	clayModels "github.com/qb0C80aE/clay/models"
)

type Environment struct {
	ID                    int                  `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	TemplateID            int                  `json:"template_id" gorm:"index" sql:"type:integer references templates(id) on delete set null"`
	Template              *clayModels.Template `json:"template"`
	TestRunnerScriptID    int                  `json:"test_runner_script_id" gorm:"index" sql:"type:integer references templates(id) on delete set null"`
	TestRunnerScript      *clayModels.Template `json:"test_runner_script"`
	GitRepositoryURI      string               `json:"git_repository_uri" gorm:"not null"`
	GitUserName           string               `json:"git_user_name" gorm:"not null"`
	GitUserEmail          string               `json:"git_user_email" gorm:"not null"`
	DesignFileName        string               `json:"design_file_name" gorm:"not null"`
	TemplateFileName      string               `json:"template_file_name" gorm:"not null"`
	TestCaseDirectoryName string               `json:"test_case_directory_name" gorm:"not null"`
}

func NewEnvironmentModel() *Environment {
	return &Environment{}
}

var sharedEnvironmentModel = NewEnvironmentModel()

func SharedEnvironmentModel() *Environment {
	return sharedEnvironmentModel
}

func init() {
	extensions.RegisterModel(sharedEnvironmentModel)
}
