package logics

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extension"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
	"strconv"
	tplpkg "text/template"
)

type TestCommandLogic struct {
}

type TestPatternLogic struct {
}

type TestCaseLogic struct {
}

func (_ *TestCommandLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testCommand := &models.TestCommand{}

	if err := db.Select(queryFields).First(testCommand, id).Error; err != nil {
		return nil, err
	}

	return testCommand, nil

}

func (_ *TestCommandLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

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

func (_ *TestCommandLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	testCommand := data.(*models.TestCommand)

	if err := db.Create(&testCommand).Error; err != nil {
		return nil, err
	}

	return testCommand, nil
}

func (_ *TestCommandLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testCommand := data.(*models.TestCommand)
	testCommand.ID, _ = strconv.Atoi(id)

	if err := db.Save(testCommand).Error; err != nil {
		return nil, err
	}

	return testCommand, nil
}

func (_ *TestCommandLogic) Delete(db *gorm.DB, id string) error {

	testCommand := &models.TestCommand{}

	if err := db.First(&testCommand, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testCommand).Error; err != nil {
		return err
	}

	return nil

}

func (_ *TestCommandLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *TestCommandLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *TestCommandLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	testCommands := []*models.TestCommand{}
	if err := db.Select("*").Find(&testCommands).Error; err != nil {
		return "", nil, err
	}
	return "test_commands", testCommands, nil
}

func (_ *TestCommandLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from test_commands;").Error
}

func (_ *TestCommandLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
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

func (_ *TestPatternLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testPattern := &models.TestPattern{}

	if err := db.Select(queryFields).First(testPattern, id).Error; err != nil {
		return nil, err
	}

	return testPattern, nil

}

func (_ *TestPatternLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

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

func (_ *TestPatternLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	testPattern := data.(*models.TestPattern)

	if err := db.Create(&testPattern).Error; err != nil {
		return nil, err
	}

	return testPattern, nil
}

func (_ *TestPatternLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testPattern := data.(*models.TestPattern)
	testPattern.ID, _ = strconv.Atoi(id)

	if err := db.Save(testPattern).Error; err != nil {
		return nil, err
	}

	return testPattern, nil
}

func (_ *TestPatternLogic) Delete(db *gorm.DB, id string) error {

	testPattern := &models.TestPattern{}

	if err := db.First(&testPattern, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testPattern).Error; err != nil {
		return err
	}

	return nil

}

func (_ *TestPatternLogic) Patch(_ *gorm.DB, _ string, _ string) (interface{}, error) {
	return nil, nil
}

func (_ *TestPatternLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *TestPatternLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	testPatterns := []*models.TestPattern{}
	if err := db.Select("*").Find(&testPatterns).Error; err != nil {
		return "", nil, err
	}
	return "test_patterns", testPatterns, nil
}

func (_ *TestPatternLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from test_patterns;").Error
}

func (_ *TestPatternLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
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

func (_ *TestCaseLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testCase := &models.TestCase{}

	if err := db.Select(queryFields).First(testCase, id).Error; err != nil {
		return nil, err
	}

	return testCase, nil

}

func (_ *TestCaseLogic) GetMulti(db *gorm.DB, queryFields string) ([]interface{}, error) {

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

func (_ *TestCaseLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	testCase := data.(*models.TestCase)

	if err := db.Create(&testCase).Error; err != nil {
		return nil, err
	}

	return testCase, nil
}

func (_ *TestCaseLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testCase := data.(*models.TestCase)
	testCase.ID, _ = strconv.Atoi(id)

	if err := db.Save(testCase).Error; err != nil {
		return nil, err
	}

	return testCase, nil
}

func (_ *TestCaseLogic) Delete(db *gorm.DB, id string) error {

	testCase := &models.TestCase{}

	if err := db.First(&testCase, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testCase).Error; err != nil {
		return err
	}

	return nil

}

func (_ *TestCaseLogic) Patch(db *gorm.DB, id string, _ string) (interface{}, error) {
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

func (_ *TestCaseLogic) Options(_ *gorm.DB) error {
	return nil
}

func (_ *TestCaseLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	testCases := []*models.TestCase{}
	if err := db.Select("*").Find(&testCases).Error; err != nil {
		return "", nil, err
	}
	return "test_cases", testCases, nil
}

func (_ *TestCaseLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Exec("delete from test_cases;").Error
}

func (_ *TestCaseLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
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

	templateFuncMaps := extension.GetTemplateFuncMaps()

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

var TestCommandLogicInstance = &TestCommandLogic{}
var TestPatternLogicInstance = &TestPatternLogic{}
var TestCaseLogicInstance = &TestCaseLogic{}

func init() {
	extension.RegisterDesignAccessor(TestCommandLogicInstance)
	extension.RegisterDesignAccessor(TestPatternLogicInstance)
	extension.RegisterDesignAccessor(TestCaseLogicInstance)
}
