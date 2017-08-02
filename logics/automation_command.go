package logics

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/pottery/models"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
)

type automationCommandLogic struct {
	*clayLogics.BaseLogic
}

func newAutomationCommandLogic() *automationCommandLogic {
	logic := &automationCommandLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func (logic *automationCommandLogic) Update(db *gorm.DB, _ string, urlValues url.Values, data interface{}) (interface{}, error) {

	automationCommand := data.(*models.AutomationCommand)
	fileMap := map[string][]byte{}

	r, err := zip.NewReader(bytes.NewReader(automationCommand.ZipFile), int64(len(automationCommand.ZipFile)))
	if err != nil {
		return "", err
	}
	for _, zf := range r.File {
		src, err := zf.Open()
		if err != nil {
			return "", err
		}
		defer src.Close()
		fileContent := &bytes.Buffer{}
		io.Copy(fileContent, src)
		fileMap[zf.Name] = fileContent.Bytes()
	}

	if _, exists := fileMap["design.json"]; !exists {
		return "", errors.New("design.json does not exist in the uploaded zip file")
	}
	if _, exists := fileMap["automation.json"]; !exists {
		return "", errors.New("automation.json does not exist in the uploaded zip file")
	}

	designJson := fileMap["design.json"]
	design := &clayModels.Design{}
	if err := json.Unmarshal(designJson, &design); err != nil {
		return "", err
	}

	if _, err := clayLogics.UniqueDesignLogic().Update(db, "", nil, design); err != nil {
		return "", err
	}

	environment, err := UpdateEnvironmentFiles(db, urlValues)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("rm", "-rf", "uploaded")
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(cmdMessage))
	}

	cmd = exec.Command("mkdir", "uploaded")
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(cmdMessage))
	}

	for key, value := range fileMap {
		if ioutil.WriteFile(fmt.Sprintf("%s/uploaded/%s",
			environment.GitRepositoryURI,
			key),
			value,
			os.ModePerm); err != nil {
			return "", err
		}
	}

	if err := CommitEnvironmentFiles(environment); err != nil {
		return "", err
	}
	return "{}", nil
}

var uniqueAutomationCommandLogic = newAutomationCommandLogic()

func UniqueAutomationCommandLogic() extensions.Logic {
	return uniqueAutomationCommandLogic
}

func init() {
}
