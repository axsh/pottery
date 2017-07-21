// +build integration

package integration

import (
	"database/sql"
	. "github.com/qb0C80aE/clay/integration"
	loamModels "github.com/qb0C80aE/loam/models"
	"github.com/qb0C80aE/pottery/models"
	"net/http"
	"os"
	"testing"
)

func beforeSetupServer() string {
	currentDir, _ := os.Getwd()
	os.Chdir("../")
	return currentDir
}

func afterSetupServer(savedDir string) {
	os.Chdir(savedDir)
}

func TestGetDiagram_Empty(t *testing.T) {
	currentDir := beforeSetupServer()
	server := SetupServer()
	afterSetupServer(currentDir)
	defer server.Close()

	responseText, code := Execute(t, http.MethodGet, GenerateSingleResourceURL(server, "diagrams", "physical", nil), nil)
	CheckResponseJSON(t, code, http.StatusOK, responseText, LoadExpectation(t, "diagram/TestGetDiagram_Empty_1.json"), &models.Diagram{})

	responseText, code = Execute(t, http.MethodGet, GenerateSingleResourceURL(server, "diagrams", "logical", nil), nil)
	CheckResponseJSON(t, code, http.StatusOK, responseText, LoadExpectation(t, "diagram/TestGetDiagram_Empty_2.json"), &models.Diagram{})
}

