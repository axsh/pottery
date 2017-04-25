package models

import (
	"github.com/qb0C80aE/clay/extensions"
	clayModels "github.com/qb0C80aE/clay/models"
)

type TestProgram struct {
	ID                     int                  `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ServiceID              int                  `json:"service_id" gorm:"not null;unique" sql:"type:integer references services(id) on delete cascade"`
	Service                *Service             `json:"service"`
	ServerScriptTemplateID int                  `json:"server_script_template_id"`
	ServerScriptTemplate   *clayModels.Template `json:"server_script_template"`
	ClientScriptTemplateID int                  `json:"client_script_template_id"`
	ClientScriptTemplate   *clayModels.Template `json:"client_script_template"`
}

func NewTestProgramModel() *TestProgram {
	return &TestProgram{}
}

var sharedTestProgramModel = NewTestProgramModel()

func SharedTestProgramModel() *TestProgram {
	return sharedTestProgramModel
}

func init() {
	extensions.RegisterModel("TestProgram", sharedTestProgramModel)
}
