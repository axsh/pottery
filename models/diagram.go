package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	loamModels "github.com/qb0C80aE/loam/models"
)

// Diagram is the model class what represents physical and logical diagrams
type Diagram struct {
	ID    int            `json:"-,omitempty" gorm:"primary_key"`
	Nodes []*DiagramNode `json:"nodes"`
	Links []*DiagramLink `json:"links"`
}

// DiagramNode is the model class what represents nodes in diagrams
type DiagramNode struct {
	ID   int              `json:"id" gorm:"primary_key" sql:"type:integer references nodes(id) on delete cascade"`
	Name string           `json:"name"`
	Icon string           `json:"icon"`
	Meta *DiagramNodeMeta `json:"meta"`
	X    float64          `json:"x"`
	Y    float64          `json:"y"`
}

// DiagramNodeMeta is the model class that represents the meta information on diagram nodes
type DiagramNodeMeta struct {
	NodeID int `json:"node_id"`
}

// DiagramInterface is the model class what represents interfaces of nodes in diagrams
type DiagramInterface struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// DiagramMeta is the model class what represents meta information in diagrams
type DiagramMeta struct {
	Interface *DiagramInterface `json:"interface"`
}

// DiagramLink is the model class what represents links between nodes in diagrams
type DiagramLink struct {
	Source string       `json:"source"`
	Target string       `json:"target"`
	Meta   *DiagramMeta `json:"meta"`
}

// NewDiagramModel creates a Diagram model instance
func NewDiagramModel() *Diagram {
	return &Diagram{}
}

// NewDiagramNodeModel creates a DiagramNode model instance
func NewDiagramNodeModel() *DiagramNode {
	return &DiagramNode{}
}

// NewDiagramInterfaceModel creates a DiagramInterface model instance
func NewDiagramInterfaceModel() *DiagramInterface {
	return &DiagramInterface{}
}

// NewDiagramMetaModel creates a DiagramMeta model instance
func NewDiagramMetaModel() *DiagramMeta {
	return &DiagramMeta{}
}

// NewDiagramLinkModel creates a DiagramLink model instance
func NewDiagramLinkModel() *DiagramLink {
	return &DiagramLink{}
}

var sharedDiagramModel = NewDiagramModel()
var sharedDiagramNodeModel = NewDiagramNodeModel()
var sharedDiagramInterfaceModel = NewDiagramInterfaceModel()
var sharedDiagramMetaModel = NewDiagramMetaModel()
var sharedDiagramLinkModel = NewDiagramLinkModel()

// SharedDiagramModel returns the diagram model instance used as a model prototype and type analysis
func SharedDiagramModel() *Diagram {
	return sharedDiagramModel
}

// SharedDiagramNodeModel returns the diagram node model instance used as a model prototype and type analysis
func SharedDiagramNodeModel() *DiagramNode {
	return sharedDiagramNodeModel
}

// SharedDiagramInterfaceModel returns the diagram interface model instance used as a model prototype and type analysis
func SharedDiagramInterfaceModel() *DiagramInterface {
	return sharedDiagramInterfaceModel
}

// SharedDiagramMetaModel returns the diagram meta model instance used as a model prototype and type analysis
func SharedDiagramMetaModel() *DiagramMeta {
	return sharedDiagramMetaModel
}

// SharedDiagramLinkModel returns the diagram link model instance used as a model prototype and type analysis
func SharedDiagramLinkModel() *DiagramLink {
	return sharedDiagramLinkModel
}

func (diagram *Diagram) SetupInitialData(db *gorm.DB) error {
	nodeTypes := []*loamModels.NodeType{
		{ID: 1, Name: "L2Switch"},
		{ID: 2, Name: "L3Switch"},
		{ID: 3, Name: "Firewall"},
		{ID: 4, Name: "Router"},
		{ID: 5, Name: "LoadBalancer"},
		{ID: 6, Name: "Server"},
		{ID: 7, Name: "Network"},
	}
	nodeExtraAttributeFields := []*loamModels.NodeExtraAttributeField{
		{ID: 1, Name: "virtual"},
		{ID: 2, Name: "remark"},
	}
	portExtraAttributeFields := []*loamModels.PortExtraAttributeField{
		{ID: 1, Name: "gateway"},
		{ID: 2, Name: "remark"},
	}

	for _, nodeType := range nodeTypes {
		if err := db.Save(nodeType).Error; err != nil {
			return err
		}
	}
	for _, nodeExtraAttributeField := range nodeExtraAttributeFields {
		if err := db.Save(nodeExtraAttributeField).Error; err != nil {
			return err
		}
	}
	for _, portExtraAttributeField := range portExtraAttributeFields {
		if err := db.Save(portExtraAttributeField).Error; err != nil {
			return err
		}
	}
	return nil
}

func init() {
	diagram := &Diagram{}
	extensions.RegisterInitialDataLoader(diagram)
	extensions.RegisterModel(diagram)

	diagramNode := &DiagramNode{}
	extensions.RegisterModel(diagramNode)
}