func TestGetDiagram(t *testing.T) {
	currentDir := beforeSetupServer()
	server := SetupServer()
	afterSetupServer(currentDir)
	defer server.Close()

	nodeType1 := &loamModels.NodeType{
		ID:   1,
		Name: "L2Switch",
	}
	nodeType2 := &loamModels.NodeType{
		ID:   2,
		Name: "L3Switch",
	}
	nodeType3 := &loamModels.NodeType{
		ID:   3,
		Name: "Firewall",
	}
	nodeType4 := &loamModels.NodeType{
		ID:   4,
		Name: "Router",
	}
	nodeType5 := &loamModels.NodeType{
		ID:   5,
		Name: "LoadBalancer",
	}
	nodeType6 := &loamModels.NodeType{
		ID:   6,
		Name: "Server",
	}
	nodeType7 := &loamModels.NodeType{
		ID:   7,
		Name: "Network",
	}
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_types", nil), nodeType1)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_types", nil), nodeType2)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_types", nil), nodeType3)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_types", nil), nodeType4)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_types", nil), nodeType5)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_types", nil), nodeType6)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_types", nil), nodeType7)

	nodeExtraAttributeField1 := &loamModels.NodeExtraAttributeField{
		ID:   1,
		Name: "virtual",
	}
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "node_extra_attribute_fields", nil), nodeExtraAttributeField1)

	router1 := &loamModels.Node{
		ID:         1,
		Name:       "router1",
		NodeTypeID: 4,
		NodeExtraAttributes: []*loamModels.NodeExtraAttribute{
			{
				NodeExtraAttributeFieldID: nodeExtraAttributeField1.ID,
				ValueBool: sql.NullBool{
					Valid: true,
					Bool:  false,
				},
			},
		},
	}
	firewall1 := &loamModels.Node{
		ID:         2,
		Name:       "firewall1",
		NodeTypeID: 3,
		NodeExtraAttributes: []*loamModels.NodeExtraAttribute{
			{
				NodeExtraAttributeFieldID: nodeExtraAttributeField1.ID,
				ValueBool: sql.NullBool{
					Valid: true,
					Bool:  false,
				},
			},
		},
	}
	l2sw1 := &loamModels.Node{
		ID:         3,
		Name:       "l2sw1",
		NodeTypeID: 1,
		NodeExtraAttributes: []*loamModels.NodeExtraAttribute{
			{
				NodeExtraAttributeFieldID: nodeExtraAttributeField1.ID,
				ValueBool: sql.NullBool{
					Valid: true,
					Bool:  true,
				},
			},
		},
	}
	server1 := &loamModels.Node{
		ID:         4,
		Name:       "server1",
		NodeTypeID: 6,
		NodeExtraAttributes: []*loamModels.NodeExtraAttribute{
			{
				NodeExtraAttributeFieldID: nodeExtraAttributeField1.ID,
				ValueBool: sql.NullBool{
					Valid: true,
					Bool:  true,
				},
			},
		},
	}
	server2 := &loamModels.Node{
		ID:         5,
		Name:       "server2",
		NodeTypeID: 6,
		NodeExtraAttributes: []*loamModels.NodeExtraAttribute{
			{
				NodeExtraAttributeFieldID: nodeExtraAttributeField1.ID,
				ValueBool: sql.NullBool{
					Valid: true,
					Bool:  true,
				},
			},
		},
	}

	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "nodes", nil), router1)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "nodes", nil), firewall1)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "nodes", nil), l2sw1)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "nodes", nil), server1)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "nodes", nil), server2)

	router1port0 := &loamModels.Port{
		ID:     1,
		NodeID: router1.ID,
		Layer:  3,
		Number: 1,
		Name:   "router1port0",
		MacAddress: sql.NullString{
			String: "00:00:00:00:00:01",
			Valid:  true,
		},
		Ipv4Address: sql.NullString{
			String: "192.168.0.1",
			Valid:  true,
		},
		Ipv4Prefix: sql.NullInt64{
			Int64: 24,
			Valid: true,
		},
	}
	firewll1port0 := &loamModels.Port{
		ID:     2,
		NodeID: firewall1.ID,
		Layer:  3,
		Number: 1,
		Name:   "firewll1port0",
		DestinationPortID: sql.NullInt64{
			Int64: int64(router1port0.ID),
			Valid: true,
		},
		MacAddress: sql.NullString{
			String: "00:00:00:00:01:01",
			Valid:  true,
		},
		Ipv4Address: sql.NullString{
			String: "192.168.0.2",
			Valid:  true,
		},
		Ipv4Prefix: sql.NullInt64{
			Int64: 24,
			Valid: true,
		},
	}
	firewll1port1 := &loamModels.Port{
		ID:     3,
		NodeID: firewall1.ID,
		Layer:  3,
		Number: 2,
		Name:   "firewll1port1",
		MacAddress: sql.NullString{
			String: "00:00:00:00:01:02",
			Valid:  true,
		},
		Ipv4Address: sql.NullString{
			String: "192.168.1.1",
			Valid:  true,
		},
		Ipv4Prefix: sql.NullInt64{
			Int64: 24,
			Valid: true,
		},
	}
	l2sw1port0 := &loamModels.Port{
		ID:     4,
		NodeID: l2sw1.ID,
		Layer:  1,
		Number: 1,
		Name:   "l2sw1port0",
		DestinationPortID: sql.NullInt64{
			Int64: int64(firewll1port1.ID),
			Valid: true,
		},
	}
	l2sw1port1 := &loamModels.Port{
		ID:     5,
		NodeID: l2sw1.ID,
		Layer:  1,
		Number: 2,
		Name:   "l2sw1port1",
	}
	l2sw1port2 := &loamModels.Port{
		ID:     6,
		NodeID: l2sw1.ID,
		Layer:  1,
		Number: 3,
		Name:   "l2sw1port2",
	}
	server1port0 := &loamModels.Port{
		ID:     7,
		NodeID: server1.ID,
		Layer:  3,
		Number: 1,
		Name:   "server1port0",
		DestinationPortID: sql.NullInt64{
			Int64: int64(l2sw1port1.ID),
			Valid: true,
		},
		MacAddress: sql.NullString{
			String: "00:00:00:00:02:01",
			Valid:  true,
		},
		Ipv4Address: sql.NullString{
			String: "192.168.1.2",
			Valid:  true,
		},
		Ipv4Prefix: sql.NullInt64{
			Int64: 24,
			Valid: true,
		},
	}
	server2port0 := &loamModels.Port{
		ID:     8,
		NodeID: server2.ID,
		Layer:  3,
		Number: 1,
		Name:   "server2port0",
		DestinationPortID: sql.NullInt64{
			Int64: int64(l2sw1port2.ID),
			Valid: true,
		},
		MacAddress: sql.NullString{
			String: "00:00:00:00:03:01",
			Valid:  true,
		},
		Ipv4Address: sql.NullString{
			String: "192.168.1.3",
			Valid:  true,
		},
		Ipv4Prefix: sql.NullInt64{
			Int64: 24,
			Valid: true,
		},
	}

	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), router1port0)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), firewll1port0)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), firewll1port1)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), l2sw1port0)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), l2sw1port1)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), l2sw1port2)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), server1port0)
	Execute(t, http.MethodPost, GenerateMultiResourceURL(server, "ports", nil), server2port0)

	responseText, code := Execute(t, http.MethodGet, GenerateSingleResourceURL(server, "diagrams", "physical", nil), nil)
	CheckResponseJSON(t, code, http.StatusOK, responseText, LoadExpectation(t, "diagram/TestGetDiagram_1.json"), &models.Diagram{})

	responseText, code = Execute(t, http.MethodGet, GenerateSingleResourceURL(server, "diagrams", "logical", nil), nil)
	CheckResponseJSON(t, code, http.StatusOK, responseText, LoadExpectation(t, "diagram/TestGetDiagram_2.json"), &models.Diagram{})
}
