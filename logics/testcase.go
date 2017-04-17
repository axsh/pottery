package logics

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
	"strconv"
	tplpkg "text/template"
)

type testCommandLogic struct {
	*clayLogics.BaseLogic
}

type testPatternLogic struct {
	*clayLogics.BaseLogic
}

type testCaseLogic struct {
	*clayLogics.BaseLogic
}

func newTestCommandLogic() *testCommandLogic {
	logic := &testCommandLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func newTestPatternLogic() *testPatternLogic {
	logic := &testPatternLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func newTestCaseLogic() *testCaseLogic {
	logic := &testCaseLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func (logic *testCommandLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testCommand := &models.TestCommand{}

	if err := db.Select(queryFields).First(testCommand, id).Error; err != nil {
		return nil, err
	}

	return testCommand, nil

}

func (logic *testCommandLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	testCommands := []*models.TestCommand{}

	if err := db.Select(queryFields).Find(&testCommands).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testCommands))
	for i, data := range testCommands {
		result[i] = data
	}

	return result, nil

}

func (logic *testCommandLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	testCommand := data.(*models.TestCommand)

	if err := db.Create(&testCommand).Error; err != nil {
		return nil, err
	}

	return testCommand, nil
}

func (logic *testCommandLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testCommand := data.(*models.TestCommand)
	testCommand.ID, _ = strconv.Atoi(id)

	if err := db.Save(testCommand).Error; err != nil {
		return nil, err
	}

	return testCommand, nil
}

func (logic *testCommandLogic) Delete(db *gorm.DB, id string) error {

	testCommand := &models.TestCommand{}

	if err := db.First(&testCommand, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testCommand).Error; err != nil {
		return err
	}

	return nil

}

func (logic *testCommandLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	testCommands := []*models.TestCommand{}
	if err := db.Select("*").Find(&testCommands).Error; err != nil {
		return "", nil, err
	}
	return "test_commands", testCommands, nil
}

func (logic *testCommandLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from test_commands;").Error
}

func (logic *testCommandLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.TestCommand{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["test_commands"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, testCommand := range container {
			if err := db.Create(testCommand).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (logic *testCommandLogic) GenerateTemplateParameter(db *gorm.DB) (string, interface{}, error) {
	testCommands := []*models.TestCommand{}
	if err := db.Select("*").Find(&testCommands).Error; err != nil {
		return "", nil, err
	}
	return "TestCommands", testCommands, nil
}

func (logic *testPatternLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testPattern := &models.TestPattern{}

	if err := db.Select(queryFields).First(testPattern, id).Error; err != nil {
		return nil, err
	}

	return testPattern, nil

}

func (logic *testPatternLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	testPatterns := []*models.TestPattern{}

	if err := db.Select(queryFields).Find(&testPatterns).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testPatterns))
	for i, data := range testPatterns {
		result[i] = data
	}

	return result, nil

}

func (logic *testPatternLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	testPattern := data.(*models.TestPattern)

	if err := db.Create(&testPattern).Error; err != nil {
		return nil, err
	}

	return testPattern, nil
}

func (logic *testPatternLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testPattern := data.(*models.TestPattern)
	testPattern.ID, _ = strconv.Atoi(id)

	if err := db.Save(testPattern).Error; err != nil {
		return nil, err
	}

	return testPattern, nil
}

func (logic *testPatternLogic) Delete(db *gorm.DB, id string) error {

	testPattern := &models.TestPattern{}

	if err := db.First(&testPattern, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testPattern).Error; err != nil {
		return err
	}

	return nil

}

func (logic *testPatternLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	testPatterns := []*models.TestPattern{}
	if err := db.Select("*").Find(&testPatterns).Error; err != nil {
		return "", nil, err
	}
	return "test_patterns", testPatterns, nil
}

func (logic *testPatternLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from test_patterns;").Error
}

func (logic *testPatternLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.TestPattern{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["test_patterns"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, testPatterns := range container {
			if err := db.Create(testPatterns).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (logic *testPatternLogic) GenerateTemplateParameter(db *gorm.DB) (string, interface{}, error) {
	testPatterns := []*models.TestPattern{}
	if err := db.Select("*").Find(&testPatterns).Error; err != nil {
		return "", nil, err
	}
	return "TestPatterns", testPatterns, nil
}

func (logic *testCaseLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testCase := &models.TestCase{}

	if err := db.Select(queryFields).First(testCase, id).Error; err != nil {
		return nil, err
	}

	return testCase, nil

}

func (logic *testCaseLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	testCases := []*models.TestCase{}

	if err := db.Select(queryFields).Find(&testCases).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testCases))
	for i, data := range testCases {
		result[i] = data
	}

	return result, nil

}

func (logic *testCaseLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	testCase := data.(*models.TestCase)

	if err := db.Create(&testCase).Error; err != nil {
		return nil, err
	}

	return testCase, nil
}

func (logic *testCaseLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testCase := data.(*models.TestCase)
	testCase.ID, _ = strconv.Atoi(id)

	if err := db.Save(testCase).Error; err != nil {
		return nil, err
	}

	return testCase, nil
}

func (logic *testCaseLogic) Delete(db *gorm.DB, id string) error {

	testCase := &models.TestCase{}

	if err := db.First(&testCase, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testCase).Error; err != nil {
		return err
	}

	return nil

}

func (logic *testCaseLogic) Patch(db *gorm.DB, id string) (interface{}, error) {
	testRunnerScript, testCommands, err := generateTestScripts(db, id)
	if err != nil {
		return "", err
	}
	a := []*models.TestCommand(testCommands)
	c := bytes.Buffer{}
	c.WriteString("# --- test runner script ---\n")
	c.WriteString(fmt.Sprintf("# %s\n", testRunnerScript))
	c.WriteString("#--------------------------------------\n")
	for _, b := range a {
		c.WriteString("#--------------------------------------\n")
		c.WriteString(fmt.Sprintf("# %s\n", b.ServiceName))
		c.WriteString("# --- server script ---\n")
		c.WriteString(fmt.Sprintf("%s\n", b.SerevrScriptTemplate))
		c.WriteString("# --- client script ---\n")
		c.WriteString(fmt.Sprintf("%s\n", b.ClientScriptTemplate))
		c.WriteString("#--------------------------------------\n")
	}
	return c.String(), nil
}

func (logic *testCaseLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	testCases := []*models.TestCase{}
	if err := db.Select("*").Find(&testCases).Error; err != nil {
		return "", nil, err
	}
	return "test_cases", testCases, nil
}

func (logic *testCaseLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from test_cases;").Error
}

func (logic *testCaseLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.TestCase{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content["test_cases"]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, testCase := range container {
			if err := db.Create(testCase).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

func (logic *testCaseLogic) GenerateTemplateParameter(db *gorm.DB) (string, interface{}, error) {
	testCases := []*models.TestCase{}
	if err := db.Select("*").Find(&testCases).Error; err != nil {
		return "", nil, err
	}
	return "TestCases", testCases, nil
}

func convertAccessibility(accessibility bool) string {
	if accessibility {
		return "allow"
	} else {
		return "deny"
	}
}

func generateTestScripts(db *gorm.DB, id string) (string, []*models.TestCommand, error) {

	result := []*models.TestCommand{}

	testCase := &models.TestCase{}
	if err := db.Preload("TestPatterns").
		Preload("TestPatterns.TestCommand").First(&testCase, id).Error; err != nil {
		return "", result, err
	}

	testCommandMap := make(map[string]*models.TestCommand)
	for _, testPattern := range testCase.TestPatterns {
		testCommandMap[testPattern.TestCommand.ServiceName] = testPattern.TestCommand
	}

	requirements := []*models.Requirement{}
	if err := db.Preload("Service").
		Preload("Service.Connections").
		Preload("SourcePort").
		Preload("SourcePort.Node").
		Preload("DestinationPort").
		Preload("DestinationPort.Node").Select("*").Find(&requirements).Error; err != nil {
		return "", result, err
	}

	anyNode := &loamModels.Node{
		ID:   0,
		Name: "Any",
	}
	anyPort := &loamModels.Port{
		ID:     0,
		Number: 0,
		Layer:  3,
		Name:   "Any",
		NodeID: anyNode.ID,
		Node:   anyNode,
		MacAddress: sql.NullString{
			String: "00:00:00:00:00:00",
			Valid:  true,
		},
		Ipv4Address: sql.NullString{
			String: "0.0.0.0",
			Valid:  true,
		},
		Ipv4Prefix: sql.NullInt64{
			Int64: 0,
			Valid: true,
		},
	}
	anyNode.Ports = []*loamModels.Port{anyPort}

	templateFuncMaps := extensions.RegisteredTemplateFuncMaps()

	for _, requirement := range requirements {
		if !requirement.SourcePortID.Valid {
			requirement.SourcePort = anyPort
		}
		if !requirement.DestinationPortID.Valid {
			requirement.DestinationPort = anyPort
		}

		testCommand := testCommandMap[requirement.Service.Name]
		serverScript := testCommand.SerevrScriptTemplate
		clientScript := testCommand.ClientScriptTemplate

		var docServerScript bytes.Buffer
		tplServerScript := tplpkg.New("template_server_script")
		for _, templateFuncMap := range templateFuncMaps {
			tplServerScript = tplServerScript.Funcs(templateFuncMap)
		}
		tplServerScript, err := tplServerScript.Parse(serverScript)
		if err != nil {
			return "", nil, err
		}
		err = tplServerScript.Execute(&docServerScript, requirement)
		if err != nil {
			return "", nil, err
		}
		serverScript = docServerScript.String()

		var docClientScript bytes.Buffer
		tplClientScript := tplpkg.New("template_server_script")
		for _, templateFuncMap := range templateFuncMaps {
			tplClientScript = tplClientScript.Funcs(templateFuncMap)
		}
		tplClientScript, err = tplClientScript.Parse(clientScript)
		if err != nil {
			return "", nil, err
		}
		err = tplClientScript.Execute(&docClientScript, requirement)
		if err != nil {
			return "", nil, err
		}
		clientScript = docClientScript.String()

		newTestCommand := &models.TestCommand{
			ServiceName:          fmt.Sprintf("%s_to_%s_%s_%s", requirement.SourcePort.Node.Name, requirement.DestinationPort.Node.Name, requirement.Service.Name, convertAccessibility(requirement.Accessibility)),
			SerevrScriptTemplate: serverScript,
			ClientScriptTemplate: clientScript,
		}
		result = append(result, newTestCommand)
	}

	return testCase.TestRunnerScriptTemplate, result, nil

}

var uniqueTestCommandLogic = newTestCommandLogic()
var uniqueTestPatternLogic = newTestPatternLogic()
var uniqueTestCaseLogic = newTestCaseLogic()

func UniqueTestCommandLogic() extensions.Logic {
	return uniqueTestCommandLogic
}

func UniqueTestPatternLogic() extensions.Logic {
	return uniqueTestPatternLogic
}

func UniqueTestCaseLogic() extensions.Logic {
	return uniqueTestCaseLogic
}

func init() {
	extensions.RegisterDesignAccessor(uniqueTestCommandLogic)
	extensions.RegisterDesignAccessor(uniqueTestPatternLogic)
	extensions.RegisterDesignAccessor(uniqueTestCaseLogic)
	extensions.RegisterTemplateParameterGenerator(uniqueTestCommandLogic)
	extensions.RegisterTemplateParameterGenerator(uniqueTestPatternLogic)
	extensions.RegisterTemplateParameterGenerator(uniqueTestCaseLogic)
}
