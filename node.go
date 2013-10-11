/**
 * Author:        Tony.Shao
 * Email:         xiocode@gmail.com
 * Github:        github.com/xiocode
 * File:          node.go
 * Description:   template node
 */

package template

type Node interface {
	Each() []Node
	Render() (string, error)
	Blocks() ([]string, error)
}

type BaseNode struct{}

func (b BaseNode) Each() []Node {
	return nil
}

func (b BaseNode) Render() (string, error) {
	return "", nil
}

func (b BaseNode) Blocks() ([]string, error) {
	for _, child := range b.Each() {
		child.Blocks()
	}
	return nil, nil
}

type TextNode struct {
	BaseNode
	value string
	line  int
}

type ExpressionNode struct {
	BaseNode
	expression string
	line       int
	raw        bool
}

type StatementNode struct {
	BaseNode
	statement string
	line      int
}

type BlockNode struct {
	BaseNode
}

type IntermediateControlBlockNode struct {
	BaseNode
}

type ExtendsBlockNode struct {
	BaseNode
}

type NamedBlockNode struct {
	BaseNode
	name string
	body []Node
	line int
}

type ChunkListNode struct {
	BaseNode
	chunks []Node
}

type FileNode struct {
	BaseNode
}
