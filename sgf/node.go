package sgf

type Node struct {
	Point      Property
	Properties []Property
	Variations []*Node
	Next       *Node
}

func (node Node) variationString() string {
	if len(node.Variations) == 0 {
		return ""
	}

	str := ""
	for _, nodevar := range node.Variations {
		nodestr := ""
		for nptr := nodevar; nptr != nil; nptr = nptr.Next {
			nodestr += nptr.String()
		}
		str += "(" + nodestr + ")"
	}
	return str
}

func (node Node) propertiesString() string {
	str := ""
	for _, prop := range node.Properties {
		str += prop.String()
	}
	return str
}

func (node Node) String() string {
	return ";" +
		node.Point.String() +
		node.propertiesString() +
		node.variationString()
}

func (node *Node) AddProperty(prop Property) {
	switch prop.Name {
	case "B", "W":
		node.Point = prop
	default:
		node.Properties = append(node.Properties, prop)
	}
}

func (node *Node) NewNode() *Node {
	node.Next = new(Node)
	return node.Next
}

func (n *Node) NewVariation() *Node {
	node := new(Node)
	n.Variations = append(n.Variations, node)
	return node
}
