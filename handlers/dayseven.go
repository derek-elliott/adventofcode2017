package handlers

// DaySevenNode holds the name, weight, children and parents of a node from Day Seven chalenge
type DaySevenNode struct {
	Name     string          `json:"name"`
	Weight   int             `json:"weight"`
	Children []*DaySevenNode `json:"children"`
	Parent   *DaySevenNode   `json:"parent,omitempty"`
}

// DaySevenInput represents the input for the day seven problem
type DaySevenInput struct {
	Name     string   `json:"name"`
	Weight   int      `json:"weight"`
	Children []string `json:"children"`
}

// DaySevenNodes holds the entire network of day seven nodes
type DaySevenNodes []DaySevenNode

// DaySevenInputs holds the entire input of the day seven problem
type DaySevenInputs []DaySevenInput

// DaySevenPartOne builds the tree for day seven and returns the root element
func DaySevenPartOne(rawData DaySevenInputs) string {
	nodes := make(map[string]*DaySevenNode)
	for _, value := range rawData {
		if node, ok := nodes[value.Name]; ok {
			node.Weight = value.Weight
			for _, child := range value.Children {
				if childNode, ok := nodes[child]; ok {
					childNode.Parent = node
					node.Children = append(node.Children, childNode)
				} else {
					childNode := DaySevenNode{Name: child, Parent: node}
					nodes[child] = &childNode
					node.Children = append(node.Children, &childNode)
				}
			}
		} else {
			node := DaySevenNode{Name: value.Name, Weight: value.Weight}
			for _, child := range value.Children {
				if childNode, ok := nodes[child]; ok {
					childNode.Parent = &node
					node.Children = append(node.Children, childNode)
				} else {
					childNode := DaySevenNode{Name: child, Parent: &node}
					nodes[child] = &childNode
					node.Children = append(node.Children, &childNode)
				}
			}
			nodes[value.Name] = &node
		}
	}

	root := false
	node := nodes[rawData[0].Name]
	for root == false {
		if node.Parent == nil {
			break
		} else {
			node = node.Parent
		}
	}
	return node.Name
}
