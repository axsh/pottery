package models

import (
	"github.com/qb0C80aE/clay/extensions"
)

type AutomationCommand struct {
	ID      int    `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ZipFile []byte `json:"zip_file" form:"zip_file"`
}

func NewAutomationCommandModel() *AutomationCommand {
	return &AutomationCommand{}
}

var sharedAutomationCommandModel = NewAutomationCommandModel()

func SharedAutomationCommandModel() *AutomationCommand {
	return sharedAutomationCommandModel
}

func init() {
	extensions.RegisterModel(sharedAutomationCommandModel)
}
