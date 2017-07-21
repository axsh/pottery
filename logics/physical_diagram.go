package logics

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	loamLogics "github.com/qb0C80aE/loam/logics"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
	"net/url"
)

var physicalNodeIconPaths = map[int]string{
	1: fmt.Sprintf("%s/%s", diagramImageRoot, "l2switch.png"),
	2: fmt.Sprintf("%s/%s", diagramImageRoot, "l3switch.png"),
	3: fmt.Sprintf("%s/%s", diagramImageRoot, "firewall.png"),
	4: fmt.Sprintf("%s/%s", diagramImageRoot, "router.png"),
	5: fmt.Sprintf("%s/%s", diagramImageRoot, "loadbalancer.png"),
	6: fmt.Sprintf("%s/%s", diagramImageRoot, "server.png"),
	7: fmt.Sprintf("%s/%s", diagramImageRoot, "network.png"),
}

type physicalDiagramLogic struct {
	*clayLogics.BaseLogic
}

func newPhysicalDiagramLogic() *physicalDiagramLogic {
	logic := &physicalDiagramLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func (logic *physicalDiagramLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {
	diagram := &models.Diagram{}

	nodes := []*loamModels.Node{}
	if err := db.Preload("NodeExtraAttributes").
		Preload("NodeExtraAttributes.NodeExtraAttributeField").
		Preload("NodeExtraAttributes.ValueNodeExtraAttributeOption").
		Preload("Ports").
		Select(queryFields).Find(&nodes).Error; err != nil {
		return nil, err
	}

	nodeMap := make(map[int]*loamModels.Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}

	ports := []*loamModels.Port{}
	if err := db.Select(queryFields).Find(&ports).Error; err != nil {
		return nil, err
	}

	portMap := make(map[int]*loamModels.Port)
	for _, port := range ports {
		portMap[port.ID] = port
	}

	for _, node := range nodes {
		nodeExtraAttributesMap, err := loamLogics.CreateNodeAttributeMap(node)
		if err != nil {
			return nil, err
		}
		var iconPathMap map[int]string
		attributes, exists := nodeExtraAttributesMap["virtual"]
		attribute := attributes[0]
		if exists && attribute.ValueBool.Valid && attribute.ValueBool.Bool {
			iconPathMap = virtualNodeIconPaths
		} else {
			iconPathMap = physicalNodeIconPaths
		}
		diagramNodeMeta := &models.DiagramNodeMeta{
			NodeID: node.ID,
		}
		diagramNode := &models.DiagramNode{
			Name: node.Name,
			Icon: iconPathMap[node.NodeTypeID],
			Meta: diagramNodeMeta,
		}
		diagram.Nodes = append(diagram.Nodes, diagramNode)
	}

	registerdPortMap := make(map[int]int)
	for _, port := range ports {
		_, exists := registerdPortMap[int(port.DestinationPortID.Int64)]
		if (port.DestinationPortID.Valid) && (!exists) {
			sourceNode := nodeMap[port.NodeID]
			destinationPort := portMap[int(port.DestinationPortID.Int64)]
			destinationNode := nodeMap[destinationPort.NodeID]

			diagramInterface := &models.DiagramInterface{
				Source: port.Name,
				Target: destinationPort.Name,
			}
			diagramMeta := &models.DiagramMeta{
				Interface: diagramInterface,
			}
			diagramLink := &models.DiagramLink{
				Source: sourceNode.Name,
				Target: destinationNode.Name,
				Meta:   diagramMeta,
			}

			diagram.Links = append(diagram.Links, diagramLink)

			registerdPortMap[port.ID] = port.ID
		}
	}

	return diagram, nil
}

var uniquePhysicalDiagramLogic = newPhysicalDiagramLogic()

// UniquePhysicalDiagramLogic returns the unique physical diagram logic instance
func UniquePhysicalDiagramLogic() extensions.Logic {
	return uniquePhysicalDiagramLogic
}

func init() {
}
