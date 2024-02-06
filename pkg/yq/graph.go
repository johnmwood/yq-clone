package yq

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

/*
PrintNodeTree prints the yaml graph in the exact format as it is in the file.
For example:

the yaml file quotes.yaml has the following content:

quotes:
  - id: 1
    quote: It is bad for a young man to sin; but it is worse for an old man to sin.
    author: Abu Bakr (R.A)
  - id: 2
    quote: If You Are Out To Describe The Truth, Leave Elegance To The Tailor.
    author: Albert Einstein
  - id: 3
    quote: >-
    O man you are busy working for the world, and the world is busy trying to
    turn you out.

PrintNodeTree will print the same text with indentation, newlines, dashes, and colons
*/
func PrintNodeTree(node *yaml.Node, indent int) {
	if node.Kind == yaml.DocumentNode {
		PrintNodeTree(node.Content[0], indent)
		return
	}
	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			PrintNodeTree(node.Content[i], indent)
			PrintNodeTree(node.Content[i+1], indent+1)
		}
		return
	}
	if node.Kind == yaml.SequenceNode {
		for i := 0; i < len(node.Content); i++ {
			fmt.Printf("- ")
			PrintNodeTree(node.Content[i], indent+1)
		}
		return
	}
	fmt.Printf("%s%s\n", strings.Repeat("  ", indent), node.Value)
}
