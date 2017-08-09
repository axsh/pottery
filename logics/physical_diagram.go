package logics

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/qb0C80aE/clay/extensions"
	clayLogics "github.com/qb0C80aE/clay/logics"
	clayModels "github.com/qb0C80aE/clay/models"
	"github.com/qb0C80aE/clay/utils/mapstruct"
	loamLogics "github.com/qb0C80aE/loam/logics"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
	"net/url"
	"strconv"
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

type physicalDiagramNodeLogic struct {
	*clayLogics.BaseLogic
}

func newPhysicalDiagramLogic() *physicalDiagramLogic {
	logic := &physicalDiagramLogic{
		BaseLogic: &clayLogics.BaseLogic{},
	}
	return logic
}

func newPhysicalDiagramNodeLogic() *physicalDiagramNodeLogic {
	logic := &physicalDiagramNodeLogic{
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

func (logic *physicalDiagramNodeLogic) GetSingle(db *gorm.DB, id string, _ url.Values, queryFields string) (interface{}, error) {

	diagramNode := &models.DiagramNode{}

	if err := db.Select(queryFields).First(diagramNode, id).Error; err != nil {
		return nil, err
	}

	return diagramNode, nil
}

func (logic *physicalDiagramNodeLogic) GetMulti(db *gorm.DB, _ url.Values, queryFields string) (interface{}, error) {
	diagramNodes := []*models.DiagramNode{}

	if err := db.Select(queryFields).Find(&diagramNodes).Error; err != nil {
		return nil, err
	}

	return diagramNodes, nil
}

func (logic *physicalDiagramNodeLogic) Update(db *gorm.DB, id string, _ url.Values, data interface{}) (interface{}, error) {

	diagramNode := data.(*models.DiagramNode)
	diagramNode.ID, _ = strconv.Atoi(id)

	if err := db.Save(&diagramNode).Error; err != nil {
		return nil, err
	}

	return diagramNode, nil

}

func (logic *physicalDiagramNodeLogic) Delete(db *gorm.DB, id string, _ url.Values) error {

	diagramNode := &models.DiagramNode{}

	if err := db.First(&diagramNode, id).Error; err != nil {
		return err
	}

	if err := db.Delete(&diagramNode).Error; err != nil {
		return err
	}

	return nil

}

func (logic *physicalDiagramNodeLogic) ExtractFromDesign(db *gorm.DB) (string, interface{}, error) {
	diagramNodes := []*models.DiagramNode{}
	if err := db.Select("*").Find(&diagramNodes).Error; err != nil {
		return "", nil, err
	}
	return extensions.RegisteredResourceName(models.SharedDiagramNodeModel()), diagramNodes, nil
}

func (logic *physicalDiagramNodeLogic) DeleteFromDesign(db *gorm.DB) error {
	return db.Delete(models.SharedDiagramNodeModel()).Error
}

func (logic *physicalDiagramNodeLogic) LoadToDesign(db *gorm.DB, data interface{}) error {
	container := []*models.DiagramNode{}
	design := data.(*clayModels.Design)
	if value, exists := design.Content[extensions.RegisteredResourceName(models.SharedDiagramNodeModel())]; exists {
		if err := mapstruct.MapToStruct(value.([]interface{}), &container); err != nil {
			return err
		}
		for _, templatePersistentParameter := range container {
			if err := db.Create(templatePersistentParameter).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

var uniquePhysicalDiagramLogic = newPhysicalDiagramLogic()
var uniquePhysicalDiagramNodeLogic = newPhysicalDiagramNodeLogic()

// UniquePhysicalDiagramLogic returns the unique physical diagram logic instance
func UniquePhysicalDiagramLogic() extensions.Logic {
	return uniquePhysicalDiagramLogic
}

// UniquePhysicalDiagramNodeLogic returns the unique physical diagram node logic instance
func UniquePhysicalDiagramNodeLogic() extensions.Logic {
	return uniquePhysicalDiagramNodeLogic
}

func init() {
	extensions.RegisterDesignAccessor(uniquePhysicalDiagramNodeLogic)
}
