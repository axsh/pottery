package models

import (
	"database/sql"
	"github.com/qb0C80aE/clay/extension"
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
	Accessibility     bool             `json:"accessibility"`
}

var ProtocolModel = &Protocol{}
var ServiceModel = &Service{}
var ConnectionModel = &Connection{}
var RequirementModel = &Requirement{}

func init() {
	extension.RegisterModelType(Protocol{})
	extension.RegisterModelType(Service{})
	extension.RegisterModelType(Connection{})
	extension.RegisterModelType(Requirement{})
}
