package logics

import (
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
	"bytes"
)

type testProgramLogic struct {
	*clayLogics.BaseLogic
}

type testPatternLogic struct {
	*clayLogics.BaseLogic
}

type testCaseLogic struct {
	*clayLogics.BaseLogic
}

func newTestProgramLogic() *testProgramLogic {
	logic := &testProgramLogic{
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

func (logic *testProgramLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {

	testProgram := &models.TestProgram{}

	if err := db.Select(queryFields).First(testProgram, id).Error; err != nil {
		return nil, err
	}

	return testProgram, nil

}

func (logic *testProgramLogic) GetMulti(db *gorm.DB, queryFields string) (interface{}, error) {

	testPrograms := []*models.TestProgram{}

	if err := db.Select(queryFields).Find(&testPrograms).Error; err != nil {
		return nil, err
	}

	result := make([]interface{}, len(testPrograms))
	for i, data := range testPrograms {
		result[i] = data
	}

	return result, nil

}

func (logic *testProgramLogic) Create(db *gorm.DB, data interface{}) (interface{}, error) {

	testProgram := data.(*models.TestProgram)

	if err := db.Create(&testProgram).Error; err != nil {
		return nil, err
	}

	return testProgram, nil
}

func (logic *testProgramLogic) Update(db *gorm.DB, id string, data interface{}) (interface{}, error) {

	testProgram := data.(*models.TestProgram)
	testProgram.ID, _ = strconv.Atoi(id)

	if err := db.Save(testProgram).Error; err != nil {
		return nil, err
	}

	return testProgram, nil
}

func (logic *testProgramLogic) Delete(db *gorm.DB, id string) error {

	testProgram := &models.TestProgram{}

	if err := db.First(&testProgram, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&testProgram).Error; err != nil {
		return err
	}

	return nil

}

func (logic *testProgramLogic) Patch(db *gorm.DB, id string) (interface{}, error) {
	testRunnerScript, testPrograms, err := generateTestScripts(db, id)
	if err != nil {
		return "", err
	}
	a := []*models.TestProgram(testPrograms)
	c := bytes.Buffer{}
	c.WriteString("# --- test runner script ---\n")
	c.WriteString(fmt.Sprintf("# %s\n", testRunnerScript))
	c.WriteString("#--------------------------------------\n")
	for _, b := range a {
		c.WriteString("#--------------------------------------\n")
		c.WriteString(fmt.Sprintf("# %s\n", b.Service.Name))
		c.WriteString("# --- server script ---\n")
		c.WriteString(fmt.Sprintf("%s\n", b.ServerScriptTemplate))
		c.WriteString("# --- client script ---\n")
		c.WriteString(fmt.Sprintf("%s\n", b.ClientScriptTemplate))
		c.WriteString("#--------------------------------------\n")
	}
	return c.String(), nil
}

func convertAccess(accessibility bool) string {
	if accessibility {
		return "allow"
	} else {
		return "deny"
	}
}

func generateTestScripts(db *gorm.DB, id string) (string, []*models.TestProgram, error) {

	result := []*models.TestProgram{}

	testCase := &models.TestCase{}
	if err := db.Preload("TestPatterns").
		Preload("TestPatterns.TestProgram").First(&testCase, id).Error; err != nil {
		return "", result, err
	}

	testProgramMap := make(map[string]*models.TestProgram)
	for _, testPattern := range testCase.TestPatterns {
		testProgramMap[testPattern.TestProgram.ServiceName] = testPattern.TestProgram
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

	internetNode := &loamModels.Node{
		ID:   0,
		Name: "Internet",
	}
	internetPort := &loamModels.Port{
		ID:     0,
		Number: 0,
		Layer:  3,
		Name:   "Internet",
		NodeID: internetNode.ID,
		Node:   internetNode,
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
	internetNode.Ports = []*loamModels.Port{internetPort}

	for _, requirement := range requirements {
		if !requirement.SourcePortID.Valid {
			requirement.SourcePort = internetPort
		}
		if !requirement.DestinationPortID.Valid {
			requirement.DestinationPort = internetPort
		}

		testProgram := testProgramMap[requirement.Service.Name]
		script, err := clayLogics.UniqueTemplateLogic().Patch(db, strconv.Itoa(testProgram.ServerScriptTemplateID))
		if err != nil {
			return "", nil, err
		}
		serverScriptTemplate := script.(*clayModels.Template)

		script, err = clayLogics.UniqueTemplateLogic().Patch(db, strconv.Itoa(testProgram.ClientScriptTemplateID))
		if err != nil {
			return "", nil, err
		}
		clientScriptTemplate := script.(*clayModels.Template)

		newTestProgram := &models.TestProgram{
			ServiceName:          fmt.Sprintf("%s_to_%s_%s_%s", requirement.SourcePort.Node.Name, requirement.DestinationPort.Node.Name, requirement.Service.Name, convertAccessibility(requirement.Access)),
			ServerScriptTemplate: serverScriptTemplate,
			ClientScriptTemplate: clientScriptTemplate,
		}
		result = append(result, newTestProgram)
	}

	return testCase.TestRunnerScriptTemplate, result, nil

}

var uniqueTestProgramLogic = newTestProgramLogic()

func UniqueTestProgramLogic() extensions.Logic {
	return uniqueTestProgramLogic
}

func init() {
	extensions.RegisterDesignAccessor(uniqueTestProgramLogic)
	extensions.RegisterTemplateParameterGenerator("TestProgram", uniqueTestProgramLogic)
}
