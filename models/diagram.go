package models

import "github.com/qb0C80aE/clay/extensions"

// Diagram is the model class what represents physical and logical diagrams
type Diagram struct {
	ID    int            `json:"-,omitempty" gorm:"primary_key"`
	Nodes []*DiagramNode `json:"nodes"`
	Links []*DiagramLink `json:"links"`
}

// DiagramNode is the model class what represents nodes in diagrams
type DiagramNode struct {
	Name string           `json:"name"`
	Icon string           `json:"icon"`
	Meta *DiagramNodeMeta `json:"meta"`
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

func init() {
	extensions.RegisterModel(&Diagram{})
}
