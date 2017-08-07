package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayModels "github.com/qb0C80aE/clay/models"
	loamModels "github.com/qb0C80aE/loam/models"
)

type Protocol struct {
	ID   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"not null;unique"`
}

type Service struct {
	ID                  int                  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name                string               `json:"name" gorm:"not null;unique"`
	Connections         []*Connection        `json:"connections"`
	FirewallTestProgram *FirewallTestProgram `json:"firewall_test_program"`
}

type Connection struct {
	ID         int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ServiceID  int       `json:"service_id" gorm:"not null" sql:"type:integer references services(id) on delete cascade"`
	ProtocolID int       `json:"protocol_id" gorm:"not null" sql:"type:integer references protocols(id) on delete cascade"`
	Protocol   *Protocol `json:"protocol"`
	PortNumber int       `json:"port_number" gorm:"not null"`
}

type Requirement struct {
	ID                int              `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	SourcePortID      sql.NullInt64    `json:"source_port_id" sql:"type:integer references ports(id) on delete cascade"`
	SourcePort        *loamModels.Port `json:"source_port"`
	DestinationPortID sql.NullInt64    `json:"destination_port_id" sql:"type:integer references ports(id) on delete cascade"`
	DestinationPort   *loamModels.Port `json:"destination_port"`
	ServiceID         int              `json:"service_id" gorm:"not null" sql:"type:integer references services(id) on delete cascade"`
	Service           *Service         `json:"service"`
	Access            bool             `json:"access"`
}

type FirewallTestProgram struct {
	ID                     int                  `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name                   string               `json:"name" gorm:"not null;unique"`
	ServiceID              int                  `json:"service_id" gorm:"not null;unique" sql:"type:integer references services(id) on delete cascade"`
	Service                *Service             `json:"service"`
	ServerScriptTemplateID int                  `json:"server_script_template_id" gorm:"not null;unique" sql:"type:integer references templates(id) on delete cascade"`
	ServerScriptTemplate   *clayModels.Template `json:"server_script_template"`
	ClientScriptTemplateID int                  `json:"client_script_template_id" gorm:"not null;unique" sql:"type:integer references templates(id) on delete cascade"`
	ClientScriptTemplate   *clayModels.Template `json:"client_script_template"`
}

func NewProtocolModel() *Protocol {
	return &Protocol{}
}

func NewServiceModel() *Service {
	return &Service{}
}

func NewConnectionModel() *Connection {
	return &Connection{}
}

func NewRequirementModel() *Requirement {
	return &Requirement{}
}

func NewFirewallTestProgramModel() *FirewallTestProgram {
	return &FirewallTestProgram{}
}

var sharedProtocolModel = NewProtocolModel()
var sharedServiceModel = NewServiceModel()
var sharedConnectionModel = NewConnectionModel()
var sharedRequirementModel = NewRequirementModel()
var sharedFirewallTestProgramModel = NewFirewallTestProgramModel()

func SharedProtocolModel() *Protocol {
	return sharedProtocolModel
}

func SharedServiceModel() *Service {
	return sharedServiceModel
}

func SharedConnectionModel() *Connection {
	return sharedConnectionModel
}

func SharedRequirementModel() *Requirement {
	return sharedRequirementModel
}

func SharedFirewallTestProgramModel() *FirewallTestProgram {
	return sharedFirewallTestProgramModel
}

func (testProgram *FirewallTestProgram) SetupInitialData(db *gorm.DB) error {
	db.Exec(`
		create trigger if not exists DeleteTestProgramTemplate delete on firewall_test_programs
		begin
			delete from templates where id = old.client_script_template_id;
			delete from templates where id = old.server_script_template_id;
		end;
	`)

	return nil
}

func init() {
	extensions.RegisterInitialDataLoader(sharedFirewallTestProgramModel)
	extensions.RegisterModel(sharedProtocolModel)
	extensions.RegisterModel(sharedServiceModel)
	extensions.RegisterModel(sharedConnectionModel)
	extensions.RegisterModel(sharedRequirementModel)
	extensions.RegisterModel(sharedFirewallTestProgramModel)
}
