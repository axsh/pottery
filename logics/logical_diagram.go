package logics

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	loamLogics "github.com/qb0C80aE/loam/logics"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
)

var virtualNodeIconPaths = map[int]string{
	1: fmt.Sprintf("%s/%s", diagramImageRoot, "l2switch_v.png"),
	2: fmt.Sprintf("%s/%s", diagramImageRoot, "l3switch_v.png"),
	3: fmt.Sprintf("%s/%s", diagramImageRoot, "firewall_v.png"),
	4: fmt.Sprintf("%s/%s", diagramImageRoot, "router_v.png"),
	5: fmt.Sprintf("%s/%s", diagramImageRoot, "loadbalancer_v.png"),
	6: fmt.Sprintf("%s/%s", diagramImageRoot, "server_v.png"),
	7: fmt.Sprintf("%s/%s", diagramImageRoot, "network.png"),
}

var segmentIconPath = fmt.Sprintf("%s/%s", diagramImageRoot, "segment.png")

type logicalDiagramLogic struct {
	*clayLogics.BaseLogic
}

func newLogicalDiagramLogic() *logicalDiagramLogic {
	logic := &logicalDiagramLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func (logic *logicalDiagramLogic) GetSingle(db *gorm.DB, id string, queryFields string) (interface{}, error) {
	nodePvs := []*loamModels.NodePv{}
	if err := db.Select(queryFields).Find(&nodePvs).Error; err != nil {
		return nil, err
	}

	nodeTypes := []*loamModels.NodeType{}
	if err := db.Select(queryFields).Find(&nodeTypes).Error; err != nil {
		return nil, err
	}

	nodes := []*loamModels.Node{}
	if err := db.Preload("Ports").Select(queryFields).Find(&nodes).Error; err != nil {
		return nil, err
	}

	ports := []*loamModels.Port{}
	if err := db.Select(queryFields).Find(&ports).Error; err != nil {
		return nil, err
	}

	nodeMap := make(map[int]*loamModels.Node)
	portMap := make(map[int]*loamModels.Port)
	consumedPortMap := make(map[int]*loamModels.Port)

	for _, node := range nodes {
		nodeMap[node.ID] = node
	}
	for _, port := range ports {
		portMap[port.ID] = port
	}

	segments := loamLogics.GenerateSegments(nodeMap, portMap, consumedPortMap)

	diagram := &models.Diagram{}

	for _, node := range nodes {
		if node.NodeTypeID != 1 {
			var iconPathMap map[int]string
			if node.NodePvID == 1 {
				iconPathMap = physicalNodeIconPaths
			} else {
				iconPathMap = virtualNodeIconPaths
			}
			diagramNode := &models.DiagramNode{
				node.Name,
				iconPathMap[node.NodeTypeID],
			}
			diagram.Nodes = append(diagram.Nodes, diagramNode)
		}
	}

	for i, segment := range segments {

		diagramNode := &models.DiagramNode{
			fmt.Sprintf("[%d]%s", i, segment.Cidr),
			segmentIconPath,
		}
		diagram.Nodes = append(diagram.Nodes, diagramNode)

		for _, port := range segment.Ports {
			diagramInterface := &models.DiagramInterface{
				Source: "",
				Target: fmt.Sprintf("%s[%s](%s/%d)",
					port.Name,
					port.MacAddress.String,
					port.Ipv4Address.String,
					port.Ipv4Prefix.Int64,
				),
			}
			diagramMeta := &models.DiagramMeta{
				Interface: diagramInterface,
			}
			diagramLink := &models.DiagramLink{
				Source: fmt.Sprintf("[%d]%s", i, segment.Cidr),
				Target: nodeMap[port.NodeID].Name,
				Meta:   diagramMeta,
			}
			diagram.Links = append(diagram.Links, diagramLink)
		}

	}

	return diagram, nil
}

var uniqueLogicalDiagramLogic = newLogicalDiagramLogic()

// UniqueLogicalDiagramLogic returns the unique logical diagram logic instance
func UniqueLogicalDiagramLogic() extensions.Logic {
	return uniqueLogicalDiagramLogic
}

func init() {
}