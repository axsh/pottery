package models

import (
	"github.com/qb0C80aE/clay/extension"
)

type TestCommand struct {
	ID                   int    `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ServiceName          string `json:"service_name" form:"service_name" gorm:"not null;unique"`
	SerevrScriptTemplate string `json:"server_script_template" form:"server_script_template"`
	ClientScriptTemplate string `json:"client_script_template" form:"client_script_template"`
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

var TestCommandModel = &TestCommand{}
var TestPatternModel = &TestPattern{}
var TestCaseModel = &TestCase{}

func init() {
	extension.RegisterModelType(TestCommand{})
	extension.RegisterModelType(TestPattern{})
	extension.RegisterModelType(TestCase{})
}
