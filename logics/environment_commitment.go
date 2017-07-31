package logics

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type environmentCommitmentLogic struct {
	*clayLogics.BaseLogic
}

func newEnvironmentCommitmentLogic() *environmentCommitmentLogic {
	logic := &environmentCommitmentLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func initGitRepository(environment *models.Environment) error {
	cmd := exec.Command("mkdir", "-p", environment.GitRepositoryURI)
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("git", "status")
	cmd.Dir = environment.GitRepositoryURI
	if err := cmd.Run(); err != nil {
		cmd := exec.Command("git", "init")
		cmd.Dir = environment.GitRepositoryURI
		cmd.Run()
	}

	cmd = exec.Command("git", "config", "--local", "user.name", environment.GitUserName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("git", "config", "--local", "user.email", environment.GitUserEmail)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	return nil
}

func updateDesignFile(db *gorm.DB, environment *models.Environment) error {
	design, err := clayLogics.UniqueDesignLogic().GetSingle(db, "", nil, "*")

	if err != nil {
		return err
	}

	jsonString, err := json.MarshalIndent(design, "", "    ")
	if ioutil.WriteFile(fmt.Sprintf("%s/%s", environment.GitRepositoryURI, environment.DesignFileName), jsonString, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func updateTemplateFile(db *gorm.DB, environment *models.Environment) error {
	templateID := fmt.Sprintf("%d", environment.TemplateID)
	template, err := clayLogics.GenerateTemplate(db, templateID, nil)

	if err != nil {
		return err
	}

	if ioutil.WriteFile(fmt.Sprintf("%s/%s", environment.GitRepositoryURI, environment.TemplateFileName), ([]byte)(template.(string)), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func updateTestCaseFile(db *gorm.DB, environment *models.Environment) error {
	testRunnerScriptID := fmt.Sprintf("%d", environment.TestRunnerScriptID)
	testRunnerScript, err := clayLogics.GenerateTemplate(db, testRunnerScriptID, nil)

	if err != nil {
		return err
	}

	requirements := []*models.Requirement{}

	if err := db.Preload("Service").
		Preload("SourcePort").
		Preload("DestinationPort").
		Select("*").Find(&requirements).Error; err != nil {
		return err
	}

	cmd := exec.Command("rm", "-rf", environment.TestCaseDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("mkdir", environment.TestCaseDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	if ioutil.WriteFile(fmt.Sprintf("%s/test_runner.sh",
		environment.GitRepositoryURI),
		([]byte)(testRunnerScript.(string)),
		os.ModePerm); err != nil {
		return err
	}

	for _, requirement := range requirements {
		testClientScript, testPattern, err := GenerateTestClientScript(db, requirement.ID)
		if err != nil {
			return err
		}
		testServerScript, testPattern, err := GenerateTestServerScript(db, requirement.ID)
		if err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_client.sh",
			environment.GitRepositoryURI,
			environment.TestCaseDirectoryName,
			testPattern),
			([]byte)(testClientScript.(string)),
			os.ModePerm); err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_server.sh",
			environment.GitRepositoryURI,
			environment.TestCaseDirectoryName,
			testPattern),
			([]byte)(testServerScript.(string)),
			os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func updateServerConfigFiles(db *gorm.DB, environment *models.Environment) error {
	cmd := exec.Command("rm", "-rf", environment.ServerConfigDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("mkdir", environment.ServerConfigDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	nodeExtraAttributes := []*loamModels.NodeExtraAttribute{}
	if err := db.Preload("Node").
		Preload("ValueNodeExtraAttributeOption").
		Select("*").Find(&nodeExtraAttributes, "node_extra_attribute_field_id = ?", 3).Error; err != nil {
		return err
	}

	for _, nodeExtraAttribute := range nodeExtraAttributes {
		templateID := nodeExtraAttribute.ValueNodeExtraAttributeOption.ValueInt.Int64
		template, err := clayLogics.GenerateTemplate(db, strconv.Itoa((int)(templateID)), nil)
		if err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_initialize.txt",
			environment.GitRepositoryURI,
			environment.ServerConfigDirectoryName,
			nodeExtraAttribute.Node.Name),
			([]byte)(template.(string)),
			os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func updateDeviceConfigFiles(db *gorm.DB, environment *models.Environment) error {
	cmd := exec.Command("rm", "-rf", environment.DeviceConfigDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	cmd = exec.Command("mkdir", environment.DeviceConfigDirectoryName)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}

	nodeExtraAttributes := []*loamModels.NodeExtraAttribute{}
	if err := db.Preload("Node").
		Select("*").Find(&nodeExtraAttributes, "node_extra_attribute_field_id = ?", 5).Error; err != nil {
		return err
	}

	for _, nodeExtraAttribute := range nodeExtraAttributes {
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_initialize.txt",
			environment.GitRepositoryURI,
			environment.DeviceConfigDirectoryName,
			nodeExtraAttribute.Node.Name),
			([]byte)(nodeExtraAttribute.ValueString.String),
			os.ModePerm); err != nil {
			return err
		}
	}

	if err := db.Preload("Node").
		Select("*").Find(&nodeExtraAttributes, "node_extra_attribute_field_id = ?", 6).Error; err != nil {
		return err
	}

	for _, nodeExtraAttribute := range nodeExtraAttributes {
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s_config.txt",
			environment.GitRepositoryURI,
			environment.DeviceConfigDirectoryName,
			nodeExtraAttribute.Node.Name),
			([]byte)(nodeExtraAttribute.ValueString.String),
			os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

/*
func updateNodeConfigFiles(_ *gorm.DB, environment *models.Environment) error {
	for _, config := range environment.NodeConfigs {
		cmd := exec.Command("mkdir", "-p", fmt.Sprintf("%s/%s/%s", environment.GitRepositoryURI, "config", config.Node.Name))
		cmdMessage, err := cmd.CombinedOutput()
		if err != nil {
			return errors.New(string(cmdMessage))
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s/initialize.txt",
			environment.GitRepositoryURI,
			"config",
			config.Node.Name),
			([]byte)(config.InitializeConfig),
			os.ModePerm); err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s/config.txt",
			environment.GitRepositoryURI,
			"config",
			config.Node.Name),
			([]byte)(config.Config),
			os.ModePerm); err != nil {
			return err
		}
		if ioutil.WriteFile(fmt.Sprintf("%s/%s/%s/firmware_version.txt",
			environment.GitRepositoryURI,
			"config",
			config.Node.Name),
			([]byte)(config.FirmwareVersion),
			os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
*/

func commit(environment *models.Environment, message string) error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}
	cmd = exec.Command("git", "commit", "-m", message)
	cmd.Dir = environment.GitRepositoryURI
	cmdMessage, err = cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(cmdMessage))
	}
	return nil
}

func (logic *environmentCommitmentLogic) Update(db *gorm.DB, _ string, _ url.Values, _ interface{}) (interface{}, error) {
	environment := &models.Environment{}

	if err := db.Select("*").First(environment, "1").Error; err != nil {
		return nil, err
	}

	if err := initGitRepository(environment); err != nil {
		return "", err
	}

	message := fmt.Sprintf("Automatic commit at %s", time.Now().String())
	if err := updateDesignFile(db, environment); err != nil {
		return "", err
	}
	if err := updateTemplateFile(db, environment); err != nil {
		return "", err
	}
	if err := updateTestCaseFile(db, environment); err != nil {
		return "", err
	}
	if err := updateServerConfigFiles(db, environment); err != nil {
		return "", err
	}
	if err := updateDeviceConfigFiles(db, environment); err != nil {
		return "", err
	}
	if err := commit(environment, message); err != nil {
		return "", err
	}
	return "{}", nil
}

var uniqueEnvironmentCommitmentLogic = newEnvironmentCommitmentLogic()

func UniqueEnvironmentCommitmentLogic() extensions.Logic {
	return uniqueEnvironmentCommitmentLogic
}

func init() {
}
