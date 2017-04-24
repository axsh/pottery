package models

import (
	"database/sql"
	"github.com/qb0C80aE/clay/extensions"
	loamModels "github.com/qb0C80aE/loam/models"
)

type Protocol struct {
	ID   int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"not null;unique"`
}

type Service struct {
	ID          int           `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name        string        `json:"name" gorm:"not null;unique"`
	Connections []*Connection `json:"connections"`
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

var sharedProtocolModel = NewProtocolModel()
var sharedServiceModel = NewServiceModel()
var sharedConnectionModel = NewConnectionModel()
var sharedRequirementModel = NewRequirementModel()

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

func init() {
	extensions.RegisterModel("Protocol", sharedProtocolModel)
	extensions.RegisterModel("Service", sharedServiceModel)
	extensions.RegisterModel("Connection", sharedConnectionModel)
	extensions.RegisterModel("Requirement", sharedRequirementModel)
}
