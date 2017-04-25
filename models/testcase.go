package models

import (
	"github.com/qb0C80aE/clay/extensions"
	clayModels "github.com/qb0C80aE/clay/models"
)

type TestCommand struct {
	ID                     int                  `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ServiceName            string               `json:"service_name" form:"service_name" gorm:"not null;unique"`
	ServerScriptTemplateID int                  `json:"server_script_template_id"`
	ServerScriptTemplate   *clayModels.Template `json:"server_script_template"`
	ClientScriptTemplateID int                  `json:"client_script_template_id"`
	ClientScriptTemplate   *clayModels.Template `json:"client_script_template"`
}

type TestPattern struct {
	ID            int          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	TestCommandID int          `json:"test_command_id" gorm:"index" sql:"type:integer references test_commands(id) on delete cascade"`
	TestCommand   *TestCommand `json:"test_command"`
	TestCaseID    int          `json:"test_case_id" gorm:"index" sql:"type:integer references test_cases(id) on delete cascade"`
	TestCase      *TestCase    `json:"test_case"`
}

type TestCase struct {
	ID                       int            `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"`
	TestPatterns             []*TestPattern `json:"test_patterns"`
	TestRunnerScriptTemplate string         `json:"test_runner_script_template" form:"test_runner_script_template"`
}

func NewTestCommandModel() *TestCommand {
	return &TestCommand{}
}

func NewTestPatternModel() *TestPattern {
	return &TestPattern{}
}

func NewTestCaseModel() *TestCase {
	return &TestCase{}
}

var sharedTestCommandModel = NewTestCommandModel()
var sharedTestPatternModel = NewTestPatternModel()
var sharedTestCaseModel = NewTestCaseModel()

func SharedTestCommandModel() *TestCommand {
	return sharedTestCommandModel
}

func SharedTestPatternModel() *TestPattern {
	return sharedTestPatternModel
}

func SharedTestCaseModel() *TestCase {
	return sharedTestCaseModel
}

func init() {
	extensions.RegisterModel("TestCommand", sharedTestCommandModel)
	extensions.RegisterModel("TestPattern", sharedTestPatternModel)
	extensions.RegisterModel("TestCase", sharedTestCaseModel)
}
